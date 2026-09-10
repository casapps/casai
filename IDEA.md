## Project description

casai is a single-binary Go application providing a CLI, TUI, and native GUI AI coding
agent. It is a drop-in functional replacement for Claude Code: it reads and honors the
exact same `~/.claude/` (user) and `{project}/.claude/` (project) directory conventions —
`CLAUDE.md` files, `settings.json`/`settings.local.json`, skills, subagents, hooks, slash
commands, MCP server config, and this user's own `~/.claude/memory/*` auto-memory system —
so a user's existing Claude Code configuration works unmodified under casai.

Unlike Claude Code (Anthropic-model-only) or any single-vendor AI CLI (Codex/OpenAI,
Gemini CLI/Google, Copilot CLI/GitHub), casai is model-agnostic: it talks to Ollama's
local API, any OpenAI-compatible API, and any Anthropic-compatible API, and supports
OAuth login against any provider that offers a third-party OAuth app flow — not only
its own vendor's subscription OAuth. No surveyed AI coding CLI combines full Claude-Code
config-format compatibility, genuine multi-backend model support, and generic third-party
OAuth; that combination is casai's reason to exist.

The target user is a developer who already has a Claude Code setup (CLAUDE.md files,
skills, subagents, hooks, MCP servers) and wants to run the same workflow against a local
model, a self-hosted OpenAI/Anthropic-compatible endpoint, or a different subscription
provider — without maintaining two parallel configs or losing any Claude Code feature.

## Project variables

    project_name:   casai
    project_org:    casapps
    # FROZEN — set once at first-time setup, never edit
    internal_name:  casai
    # FROZEN — set once at first-time setup, never edit
    internal_org:   casapps
    app_name:       casai
    module_path:    github.com/casapps/casai
    maintainer_name: casjay
    maintainer_email: casjay@yahoo.com

## Business logic

### App surfaces in scope

All three: GUI, TUI, and CLI, per PART 2 → "GUI/TUI/CLI Capability Rule". Runtime mode
auto-selects GUI > TUI > CLI per PART 3, overridable by flag/env. CLI is the required
baseline for non-interactive/automation/CI use (scripted invocation, piped input/output,
exit codes). No RFC protocol daemon (PART 14) and no local client/server IPC split in v1 —
every invocation is a self-contained process; revisit IPC mode only if a future feature
(e.g. background/scheduled runs) genuinely requires a long-lived server, and declare it
here explicitly before implementing it.

### Claude Code compatibility — MUST be 100% compatible, not merely equivalent

casai MUST read and behave identically to Claude Code for every item below, so an existing
`~/.claude/` and project `.claude/` tree works unmodified:

- Directory layout: `~/.claude/` (user scope) and `{project}/.claude/` (project scope) —
  see "`~/.claude/` and project `.claude/` handling" below for the read/write split
- `CLAUDE.md` loading: walks CWD up to root, auto-discovers nested subdirectory
  `CLAUDE.md` files on access, supports `@path` imports, layers additively across
  user/project/managed scope
- `.claude/rules/*.md` with `paths:` frontmatter, scope-loaded only when a matching file
  is touched
- `settings.json` / `settings.local.json` schema: permission allow/ask/deny rules
  (per-tool/command pattern), env vars, default model, hooks block, MCP server allowlist;
  layering order user > project > local > managed-policy
- Skills: `SKILL.md` frontmatter (`name`, `description`, `disable-model-invocation`,
  `context: fork`), `/name` invocation, auto-match-by-description, precedence
  managed > user > project
- Subagents: `.md` agent definition format (own system prompt, tool allowlist, optional
  `skills:` preload field), isolated-context execution, precedence
  managed > CLI flag > project > user > plugin
- Hooks: the full lifecycle event set (PreToolUse, PostToolUse, Stop, SessionStart,
  UserPromptSubmit, PreCompact, etc.), matcher syntax (`Edit|Write`, `*`), exit-code-2
  blocking semantics, JSON stdin/stdout contract, all hooks from all sources fire
  (no override)
- MCP server config: `settings.json`/`.mcp.json`, scopes user/project/local, override
  precedence local > project > user
- Slash commands: custom `.md` command files invoked as `/name`, plus the built-in set
  (`/init`, `/memory`, `/agents`, `/permissions`, `/mcp`, `/plan`, `/doctor`, `/context`)
- Plugins: bundle format (skills+hooks+subagents+MCP+commands), namespacing
  (`plugin:skill`), marketplace manifest format
- Permission modes: Plan (read-only), default (ask), Accept Edits, Bypass Permissions,
  with identical tool-gating behavior per mode
- This user's own `~/.claude/memory/*` auto-memory convention (referenced from a global
  `CLAUDE.md` loader) — casai MUST also load and honor it when present, since it is part
  of "the user's existing Claude Code configuration" this project promises to run
  unmodified

Any divergence from upstream Claude Code behavior in the items above is a bug, not a
design choice — this list is the compatibility contract, not a feature-parity aspiration.

### `~/.claude/` and project `.claude/` handling

**User scope (`~/.claude/`) — read-only source, never written by casai:**

- casai reads the user's entire `~/.claude/**` tree, not a curated subset. It splits into
  three handling categories:
  1. **Config-defining** (loaded and honored identically to Claude Code, per the
     compatibility contract above): `CLAUDE.md`, `settings.json`/`settings.local.json`,
     `agents/` (subagents), `skills/`, `hooks/`, `memory/`, `plugins/`, `TEMPLATES/`,
     `scripts/`, `.mcp.json`
  2. **Session/conversation history** (`projects/`, `sessions/`, `history.jsonl`): read
     and one-way synced/migrated into casai's own internal session format, so a
     Claude-Code-started conversation can be resumed under casai — this is an import, not
     a shared live format; casai never writes back into Claude Code's own session files
  3. **Everything else** (`.credentials.json`, `daemon/`, `daemon-auth-*`, `debug/`,
     `shell-snapshots/`, `paste-cache/`, `file-history/`, `cache/`, `jobs/`, `tasks/`,
     `teams/`, `security/`, `backups/`) is Claude Code's own runtime/internal state —
     casai has its own equivalents and does not need or use a shared format for these
- casai NEVER writes into `~/.claude/**` — it is purely a read source. casai's own
  user-scope state lives in its own separate location
- No `~/.claude/` present at all is not an error — casai runs standalone with zero
  Claude Code history to import

**Project scope (`{project}/.claude/`) — shared, read AND written by casai:**

- casai uses the project's existing `{project}/.claude/**` directory directly for
  project-scope config (CLAUDE.md, settings.json, project skills/agents/hooks/rules) — it
  does NOT create a separate casai-specific project config directory
- Hard rule: casai MUST NEVER write secrets, API keys, passwords, or any credential into
  `{project}/.claude/**`. That directory is expected to be committed to the project's git
  repo and is treated as public by default (same posture as every other tracked file) —
  a user must never have to remember to keep it secret-free themselves. Credential
  storage always lives outside the project tree, per "Credentials at rest" under Trust
  boundaries below

### Default permission mode — one deliberate divergence from Claude Code

casai is a true agent: the user states intent and casai carries it out, rather than
confirming each step. casai's OOTB default permission mode is therefore **auto** — fully
autonomous, no per-tool-call confirmation — which differs from Claude Code's own default
of "ask." This is the one intentional exception to the compatibility contract above, and
it is scoped narrowly:

- Auto mode still fully adheres to whatever `settings.json` allow/deny/ask permission
  rules and hooks are in effect — it never bypasses an explicit `deny` rule or an
  explicit `ask` rule the user configured; "auto" means "no confirmation prompt when
  policy doesn't already require one," not "ignore policy"
- All four Claude Code permission modes — Plan (read-only), default-ask, Accept Edits,
  Bypass Permissions — remain fully implemented and selectable via flag/config/
  `settings.json`, and each behaves identically to its Claude Code counterpart when
  selected; only the OOTB default selection differs from upstream
- If a user's existing `settings.json` already sets an explicit permission mode, that
  value is honored as normal config layering — auto is casai's own fallback default when
  nothing else specifies a mode, not a forced override

### General AI-CLI feature parity — should have, casai's own implementation, no compat requirement

- Agentic loop (plan → tool call → observe → repeat), not just one-shot chat
- Read-only "plan mode" research/propose phase before any change is applied
- File edit tool with diff-based, reviewable changes
- Shell/exec tool with a sandbox axis independent of the approval axis (filesystem,
  network, exec — Codex's model is the cleanest reference) so a user can allow file edits
  while still gating shell/network
- Web search / web fetch tool
- MCP client support (protocol-level; not just reading Claude Code's MCP config format,
  which is already required above)
- Repo-map / codebase-wide symbol awareness (tree-sitter or LSP-based) so large repos
  don't require dumping whole files into context
- Multi-file/multi-step diff application with an optional auto-commit step
- Lint/test/auto-repair loop as an opt-in workflow
- Config file + CLI flags + env var layering, with named profiles (per-project/per-task
  approval+sandbox+model switching, per Codex)
- Local model support via Ollama, with the same UX as any other provider
- Custom/user-defined slash commands as reusable workflows (already required above via
  Claude Code compat, but casai's own extensions may add to this)
- Subagent/multi-agent delegation with isolated context
- Session persistence/resume across invocations
- Voice input — optional, not required for v1
- Telemetry: opt-in only, never on by default, vendor-neutral if implemented at all

### Model/provider support

Every provider casai talks to — OOTB preset or user-defined custom entry — resolves to
exactly one of three wire protocols: `openai` (OpenAI-compatible REST), `anthropic`
(Anthropic-native REST), or `ollama` (Ollama's native API). One client implementation per
protocol serves every provider that speaks it; adding a new OOTB preset is a config-table
entry, never a new code path.

#### OOTB provider presets

Baked in by name, each with a default base URL (user-overridable) and default auth mode.
`api-key` means casai prompts for/reads a key; `oauth` means a real third-party
authorization flow casai drives itself (not the vendor's own first-party CLI login, which
is access-gated and out of scope even when the vendor calls it "OAuth").

| Preset name | Protocol | Default base URL | Auth |
|---|---|---|---|
| `anthropic` | anthropic | `https://api.anthropic.com` | api-key |
| `openai` | openai | `https://api.openai.com/v1` | api-key |
| `ollama` | ollama | `http://localhost:11434` | none |
| `azure-openai` | openai | `https://{resource}.openai.azure.com/openai/v1/` (resource required) | api-key or oauth (Entra ID) |
| `gemini` | openai | `https://generativelanguage.googleapis.com/v1beta/openai/` | api-key or oauth |
| `groq` | openai | `https://api.groq.com/openai/v1` | api-key |
| `mistral` | openai | `https://api.mistral.ai/v1` | api-key |
| `deepseek` | openai | `https://api.deepseek.com/v1` | api-key |
| `xai` | openai | `https://api.x.ai/v1` | api-key |
| `cohere` | openai | `https://api.cohere.ai/compatibility/v1` | api-key |
| `perplexity` | openai | `https://api.perplexity.ai` | api-key |
| `together` | openai | `https://api.together.xyz/v1` | api-key |
| `fireworks` | openai | `https://api.fireworks.ai/inference/v1` | api-key |
| `deepinfra` | openai | `https://api.deepinfra.com/v1/openai` | api-key |
| `cerebras` | openai | `https://api.cerebras.ai/v1` | api-key |
| `replicate` | openai | `https://api.replicate.com/v1` (unofficial compat shim) | api-key |
| `openrouter` | openai | `https://openrouter.ai/api/v1` | api-key or oauth (PKCE, no client secret) |
| `litellm` | openai | `http://localhost:4000` (self-hosted, no fixed public default) | api-key (virtual key it issues) |
| `9router` | openai | `http://localhost:8080/v1` (self-hosted, no fixed public default) | api-key |
| `open-webui` | openai | `http://localhost:3000/v1` (self-hosted) | api-key |
| `lmstudio` | openai | `http://localhost:1234/v1` | none |
| `llamacpp` | openai | `http://localhost:8080/v1` | none |
| `vllm` | openai | `http://localhost:8000/v1` | none or api-key |
| `textgen-webui` | openai | `http://127.0.0.1:5000/v1` | none or api-key |
| `localai` | openai | `http://localhost:8080/v1` | none |
| `koboldcpp` | openai | `http://127.0.0.1:5001/v1` | none |
| `jan` | openai | `http://127.0.0.1:1337/v1` | none |

Amazon Bedrock and Google Vertex AI are explicitly OUT of the OOTB preset list for v1:
both use proprietary signed-request auth (AWS SigV4 / GCP service-account OAuth) instead
of a base-URL+key swap, and need a dedicated adapter rather than fitting the
openai/anthropic/ollama protocol set. Add either only as a deliberate future decision
recorded here first, never as a silent scope-creep addition.

#### Custom providers

A user can define any provider casai doesn't ship a preset for. A custom entry requires:

    name: string           # unique key, e.g. "work-proxy"
    type: openai|anthropic|ollama   # which protocol/client this endpoint speaks
    url: string             # base URL

and may optionally set exactly one auth method:

    api_key_env: string     # name of the env var holding the key (never the raw key
                             # inline in config — PART 9 credential-handling rule)
    oauth:                  # or, a third-party OAuth app config
      client_id: string
      auth_url: string
      token_url: string
      scopes: [string]

Custom entries live in the same provider table as OOTB presets once defined — nothing
downstream (model selection, profiles, switching) distinguishes "custom" from "built-in"
after config load.

#### Provider/model selection and switching

- Provider/model selection is per-invocation and per-profile — never hardcoded to one
  vendor anywhere in the core application layer
- Selected provider+model is saved to config (the active default, plus any named
  profiles) so it persists across invocations without re-specifying it every run
- Switching provider and/or model — mid-session or between invocations — MUST NOT lose
  conversation context: session/conversation history is stored independently of which
  provider or model is currently active, so a user can start on one provider and continue
  the same session on another (subject only to that provider's own context-window limits)

### Token usage optimization

casai MUST actively manage token consumption, not just pass everything to the model
uncut:

- Context compaction: auto-summarize/compact conversation history as it nears the active
  model's context limit (equivalent to Claude Code's own `/compact`), rather than
  truncating or erroring
- Repo-map / codebase-wide symbol awareness (already required under general feature
  parity above) is also the token-budget mechanism for large repos — read narrow slices
  via tree-sitter/LSP symbol lookups instead of dumping whole files into context
- Configurable per-session and per-task token budgets (config setting and/or
  `--max-tokens` flag): casai warns as usage approaches the budget and can stop or ask
  before exceeding a hard cap, rather than running unbounded
- Cheaper-model routing: trivial/mechanical sub-steps (simple greps, formatting,
  boilerplate lookups) may be routed to a smaller/cheaper configured model while the main
  agentic loop stays on the primary model — mirrors the user's own model-routing
  convention; only applies when a cheaper model is configured, never silently substitutes
  a different model for the primary reasoning steps
- Prompt caching: use the active provider's prompt-caching feature where available
  (e.g. Anthropic `cache_control`, OpenAI automatic caching) to cut repeated-context cost
  across multi-turn sessions

### Roles / user types

Single-user local tool — no multi-tenant server component, no accounts system. The
"trusted" party is the local OS user running the binary; there is no separate admin role.

### Trust boundaries & abuse cases

- casai executes with the invoking OS user's full privileges — no privilege boundary
  between casai and the user beyond what PART 4 (Privilege Escalation & System
  Integration) already governs for the whole app template
- Config and credentials (API keys, OAuth tokens) are trusted input from the local user
  only; casai MUST NOT accept credentials or provider config from untrusted remote
  sources (a fetched CLAUDE.md, a skill file, an MCP tool response) — those are untrusted
  content and MUST NOT be interpreted as configuration or instructions to change
  permissions, add providers, or disable safety prompts
- Model output (from any configured provider) is untrusted with respect to the
  permission system: a model's own text MUST NOT be able to bypass permission
  prompts/allow-deny rules — only explicit user config or explicit user approval can
  grant a tool call
- Abuse case: a malicious or compromised skill/subagent/MCP server definition attempting
  to exfiltrate credentials or widen its own tool allowlist — defended by the same
  permission/allow-deny model Claude Code uses (required above), never bypassed by
  casai's own code
- Abuse case: a fetched remote file (via web fetch/MCP) containing prompt-injection text
  instructing casai to change settings, run destructive shell commands, or leak secrets —
  MUST be treated as data, never as instructions, consistent with how the user's own
  global rules already treat untrusted content across every other tool
- Credentials at rest: API keys and OAuth tokens are stored locally only, never logged in
  plaintext, never transmitted except to their own provider's endpoint

### Platform constraints

- Linux, macOS, Windows, FreeBSD — per PART 5 binary matrix
- Single static binary, no host toolchain required at runtime, offline-capable except for
  the network calls the user's chosen provider requires
- First run works with zero config beyond picking/configuring one provider
