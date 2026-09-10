# {PROJECT_NAME} Specification

**Name**: {project_name}

**About this file:** `AI.md` is the complete, authoritative specification for this project. Throughout this document, all references to `AI.md` refer to this file.

**Note:** `{PROJECT_NAME}` and `{project_name}` in this file are reference tokens, not setup-time text replacements. Their values are resolved from `IDEA.md ## Project variables` while `AI.md` remains read-only.

---

# PROJECT DESCRIPTION

**See `IDEA.md` for project-specific details.**

---

# SOURCE OF TRUTH AND IDEA.md PRECEDENCE

**See `IDEA.md` for features, data models, and business rules.**

IDEA.md is the project PLAN. AI.md (this file) is the SOURCE OF TRUTH.

| File | Role | Update When |
|------|------|-------------|
| **AI.md** | SOURCE OF TRUTH - implementation rules (read-only) | No — use SPEC.md for project-specific rule overrides |
| **SPEC.md** | Project-specific rule overrides — created only when a rule must contradict this specification or global. May be empty. SPEC.md wins over AI.md. | When a project rule must differ from this specification or global |
| **IDEA.md** | PROJECT PLAN - must follow AI.md | Features change, project variables change |

**Rule hierarchy:** SPEC.md > AI.md > global CLAUDE.md. If SPEC.md and AI.md conflict, SPEC.md wins — that is its purpose.
**Rule:** If AI.md and IDEA.md conflict, AI.md wins. Fix IDEA.md.

## IDEA.md Required Layout

**Every IDEA.md MUST have exactly these three top-level sections, in this order. For a fillable template plus worked examples (Notes / Feeds / dotctl), see PART 13 → "IDEA.md REFERENCE".**

```markdown
## Project description

(Full project description — what the project is, who uses it, what problem it solves.)

## Project variables

(All project variables in `key: value` form. Required keys at minimum: `project_name`,
`project_org`, `internal_name`, `internal_org`. Add more as the project needs — `app_name`,
`official_site`, `maintainer_name`, `maintainer_email`, `module_path`, etc.)

Example:

    project_name:  notes
    project_org:   casjay
    # FROZEN — set once at first-time setup, never edit
    internal_name: notes
    # FROZEN — set once at first-time setup, never edit
    internal_org:  casjay
    app_name:      Notes
    module_path:   github.com/casjay/notes
    official_site: https://notes.example.com

## Business logic

(Full business spec — the WHAT, not the HOW. Features, data models, user flows,
permission rules, business invariants, platform targets, input modes, accessibility,
security assumptions, and any exceptions.)
```

**Rules for `## Project variables`:**
- One variable per line: `key: value`
- Keys are **lower_snake_case** only
- The setup flow renders `{KEY_UPPER}` automatically by uppercasing the lowercase key
- Never guess values: use commands and existing files
- If a placeholder referenced by AI.md has no entry in `## Project variables`, setup MUST stop and ask instead of inventing a value

**Rules for `## Business logic`:**
- It MUST define the actual product scope for THIS project - not generic boilerplate
- It MUST state which app surfaces exist: GUI, TUI, CLI, or a subset — and whether the project is an RFC protocol daemon (PART 14) or a local client/server (IPC) split (PART 2 → "Local Client/Server (IPC) Mode")
- It MUST define user flows, stored data, trust boundaries, abuse cases, and platform constraints
- If a security-sensitive choice is intentionally allowed, the reason MUST be documented there

## Migrating Existing `CLAUDE.md` Into `IDEA.md`

**If a repository already has a pre-existing `CLAUDE.md` or `.claude/CLAUDE.md` with real project details, those project details MUST be migrated into `IDEA.md`.**

**What belongs in `IDEA.md`:**
- project description / elevator pitch
- project-specific terminology
- project variables that can be expressed as `key: value`
- business logic, roles, flows, constraints, trust boundaries, abuse cases, and security exceptions

**What does NOT belong in `IDEA.md`:**
- generic Claude/Copilot usage instructions
- loader boilerplate whose job is only to point at `AI.md`
- duplicated global implementation rules that already live in `AI.md`
- stale code snippets, one-off notes, or tool chatter with no business/spec value

**Migration rules:**
1. Read existing `CLAUDE.md` and `.claude/CLAUDE.md` first - never overwrite blindly
2. Extract valid project-specific content and reorganize it into the required `IDEA.md` layout
3. Normalize discovered variables into lower_snake_case `key: value` entries
4. If `internal_name` cannot be proven, initialize it to `project_name` on first migration and treat it as frozen after that. Do the same for `internal_org` ← `project_org`.
5. If statements from `CLAUDE.md` or `.claude/CLAUDE.md` conflict with `AI.md`, `AI.md` wins
6. After migration, keep root `CLAUDE.md` and/or `.claude/CLAUDE.md` only as short efficient loaders and keep the real plan/spec in `IDEA.md`
7. Never silently discard meaningful project-specific content; migrate it, trim it, or explicitly ask where it belongs

---

# PART 0: CRITICAL RULES - READ FIRST

## THIS IS A STRICT SPECIFICATION - NOT GUIDELINES

- Every item in this specification MUST be followed exactly unless explicitly marked optional
- This is not a suggestion document
- There are no silent exceptions
- If the spec says X, do X - not "improved X"
- If something seems wrong, follow it and flag it; do not silently rewrite intent

## Attribution

**AI operates on behalf of the user in a Senior Developer / UI-UX Designer capacity.**

## ⚠️ CRITICAL: File Paths and Project Root

- All paths are relative to the project root unless explicitly noted
- Do not scatter top-level files unnecessarily
- Runtime-generated files are not committed
- AI must not move the project root or invent sibling repositories
- **Never edit files outside the project root to work around a problem inside it** — no host system files (`/etc/*`, shell rc files, systemd units), no sibling/other repositories, no global tool configs. A failing build, missing dependency, or broken tool inside the project is fixed in the project's own code/config/Docker image, never by reaching outside it. The only exception is an external path the user explicitly names for that specific task

## ⚠️ CRITICAL: AI.md is the Source of Truth

- `AI.md` is read-only during routine work
- `IDEA.md` is where project-specific values and product rules live
- Loader files (`CLAUDE.md`, `.claude/CLAUDE.md`) stay short and point back to `AI.md`
- If a loader file and `AI.md` disagree, `AI.md` wins
- `.claude/memory/` holds durable, project-specific knowledge discovered during
  development — decisions, gotchas, conventions unique to this codebase — distinct
  from `AI.md` (spec) and the loader files above. Committed to the repo, not
  gitignored. One markdown file per topic, YAML frontmatter (`name`, `description`,
  `type: project`), indexed by `.claude/memory/MEMORY.md`, read on demand. Same
  credential-masking rule as everywhere else. `~/.claude/**` (global) stays
  read-only, deployed only via `claudemgr/config`'s `install.sh`; `.claude/memory/`
  here is read/write in this repo directly

## ⚠️ CRITICAL: Keep Documentation in Sync

Update these when their subject changes:
- `IDEA.md` when features or variables change
- `README.md` when install, usage, or packaging changes
- `LICENSE.md` when dependencies or attribution changes
- CI/CD docs or scripts when release mechanics change

## ⚠️ CRITICAL: One Coherent Product

This specification defines **one Go application** with shared core logic and up to three presentation layers:
- GUI (preferred when available)
- TUI (fallback for interactive terminals)
- CLI (fallback for non-interactive/plain execution)

Single-process, single binary. The app may make outbound network calls to remote services it consumes; it does not host any.

## ⚠️ CRITICAL: No Host Toolchain or Binary Execution

The Go toolchain and any compiled artifact from this project MUST NOT run on the host machine.

- Never invoke `go`, `go build`, `go test`, `go run`, `go fmt`, `go vet`, `golangci-lint`, `govulncheck`, `go install`, or any project binary built from source directly on the host
- Every build, test, lint, format-check, doc-gen, run, install, package, and manual debug session executes inside a project-provided Docker container
- The host's role is limited to editing source files, version control, and orchestrating Docker
- CI and local commands documented in this spec are illustrative shapes; the **actual invocation** is always Docker-wrapped (see PART 5 → "Docker Rule")
- If a contributor's environment cannot run Docker, they cannot build/test this project — that is intentional, not a bug

This rule has no opt-out. There is no "just this once" exception for `go test` on the host.

## ⚠️ CRITICAL: X11 AND Wayland Are Both Required

If this project ships a GUI surface, it MUST support **both** X11 and Wayland as first-class display backends.

- Wayland-only GUIs are not acceptable
- X11-only GUIs are not acceptable
- The chosen GUI stack/toolkit must be one that natively supports both, or the app must integrate both backends
- Display detection at runtime must consider `WAYLAND_DISPLAY` AND `DISPLAY` and pick the appropriate backend
- GUI smoke testing inside Docker MUST be runnable against both an X11 socket and a Wayland socket forwarded from the host (see PART 5 → "Docker Rule")
- IDEA.md may declare GUI out of scope, but it may NOT declare "X11 only" or "Wayland only"
- Reconciling X11/Wayland with the static-binary rule (see "Single Static Binary" below): use display packages that avoid **link-time** C dependencies. With `CGO_ENABLED=0` this is automatic — for X11, the preferred option is a pure-Go X11 client (`github.com/jezek/xgb`, talks to the X server's Unix socket directly with no `libX11` involvement); for Wayland, use a pure-Go Wayland client (`github.com/rajveermalviya/go-wayland`, speaks the wire protocol directly with no `libwayland-client` involvement). In all of these cases the binary has no link-time C dependency for display I/O — that is the rule. The pure-Go GUI toolkits (e.g., `gioui.org`, and toolkits built on it) follow this path natively; that path is compliant.

## ⚠️ CRITICAL: Go-Only Application

This project's source code is **exclusively Go**.

- All application code, library code, build automation, and test code in this repository is written in Go
- No C, C++, Objective-C, Swift, Python, JavaScript, TypeScript, or shell-script source files contribute to the produced binary
- `Makefile` and `scripts/` are used for build automation — they orchestrate Docker invocations but MUST NOT contain application logic
- Small `docker/rootfs/usr/local/bin/entrypoint.sh` and `docker/` shell helpers are tolerated because they orchestrate the container, not the application; they MUST NOT contain application logic
- `CGO_ENABLED=0` ALWAYS — pure Go, no C, no exceptions
- Third-party packages that are pure Go are strongly preferred; packages that vendor C are allowed only when (a) no pure-Go equivalent is viable, (b) the C code is statically compiled into the final binary (c) it does NOT require a system C lib at runtime, and (d) the dependency is documented in `IDEA.md` and `LICENSE.md`
- **Prefer pure-Go packages whenever a viable one exists.** Pure-Go packages are what make the single-static-binary rule and cross-platform GUI (Windows / macOS / Linux / BSD) achievable in practice — every C dependency dragged in becomes a portability and build-system tax. See PART 5 → "Pure-Go Library Stack" for the recommended package list
- Never introduce a build that requires a system C/C++ toolchain on the user's machine — the project Docker image is the only required build environment

## ⚠️ CRITICAL: Single Static Binary

The deliverable for each supported target is **one statically linked binary** that runs without external runtime dependencies on the user's system beyond the kernel and (where applicable) display server sockets.

**Required build outputs per target:**

| Target | Linkage | Notes |
|--------|---------|-------|
| `GOOS=linux GOARCH=amd64 CGO_ENABLED=0` | fully static | Default Linux release target |
| `GOOS=linux GOARCH=arm64 CGO_ENABLED=0` | fully static | Default ARM64 Linux release target |
| `GOOS=windows GOARCH=amd64 CGO_ENABLED=0` | fully static | Single `.exe` with no runtime DLL dependency |
| `GOOS=windows GOARCH=arm64 CGO_ENABLED=0` | fully static | Single `.exe` |
| `GOOS=darwin GOARCH=amd64 CGO_ENABLED=0` | system frameworks only (Apple does not allow static libSystem) | "Static" here means: no third-party dynamic libraries; only Apple-provided frameworks resolved at runtime |
| `GOOS=darwin GOARCH=arm64 CGO_ENABLED=0` | system frameworks only | Same rule as amd64 macOS |
| `GOOS=freebsd GOARCH=amd64 CGO_ENABLED=0` | fully static | Default FreeBSD release target |
| `GOOS=freebsd GOARCH=arm64 CGO_ENABLED=0` | fully static | Default ARM64 FreeBSD release target |

**Rules:**
- No `glibc` runtime dependency on Linux release artifacts — `CGO_ENABLED=0` produces a fully static binary on Linux without needing musl
- No third-party `.so` / `.dylib` / `.dll` shipped alongside the binary
- No `LD_LIBRARY_PATH`, `DYLD_LIBRARY_PATH`, or wrapper-script tricks to find runtime libs
- `go build` inside the Docker image MUST produce a binary that passes a "no unexpected dynamic deps" check (`ldd`, `otool -L`, `dumpbin /dependents`) appropriate to the target
- X11 and Wayland are runtime-discovered display sockets, not link-time dependencies — see "X11 AND Wayland Are Both Required"
- Plugin systems, `dlopen` of arbitrary user code, and runtime extension loading from disk are forbidden unless IDEA.md explicitly defines a hardened plugin contract

## ⚠️ CRITICAL: Self-Contained Assets

The single binary contains **everything the app needs to function**. The user is not required to install fonts, themes, icons, templates, schemas, locales, or any other support file separately.

- Embed assets at build time using `//go:embed` (standard library `embed` package)
- Fonts, icons, theme data, default config, JSON/YAML schemas, SQL migrations, web assets (if any), localization catalogs, default templates, and licenses-to-display all live inside the binary
- The `assets/` directory in the repo is a **build-time source tree** for embedding — it is not a runtime install target
- The app may **read** user-provided config and data from per-user paths (PART 4 → "Path Rule"), and may **write** cache/state there, but it must run with full default functionality when those paths are empty
- Network/CDN fetches at first run to "download missing assets" are forbidden
- A new install on an air-gapped machine MUST work end-to-end with only the binary present

## Licensing & Features

| Rule | Description |
|------|-------------|
| **MIT License** | All project code is MIT licensed unless IDEA.md explicitly states an additional compatible license policy |
| **3rd party attribution** | All third-party licenses are listed in `LICENSE.md`. Generation, allowlist/denylist, IDEA.md exception path, and the user-visible licenses surface live in PART 11 → "License Compliance" |
| **GPL / AGPL / LGPL denied by default** | Static linking would relicense the distributed binary away from MIT. Allowed only via a documented IDEA.md exception (PART 11 → "License Compliance") |
| **Free & open source** | No paid tiers, enterprise gating, or artificial feature segmentation |
| **No premium features** | GUI, TUI, and CLI surfaces expose the same core product capabilities where applicable |
| **No activation gates** | No license keys, phone-home unlocks, or paywalled code paths |

**NEVER implement:**
- upgrade/paywall prompts for core functionality
- hidden features enabled by payment tier
- telemetry-based licensing enforcement
- artificial limits used for monetization

---

# PART 1: PROJECT FILES & GOVERNANCE

## Project Files

| File | Purpose | Update When |
|------|---------|-------------|
| **AI.md** | Implementation spec (HOW) - SOURCE OF TRUTH, read-only | No — use SPEC.md for project-specific rule overrides |
| **SPEC.md** | Project-specific rule overrides (optional, may be empty) | When a project rule must contradict this specification or global |
| **IDEA.md** | Project plan (WHAT) | Features or variables change |
| **TODO.AI.md** | Task tracking (AI-owned) | Tasks added/completed |
| **TODO.md** | Task tracking (human-owned) | AI may mark done; never delete/empty |
| **PLAN.AI.md** | Implementation plan (AI-owned) | Planning new work |
| **PLAN.md** | Implementation plan (human-owned) | AI may mark done; never rewrite wholesale |
| **README.md** | User-facing install/usage docs | Usage changes |
| **LICENSE.md** | Project + dependency licenses | Dependency set changes |
| **release.txt** | Canonical release version when present | Release version changes |
| **site.txt** | Optional official site/homepage URL | Official site changes |
| **Makefile** | Build, test, run targets for local development. Convenience wrappers around the Docker-wrapped commands. **Never called from CI workflow `run:` steps** (PART 10 → "CI/CD Rules") | Local-dev targets change |
| **.dockerignore** | Build-context exclusions; Go-specific entries: `binaries/`, `releases/`. See `dockerfile_conventions.md` → ".dockerignore" for the full standard entry set (version control, `.env`, build artifacts, `volumes/`, OS files, markdown/LICENSE). `docker/`, `src/`, `go.mod`, `go.sum`, `release.txt` are NEVER excluded. | Build-context surface changes |
| **renovate.json** | Single Renovate config covering Go module versions, GitHub Actions SHAs, and Docker digests across all five providers. Dependabot forbidden (PART 9 → "Dependency Governance") | Update-policy changes |

## Mandatory Compliance Schedule

| When | Action | Purpose |
|------|--------|---------|
| Before each task | Read only the spec parts relevant to what you are about to implement — do not pre-load speculatively | Prevent token waste |
| Every 3-5 changes | Stop and verify against spec | Catch drift early |
| Before task completion | Full compliance check | Ensure correctness |
| When uncertain about a spec requirement | Read that specific section — never guess, never rely on prior-session memory | Accuracy without waste |

**Hook-enforced:** `spec-guard.sh` blocks Edit/Write on project files until AI.md/SPEC.md has been Read this session; the gate re-arms after every compaction.

## Self-Validation Loop

**AI MUST verify its own work with real tools before reporting a task as done. Do not rely on "the code looks right."**

**This rule applies to EVERY change type covered by this specification — library/API logic, GUI/TUI/CLI binaries, single-static-binary build, asset embedding, Docker, CI/CD, configuration, documentation, security — not only one category.** Whatever you touched, you verify.

Getting code correct on the first try is much harder than iterating with feedback. Close the loop every time. All execution goes through the project's containerised targets — never bare host Go toolchain.

| Change type | How to verify |
|-------------|---------------|
| Library / API logic | Run the project's test target inside the container; exercise the API directly; compare output against expected |
| Behavior-preserving refactor | Diff outputs of old vs. new path on representative inputs (don't trust that the diff "looks right") |
| CLI binary | Run the binary in the container; exercise relevant flags including `--help`/`--version`; check stdout, stderr, and exit code |
| TUI binary | Run in the container; verify rendering, keyboard input, and screen redraw on resize/exit |
| GUI binary | Run in the container; verify the rendered window under BOTH X11 AND Wayland forwarding (per the X11/Wayland mandate); confirm input events reach the app |
| Single static binary requirement | Confirm the artifact is a single self-contained file and that it runs on a clean container with no extra runtime install |
| Asset embedding | Confirm assets are loaded from the binary itself (not from a host path); test on a container without the source tree mounted |
| Performance change | Measure before AND after — don't assume parallelism, caching, or "cleaner" code is faster |
| Bug fix | Reproduce the bug FIRST so you have a failing signal, then verify the fix makes it disappear; add a regression test where feasible |
| Configuration / settings | Start the binary with the new config; verify defaults; verify validation rejects bad input with a useful error |
| Docker / container build | Build and verify each image variant as appropriate to the change: `:devel` (debug binary) for logic changes, `:latest` (release binary) for release verification. Run the container; smoke-test the binary inside it; for GUI, verify display forwarding still works. |
| CI/CD workflow | Run the workflow on a branch (or equivalent dry-run); verify each job's exit status, not just YAML validity |
| Logging / error paths | Trigger the error path; verify the log line/structured event was emitted with expected fields |
| Security-sensitive change (auth, crypto, input validation, plugin contracts) | Test both the success path AND attempted bypass paths; never assume a guard works without exercising it |
| Documentation / README | Render markdown locally; verify links, code samples, and example commands actually work |
| Type / lint / build correctness | The project's containerised check, golangci-lint, and build targets — green across all |

**Iteration rules:**
- A failed check is data, not failure — adjust and re-run until green
- Never report "done" while any verification is still red
- If a check reveals the change is wrong in a way that can't be patched, revert and re-plan; do not paper over a failing check
- When verification is genuinely impossible in this environment (no display, no DB, no network): say so explicitly. List what was checked and what could not be, so the user knows where to look

**Reference:** based on published guidance about AI coding agent self-validation (Eivind Kjosbakken, Towards Data Science, 2026) — when an AI agent is given verification tools (output diffing, browser/display, test runners) and allowed to iterate, one-shot success rate, run length, and task complexity all improve substantially.

## Loader Files

| Tool | Primary Loader | Alternate Loader | Personal Override |
|------|----------------|------------------|------------------|
| Claude Code | `CLAUDE.md` | `.claude/CLAUDE.md` | `CLAUDE.local.md` |

**Loader rule:** loader files stay short. Long-form product content belongs in `IDEA.md`; long-form implementation policy belongs in `AI.md`.

---

# PART 2: GO APPLICATION MODEL

## Product Model

This specification targets a **single-binary, fully self-contained Go application** that may expose:
- a native GUI
- a terminal UI (TUI)
- a plain CLI
- an **RFC protocol server (daemon)** — conditional; only when IDEA.md declares it (PART 14)
- a **local client/server (IPC) split** — conditional; only when IDEA.md declares it (see "Local Client/Server (IPC) Mode" below)

The application may make outbound network calls to consume remote services it depends on (APIs, databases, object stores, update endpoints, etc.).

**Network listening sockets (TCP/UDP) are allowed only in daemon mode** — a project that accepts inbound network connections MUST satisfy PART 14; otherwise the application opens no network listening sockets. Unix domain sockets and named pipes are local IPC, not network listening — they are governed by "Local Client/Server (IPC) Mode" below, not by PART 14.

**Distribution model:** one statically linked binary per supported target. Everything the app needs at runtime — UI assets, fonts, icons, default config, schemas, templates, locales — is embedded inside that binary. See PART 0 → "Single Static Binary" and "Self-Contained Assets."

## Local Client/Server (IPC) Mode

Some applications split into a thin client and a long-lived local server the way tmux, `nvim --listen`, and atuin's daemon do. This is allowed for any surface when IDEA.md `## Business logic` declares the client/server architecture — it does NOT trigger PART 14, which governs network protocol daemons only.

**Socket placement & permissions:**
- All sockets live in `$XDG_RUNTIME_DIR/{project_name}/` (fallback: the per-user state dir, PART 4 → "Path Rule"); directory mode `0700`, sockets `0600` — never world-accessible, never in shared `/tmp` paths
- Private wire protocols are allowed on IPC sockets — PART 14's "never invent wire protocols" rule applies to network sockets only
- No TCP listener by default, not even loopback — a TCP listener changes the trust boundary and requires its own explicit IDEA.md declaration

**Main socket + named servers:**
- The first server started is the **main server** and owns the **main socket**: `$XDG_RUNTIME_DIR/{project_name}/{project_name}.sock`
- The user may start additional **named servers**; each gets its own socket: `$XDG_RUNTIME_DIR/{project_name}/{server_name}.sock` (server names validated against `[a-z0-9._-]+`; no path separators, never colliding with the main socket name)
- **The main server is aware of all servers** — it keeps a registry of every named server; named servers announce themselves to the main socket on startup and deregister on clean shutdown
- Clients enumerate all running servers through the main socket (e.g. a `{project_name} servers` / `list-servers` command)
- **Registry self-heal:** on main-server start and on every listing, the registry reconciles against a scan of the socket directory — each socket is connect-tested, and dead/stale entries are dropped from the registry and their socket files unlinked
- A missing main server never blocks a named server from starting — the named server retries registration in the background and is picked up by the main server's next reconcile scan

**Advanced capabilities (optional — each one used only when IDEA.md `## Business logic` declares it):**
- **Session state awareness** — the server may track live per-session/per-client state: working directory, foreground command, environment snapshot, window/pane layout, whatever the product needs; clients query it over the IPC socket
- **Save/restore** — full server state (sessions, layout, per-session working directory, running-command metadata) serializes to the per-user state dir (PART 4 → "Path Rule") as `0600` files, written atomically (temp file + rename); an explicit `save` / `restore` command pair
- **Continuous persistence (continuum-style)** — periodic autosave on a configurable interval plus save on clean shutdown, and restore-on-start reattaches the previous state; autosave runs in the background and never blocks the serving path
- **Restore never re-executes commands by default** — a default restore recreates sessions, layout, and working directories only; re-launching previously running programs happens only when the user has explicitly enabled it (per program or per profile), mirroring tmux-resurrect's opt-in model
- **Privacy:** tracked state (working directories, command lines, environment) is sensitive data — stored only in per-user `0600` files, never logged at default verbosity, never transmitted off-machine, and excluded from any diagnostics output unless the user explicitly opts in

**Trust boundary:** the IPC surface trusts same-UID local processes only — access control is the `0700`/`0600` filesystem permissions above; peer-credential checks (`SO_PEERCRED` on Linux, `getpeereid` on BSD/macOS) are the defense-in-depth verification. IDEA.md's **Trust boundaries & abuse cases** block MUST name this boundary (abuse case: another local user or process attempting to reach a socket → refused by permissions and peer check).

## Architectural Rule

Use **one shared core application layer** and thin presentation adapters:

```text
main.go                     # auto-detects GUI/TUI/CLI mode unless explicitly overridden
                            # (or cmd/{project_name}/main.go for multi-binary workspaces)
internal/
├── app/                    # shared domain/application logic
├── config/                 # configuration loading + defaults
├── platform/               # OS/platform integration
├── state/                  # app state and persistence adapters
├── support/                # logging, errors, utilities
└── ui/
    ├── gui/                # optional native GUI
    ├── tui/                # optional terminal UI
    └── cli/                # optional plain CLI commands/output
```

(See PART 5 → "Project Layout" for the full repository tree, including `docker/`, `assets/`, `packaging/`, `Makefile`, `scripts/`, etc.)

**Rules:**
- Core behavior MUST live in shared packages, not be duplicated across GUI/TUI/CLI
- UI layers are adapters over the same business rules
- Optional `cmd/*/main.go` utilities are allowed only for real auxiliary tools, not as a substitute for clean architecture

## Binary Model

| Binary | Status | Purpose |
|--------|--------|---------|
| `{project_name}` | REQUIRED | Primary application binary — single statically linked artifact |
| `cmd/*` helpers | DISCOURAGED | Permitted only when IDEA.md justifies a real auxiliary tool; each helper is itself a single static binary |

**Binary naming rules:**

Distribution artifact names use Go's native GOOS/GOARCH terms directly — no mapping or normalization:

```
{project_name}-{GOOS}-{GOARCH}{.ext}
```

| Token | Allowed values |
|-------|----------------|
| `{GOOS}` | `linux`, `darwin`, `windows`, `freebsd` (raw Go GOOS value — never aliases like `macos`) |
| `{GOARCH}` | `amd64`, `arm64` (raw Go GOARCH value — never aliases like `x86_64` or `aarch64`) |
| `{.ext}` | `.exe` on Windows; empty everywhere else |

**Worked examples:**

| GOOS / GOARCH | Artifact name |
|---------------|---------------|
| `linux` / `amd64` | `{project_name}-linux-amd64` |
| `linux` / `arm64` | `{project_name}-linux-arm64` |
| `darwin` / `amd64` | `{project_name}-darwin-amd64` |
| `darwin` / `arm64` | `{project_name}-darwin-arm64` |
| `windows` / `amd64` | `{project_name}-windows-amd64.exe` |
| `windows` / `arm64` | `{project_name}-windows-arm64.exe` |
| `freebsd` / `amd64` | `{project_name}-freebsd-amd64` |
| `freebsd` / `arm64` | `{project_name}-freebsd-arm64` |

**Other rules:**
- Local (in-tree) primary binary name: `{project_name}` (no platform/arch suffix during local development inside the Docker image)
- If optional helper binaries exist, use `{project_name}-{tool}` for the in-tree name and `{project_name}-{tool}-{GOOS}-{GOARCH}{.ext}` for distribution
- Release checksums are aggregate: `sha256.txt` and `sha512.txt` in the artifact directory cover every published artifact (standard `sha256sum` / `sha512sum` `{hash}  {filename}` lines) — no per-artifact sidecar checksum files

**Single-binary rule:** the default user experience is "download one file, run it." The primary binary MUST be self-sufficient (PART 0 → "Self-Contained Assets"). Helper binaries are not a substitute for putting features into the primary binary; if a feature can live behind a CLI subcommand of the primary binary, it MUST.

## GUI/TUI/CLI Capability Rule

- A project may implement one, two, or all three surfaces
- If GUI exists, it is the preferred interactive experience on capable local desktops
- If TUI exists, it is the preferred interactive fallback for capable terminals
- CLI is the universal fallback and the default for automation/non-interactive use
- IDEA.md MUST declare which surfaces are actually in scope

---

# PART 3: RUNTIME MODE SELECTION

## Selection Priority

**Automatic runtime selection priority is:**
1. **GUI**
2. **TUI**
3. **CLI**

**Override priority is:**
1. explicit CLI flag / subcommand (`--ui gui|tui|cli`, `gui`, `tui`, `cli`, etc.)
2. config file value
3. environment override
4. automatic detection using the priority above

## Smart Detect Rules

### GUI

Choose GUI when **all** are true:
- GUI support is compiled and enabled for the project
- the process is **not** running in SSH/MOSH/remote-shell style contexts
- a desktop/session display stack is available
- the invocation is interactive enough for windowed UX

**Treat these as remote-shell / non-local GUI blockers:**
- `SSH_CONNECTION`, `SSH_CLIENT`, `SSH_TTY`
- `MOSH_IP`, `MOSH_KEY`
- explicit headless environment from config or flag

**Treat these as positive GUI/display signals (platform-appropriate):**
- `WAYLAND_DISPLAY` — Wayland session (MUST be supported by the app on Linux/BSD via the GUI toolkit)
- `DISPLAY` — X11 / XWayland session (MUST be supported by the app on Linux/BSD via the GUI toolkit)
- a local Windows desktop session
- a local macOS Aqua/session launch

**Backend selection on Linux/BSD when both signals are present:**
- Prefer Wayland when `WAYLAND_DISPLAY` is set; if connecting to that Wayland socket fails at runtime, fall back to X11 instead of erroring out
- Fall back to X11 (`DISPLAY`) when `WAYLAND_DISPLAY` is unset or the user/config explicitly requests X11
- Both backends MUST be exercised in tests (PART 0 → "X11 AND Wayland Are Both Required")

### TUI

Choose TUI when GUI is not selected and **all** are true:
- TUI support is implemented
- stdin and stdout are TTYs
- `TERM` is not `dumb`
- output is interactive (not piped for machine consumption)
- the environment is not clearly no-format/no-interaction

**Avoid TUI when any of these are true:**
- `TERM=dumb`
- stdout or stdin is not a TTY
- `CI=true` or similar automation context
- `NO_COLOR=1` or another plain-output requirement **when the TUI depends on color/alternate-screen/cursor control**
- explicit `--plain`, `--json`, `--quiet`, or script-oriented modes

### CLI

Choose CLI when:
- GUI and TUI are unavailable or inappropriate
- the app is running non-interactively
- output is piped/redirected
- machine-readable output is requested
- automation or plain-text mode is requested

CLI is the **required** fallback and the default for non-interactive work.

## Reference Detection Logic

```go
type UIMode int

const (
    UIModeCLI UIMode = iota
    UIModeTUI
    UIModeGUI
)

func DetectUIMode(env Env, caps Capabilities) UIMode {
    if forced, ok := env.ForcedUIMode(); ok {
        return forced
    }

    remoteShell := env.HasAny("SSH_CONNECTION", "SSH_CLIENT", "SSH_TTY", "MOSH_IP", "MOSH_KEY")
    displayAvail := env.Has("WAYLAND_DISPLAY") || env.Has("DISPLAY") ||
        caps.LocalWindowsSession || caps.LocalMacOSSession

    if caps.GUISupported && !remoteShell && displayAvail && caps.InteractiveLaunch {
        return UIModeGUI
    }

    plainOrNoninteractive := !caps.StdinTTY || !caps.StdoutTTY ||
        env.Is("TERM", "dumb") || env.Truthy("CI") ||
        (env.Truthy("NO_COLOR") && caps.TUIRequiresFormatting) ||
        caps.MachineOutputRequested

    if caps.TUISupported && !plainOrNoninteractive {
        return UIModeTUI
    }

    return UIModeCLI
}
```

## Behavior Rules by Mode

| Mode | Primary UX | Good For | Avoid |
|------|------------|----------|-------|
| GUI | native windows/dialogs/desktop integration | local desktop use | headless/SSH automation |
| TUI | structured terminal app | interactive terminal workflows | `TERM=dumb`, plain-output contexts |
| CLI | plain commands/text/json | automation, scripts, CI, pipes | none |

---

# PART 4: PRIVILEGE ESCALATION & SYSTEM INTEGRATION

## Core Rule

Privilege escalation support is allowed, but it is **optional** and must be used **only when the user explicitly requests it**.

## NEVER Do

- Never auto-run `sudo`, `pkexec`, UAC elevation, or similar escalation just because a path is protected
- Never silently switch from user-scope install to system-scope install
- Never require elevation for normal app launch, per-user config, or per-user updates
- Never block core app usage behind root/admin requirements unless IDEA.md explicitly defines a genuinely privileged tool

## Allowed Uses (Explicit Request Only)

- system-wide install into protected directories
- system-wide uninstall or cleanup
- privileged file associations / shell integration / desktop registration when explicitly requested
- service registration only if the project truly implements a background service and IDEA.md says so

## Default Scope Rule

Default to **user scope**:
- user config
- user data
- user cache
- per-user desktop integration
- user-local updates/installations

## Path Rule

Prefer platform-standard user directories:

**On-disk paths use the frozen pair `{internal_org}` and `{internal_name}` only — never the mutable `{project_org}` / `{project_name}`. This is the rule that protects user data across project/org renames.**

| Purpose | Linux / BSD | macOS | Windows |
|---------|-------------|-------|---------|
| Config | `~/.config/{internal_org}/{internal_name}/` | `~/Library/Application Support/{internal_org}/{internal_name}/config/` | `%AppData%\\{internal_org}\\{internal_name}\\config\\` |
| Data | `~/.local/share/{internal_org}/{internal_name}/` | `~/Library/Application Support/{internal_org}/{internal_name}/data/` | `%LocalAppData%\\{internal_org}\\{internal_name}\\data\\` |
| Cache | `~/.cache/{internal_org}/{internal_name}/` | `~/Library/Caches/{internal_org}/{internal_name}/` | `%LocalAppData%\\{internal_org}\\{internal_name}\\cache\\` |
| Logs | `~/.local/state/{internal_org}/{internal_name}/logs/` | `~/Library/Logs/{internal_org}/{internal_name}/` | `%LocalAppData%\\{internal_org}\\{internal_name}\\logs\\` |

**Rule:** Both `{internal_name}` and `{internal_org}` anchor on-disk identifiers and stable OS-registered names (Bundle IDs, package IDs, dbus names, keychain entries, updater channels). A rename of `{project_name}` or `{project_org}` MUST NOT silently move user data or change those identifiers.

---

# PART 5: GO TOOLCHAIN, BUILD, AND PACKAGING

## Toolchain Rules

| Rule | Description |
|------|-------------|
| Go version | Use the current stable Go version declared in `go.mod` and pinned in `.go-version` |
| Version pin | Single-line `.go-version` file at project root (e.g., `go1.23.0`) — used by toolchain managers and CI |
| Targets | Build targets set via `GOOS`/`GOARCH` env vars inside the Docker image; no toolchain installation beyond the base `casjaysdev/go:latest` image is required |
| Formatting | `gofmt` (or `goimports`) is required; no enforced line length — `gofmt` does not wrap lines |
| Line width | No hard limit; `golangci-lint` `lll` soft guide at 120, hard fail over 180 |
| Linting | `golangci-lint` is required |
| Testing | `go test` is required |
| CGO | `CGO_ENABLED=0` ALWAYS — pure Go, no C (PART 0 → "Go-Only Application") |
| Source language | Go only (PART 0 → "Go-Only Application") |

## Pure-Go Library Stack

This is the recommended starting point for satisfying common application needs with packages that do **not** drag in C dependencies. Pure-Go packages are what make a single static binary that works on Windows, macOS, Linux, and BSD achievable; every C-linked package becomes a portability tax across that target matrix.

**Rule:** when a capability below is needed, prefer the pure-Go option. Only deviate when (a) the pure-Go option is not viable for the project's actual requirements, AND (b) the deviation is documented in `IDEA.md` and `LICENSE.md`. All deviations from pure Go require both an IDEA.md exception and LICENSE.md attribution.

| Capability | Go preferred |
|------------|-------------|
| TLS | `crypto/tls` (standard library) |
| HTTP client | `net/http` (standard library) |
| Crypto | `crypto/*` (standard library); `golang.org/x/crypto` for extras |
| Compression | `compress/*` (standard library) |
| Serialization | `encoding/json`, `encoding/xml` (stdlib); `gopkg.in/yaml.v3`, `github.com/BurntSushi/toml` |
| Async / concurrency | goroutines + channels (language built-in) |
| CLI parsing | `github.com/spf13/cobra` + `github.com/spf13/pflag`; or stdlib `flag` |
| Logging | `log/slog` (stdlib, Go 1.21+); `go.uber.org/zap` |
| Error handling | `fmt.Errorf` + `%w` wrapping; `errors.Is` / `errors.As` (stdlib) |
| Date / time | `time` (standard library) |
| Config files | `github.com/spf13/viper`; or plain TOML/YAML |
| GUI (cross-platform) | `gioui.org` (Gio — pure Go, X11+Wayland+macOS+Windows); `fyne.io/fyne/v2` |
| TUI | `github.com/charmbracelet/bubbletea` + `github.com/charmbracelet/lipgloss` |
| Font rendering | handled by GUI toolkit |
| Image decoding | `image/*` (standard library) |
| Clipboard | `golang.design/x/clipboard` |
| Desktop notifications | `github.com/gen2brain/beeep` |
| File dialogs | `github.com/sqweek/dialog` |
| Theme detection | `github.com/adrg/xdg` + OS theme APIs |
| Secrets storage | `github.com/zalando/go-keyring` |
| Local SQL | `modernc.org/sqlite` (pure Go SQLite — `CGO_ENABLED=0` compatible, no bundled C needed) |
| UUID / hashing | `github.com/google/uuid`; `crypto/sha256` (stdlib) |
| Filesystem watching | `github.com/fsnotify/fsnotify` |
| Process / IPC | `net` (Unix socket / named pipe) from stdlib |
| Networking primitives | `net`, `net/http` from stdlib |

This list is not exhaustive; treat it as the starting point. When introducing a new dependency:

1. Search pkg.go.dev for a pure-Go option first
2. Check for any transitive C dependency (`go mod graph`, review each package's source)
3. If the only viable option has a C dependency, document the exception in `IDEA.md`, confirm it can be statically compiled into the final binary with `CGO_ENABLED=0`, and add attribution to `LICENSE.md`

## Go Commands

**All Go invocations execute inside the project Docker container — never on the host.** The table below shows the *logical* command; the **actual** invocation is wrapped via the `GO_DOCKER` macro (see "Build Metadata" in PART 6 for the full macro definition). See "Docker Rule" below.

| Logical Command | Purpose |
|-----------------|---------|
| `gofmt -l .` | Format verification (list files needing formatting) |
| `gofmt -w .` | Format code in-place |
| `golangci-lint run ./...` | Lint enforcement |
| `go test ./...` | Test suite |
| `go build ./...` | Debug build (all packages) |
| `go build -ldflags="-s -w" -o {project_name} .` | Optimized build |
| `go build -ldflags="-s -w -X main.Version=…" …` | Release build with version injection |
| `go vet ./...` | Vet checks |
| `govulncheck ./...` | Vulnerability scanning |
| `go run . -- [args]` | Execution (in container; with X11/Wayland forwarding if GUI) |

**Examples of the actual wrapped form (using the `GO_DOCKER` macro from the Makefile):**

```bash
# Format check
$(GO_DOCKER) gofmt -l .

# Lint
$(GO_DOCKER) golangci-lint run ./...

# Tests
$(GO_DOCKER) go test ./...

# Release build (artifacts land in binaries/ via the mounted workspace)
$(GO_DOCKER) \
  sh -c "GOOS=linux GOARCH=amd64 go build -buildvcs=false -trimpath \
    -ldflags=\"-s -w -X main.Version=\$(cat release.txt)\" \
    -o binaries/{project_name}-linux-amd64 ./src"

# Run with display forwarding via the `gui` compose service, which mounts
# the X11 / Wayland sockets and exports DISPLAY / WAYLAND_DISPLAY (see
# "Docker Rule" → "X11 and Wayland Forwarding" for socket/env wiring).
docker compose run --rm gui go run ./src -- --ui gui
```

Bare `go …` invocations on the host are forbidden by PART 0 → "No Host Toolchain or Binary Execution."

## Build Rules

- **Pure Go by default** — `CGO_ENABLED=0` always; every dependency is pure Go unless a specific exception is documented in `IDEA.md` (PART 0 → "Go-Only Application")
- Release builds use `go build` with `-ldflags="-s -w"` to strip symbols, reducing binary size
- Use `-ldflags="-X main.Version=…"` (and similar `-X` flags) to inject version, commit ID, and build epoch at compile time (`BuildDate` is derived from `BuildEpoch` at process start, never embedded)
- The final artifact MUST be a single statically linked binary per target (PART 0 → "Single Static Binary")
- Set `GOOS`, `GOARCH`, and `CGO_ENABLED=0` as environment variables when cross-compiling inside the Docker image
- All runtime assets are embedded at compile time via `//go:embed`; the repo's `assets/` directory is build-time input only (PART 0 → "Self-Contained Assets")
- A static-binary self-check runs as part of CI (`ldd` on Linux artifacts must show "not a dynamic executable" or only kernel-vDSO; `otool -L` on macOS must show only Apple-provided frameworks; `dumpbin /dependents` on Windows must show no unexpected DLLs)

## Build Metadata

Inject version, commit ID, build epoch, and official site via `-ldflags`. `BuildDate` is never embedded — it is derived from `BuildEpoch` at process start:

```go
// Build info - BuildEpoch is set via -ldflags at build time; BuildDate is derived from it
var (
    // overridden by: -ldflags="-X main.Version=$(cat release.txt)"
    Version      = "devel"
    // overridden by: -ldflags="-X main.CommitID=$(git rev-parse --short=7 HEAD)"
    CommitID     = "N/A"
    // BuildDate is derived from BuildEpoch in init(); "N/A" when BuildEpoch is unset
    BuildDate    = "N/A"
    // overridden by: -ldflags="-X 'main.BuildEpoch=${BUILD_EPOCH}'" - Unix build timestamp (seconds, UTC)
    BuildEpoch   = "0"
    // overridden by: -ldflags="-X main.OfficialSite=$(cat site.txt)"
    OfficialSite = ""
)

// buildEpoch parses the embedded BuildEpoch ldflag; 0 when unset or invalid.
// Requires "strconv" in this file's import block.
func buildEpoch() int64 {
    n, err := strconv.ParseInt(BuildEpoch, 10, 64)
    if err != nil {
        return 0
    }
    return n
}

// init derives BuildDate (RFC 3339 UTC) from the embedded BuildEpoch
// Requires "time" in this file's import block.
func init() {
    if n := buildEpoch(); n > 0 {
        BuildDate = time.Unix(n, 0).UTC().Format("2006-01-02T15:04:05Z")
    }
}
```

**Full release build command (logical form — run inside the Docker container, not on the host):**
```bash
# BUILD_EPOCH is the single captured time source - captured once per build
BUILD_EPOCH="$(date -u +%s)"

go build \
  -buildvcs=false -trimpath \
  -ldflags="-s -w \
    -X main.Version=$(cat release.txt) \
    -X main.CommitID=$(git rev-parse --short=7 HEAD) \
    -X 'main.BuildEpoch=${BUILD_EPOCH}' \
    -X main.OfficialSite=$(cat site.txt 2>/dev/null || true)" \
  -o binaries/{project_name}-linux-amd64 ./src
```

## Project Layout

```text
go.mod
go.sum
.go-version                 # single-line Go version pin, e.g. "go1.23.0"
src/
├── main.go                 # entry point — auto-detects GUI/TUI/CLI mode
├── app/                    # shared domain/application logic
├── config/                 # configuration loading + defaults
├── platform/               # OS/platform integration
├── state/                  # app state and persistence adapters
├── support/                # logging, errors, utilities
└── ui/
    ├── gui/                # optional native GUI
    ├── tui/                # optional terminal UI
    └── cli/                # optional plain CLI commands/output
assets/                     # BUILD-TIME ONLY: source tree embedded via //go:embed
deps/                       # optional: committed, project-specific support
                             # files not part of build/release output (e.g.
                             # scripts or Dockerfiles for building a
                             # dependency) — never a cache or temp/output dir
docker/                     # REQUIRED: Dockerfile, docker-compose.yml, rootfs/usr/local/bin/entrypoint.sh, README.md
packaging/                  # installer/manifests/bundle metadata
Makefile                    # build targets: build, release, docker, test, dev, clean
scripts/                    # optional host-side Docker wrapper scripts; MUST NOT contain application logic
tests/                      # integration tests
```

The `assets/` directory in the repo holds source files (fonts, icons, default themes, schemas, locales, default config). They are embedded into the binary at compile time via `//go:embed`. The directory is **never** shipped or installed alongside the binary on a user's machine.

## Naming Rules

| Item | Rule | Example |
|------|------|---------|
| Files | `snake_case` or flat lowercase | `config_loader.go`, `platform.go` |
| Package names | lowercase, no underscores | `package config`, `package ui` |
| Types/interfaces | `PascalCase` | `AppConfig`, `UIMode` |
| Functions | `camelCase` (unexported) / `PascalCase` (exported) | `detectUIMode`, `DetectUIMode` |
| Constants/vars | `PascalCase` (exported) / `camelCase` (unexported) (package-level consts) | `DefaultTheme` (exported), `defaultTimeout` (unexported) |
| Go module path (`go.mod`) | lowercase, URL-style | `github.com/myorg/myapp` |
| Go package name | lowercase, no underscores, no hyphens | `package myapp` |

## Release Artifacts

- Primary binary follows the naming schema `{project_name}-{GOOS}-{GOARCH}{.ext}` (Go's native GOOS/GOARCH terms: `linux/darwin/windows/freebsd`, `amd64/arm64`; `.exe` for windows) — single statically linked file, no platform-alias translations
- Each release MUST publish artifacts for at minimum: `{project_name}-linux-amd64`, `{project_name}-linux-arm64`, `{project_name}-darwin-amd64`, `{project_name}-darwin-arm64`, `{project_name}-windows-amd64.exe`, `{project_name}-windows-arm64.exe`, `{project_name}-freebsd-amd64`, `{project_name}-freebsd-arm64` (subset acceptable only when IDEA.md narrows platform scope)
- A static-linkage verification step is part of release: `ldd` / `otool -L` / `dumpbin /dependents` output is captured and checked against an allowlist (kernel vDSO, Apple system frameworks, Windows kernel32/user32 etc.) — anything outside the allowlist fails the release
- No companion files (no `.so`, `.dylib`, `.dll`, no asset bundles, no font directories) ship next to the binary
- Include aggregate `sha256.txt` and `sha512.txt` checksum files covering every published artifact, uploaded as release assets alongside the binaries
- Include release notes that describe actual changes
- Include an SBOM (always — generated via `cyclonedx-gomod`; see PART 10 → "Suggested CI Steps" for the invocation). Include provenance/attestation via `actions/attest-build-provenance` when the release platform supports it; always set `provenance: false` on `docker/build-push-action` steps
- If GUI packaging exists (MSI, DMG, AppImage, deb, rpm, etc.), the package wraps the same single static binary plus desktop integration metadata; package metadata lives in `packaging/`

## Docker Rule

Docker is **REQUIRED**. The project MUST ship a working Docker setup, and every toolchain command and every produced binary runs inside it (PART 0 → "No Host Toolchain or Binary Execution").

### Required Docker Assets

Place Docker assets under `docker/`:

```text
docker/
├── Dockerfile              # production runtime image — two-stage (builder + minimal Alpine); tagged :latest
├── Dockerfile.dev          # devel image — same as release but binary runs in debug mode; tagged :devel   (project-specific)
├── rootfs/                 # build-time filesystem overlay copied into image at /   (project-specific)
│   └── usr/local/bin/entrypoint.sh   # prepares cache dirs; user creation and privilege drop happen in the binary; called by tini → entrypoint.sh → app
├── docker-compose.yml      # production compose — HUMAN USE ONLY
├── docker-compose.dev.yml  # development compose — runs `:devel` image in debug mode; HUMAN USE ONLY
├── docker-compose.test.yml # automated test compose — AI prefers the tests/ scripts over running this directly
└── README.md               # how to build the image, run tests, run GUI with display forwarding
```

`docker-compose.yml` always lives under `docker/` — never at the repo root. See `dockerfile_conventions.md` → "Docker Compose" for the canonical layout: `docker-compose.yml` and `docker-compose.dev.yml` are human-only, and AI's preferred interface for `docker-compose.test.yml` is the project's `tests/` scripts rather than a direct invocation.

### Toolchain Image

**All Go projects use `casjaysdev/go:latest` — never create `docker/Dockerfile.build` or `build-toolchain.yml` for Go.** The maintained image includes everything any Go project needs:

- Latest stable Go toolchain (`go`, `gofmt`, `go vet`); image defaults: `CGO_ENABLED=0`, `GOFLAGS=-buildvcs=false`, `GOTOOLCHAIN=auto`, `GOTELEMETRY=off`
- `golangci-lint`, `staticcheck`, `gofumpt`, `goimports` — linting and formatting
- `govulncheck` — vulnerability scanner
- `go-licenses` — dependency license reporter
- `cyclonedx-gomod` — CycloneDX SBOM generator
- `goreleaser` — release automation
- `gotestsum` — structured test runner
- `gopls` — official Go language server
- `dlv` (Delve) — source-level debugger
- `ko` — build container images from Go source without a Dockerfile
- `air` — live-reload dev server
- `buf` — protobuf toolchain; `protoc-gen-go`, `protoc-gen-go-grpc`
- `goose` — DB migration runner
- `wire` — compile-time dependency injection
- `mockgen` (uber/mock) — interface mock generator
- `stringer` — `String()` method generator for iota types
- `benchstat` — statistically sound benchmark comparison
- `gops` — live Go process diagnostics

CI workflows reference this image directly: `container: image: casjaysdev/go:latest`. No `apk add`, no `go install`, no `ensure-build-image` job, no `build-toolchain.yml`.

### OCI Annotations (No LABEL Policy)

All image metadata is applied as **OCI annotations at build time** — never as `LABEL` blocks in any Dockerfile. Labels attach to per-platform image layers; multiarch manifest indexes do not inherit labels, so they appear missing on multiarch pulls. Annotations attach to the manifest index and are visible across all platforms.

- No `LABEL` blocks in `docker/Dockerfile` or `docker/Dockerfile.dev`
- GitHub Actions: use `docker/metadata-action@030e881283bb7a6894de51c315a6bfe6a94e05cf  # v6.0.0` with an `annotations:` input listing the required OCI keys
- `docker/build-push-action@bcafcacb16a39f128d818304e6c9c0c18556b85f  # v7.1.0` MUST set `annotations: ${{ steps.meta.outputs.annotations }}`, `labels: ""` to suppress label output, AND `provenance: false` to prevent a spurious `unknown/unknown` platform entry in the manifest list (use `actions/attest-build-provenance` for release binary attestation instead)

See `dockerfile_conventions.md` → "OCI Annotations" for the full required annotation set (`org.opencontainers.image.{title,description,url,source,documentation,vendor,authors,vcs-type,version,revision,created,licenses,...}`).

### Build-Time Metadata Args

`docker/Dockerfile` declares `ARG BUILD_EPOCH` in the builder stage and exports it as an environment variable for the `go build` step, so the `-ldflags` `-X 'main.BuildEpoch=…'` injection embeds it (PART 5 → "Build Metadata"). `ARG BUILD_DATE` is declared only where it feeds the `org.opencontainers.image.created` value — it is never passed to the go build. The caller captures `BUILD_EPOCH` once, derives `BUILD_DATE` from it, and passes both:

```bash
# Captured ONCE; BUILD_DATE is derived from it, never independently captured
BUILD_EPOCH="$(date -u +%s)"
BUILD_DATE="$(date -u -d "@${BUILD_EPOCH}" +%Y-%m-%dT%H:%M:%SZ)"
docker build -f docker/Dockerfile \
  --build-arg BUILD_EPOCH="$BUILD_EPOCH" \
  --build-arg BUILD_DATE="$BUILD_DATE" \
  ...
```

### Container Runtime Rules

Every production image MUST satisfy:

- **Startup chain `tini → entrypoint.sh → app`** — `ENTRYPOINT [ "tini", "-p", "SIGTERM", "--", "/usr/local/bin/entrypoint.sh" ]`. Never override `ENTRYPOINT` or `CMD` to bypass `tini` or the entrypoint shim. All startup customization goes in `docker/rootfs/usr/local/bin/entrypoint.sh`, which MUST end with `exec "$@"` to preserve PID 1 signal handling.
- **`STOPSIGNAL SIGTERM`** (or `SIGRTMIN+3` for s6-based images) for graceful shutdown
- **`HEALTHCHECK`** — every production image declares a `HEALTHCHECK` that exits non-zero when the binary is unhealthy
- **Privilege drop, not Dockerfile users** — containers start as root with NO `USER` directive and no user/group creation in the Dockerfile. The binary itself creates its dedicated user/group, creates its directories, sets permissions, then drops privileges once initialization completes. `entrypoint.sh` may export UID/GID env vars so the binary can match host ownership of mounted volumes. Running permanently as root (never dropping) is the exception and MUST be justified in `IDEA.md`.

### Mandatory `docker run` Naming Convention

Every `docker run` invocation in this project (CI, scripts, docs, examples) MUST use:

```bash
PROJECT_NAME="{project_name}"
PROJECT_IMAGE="casjaysdev/go:latest"

docker run --rm \
  --name "${PROJECT_NAME}-$(tr -dc 'a-z0-9' </dev/urandom | head -c8)" \
  ...
```

`PROJECT_NAME` and `PROJECT_IMAGE` are shell variables set once (e.g., in the Makefile or the calling script) before any of the `docker run` examples below — `PROJECT_NAME` resolves to the project's `{project_name}` and `PROJECT_IMAGE` to the toolchain image selected under "Toolchain Image" above (normally `casjaysdev/go:latest`).

- `--rm` — self-remove on exit (no orphaned containers)
- `-it` — interactive-capable for log streaming and signal handling
- `--name "${PROJECT_NAME}-XXXX"` — traceable name; `XXXX` is the 8-char lowercase-alphanumeric random suffix produced by `tr -dc 'a-z0-9' </dev/urandom | head -c8` (Makefile form: `$$(tr -dc 'a-z0-9' </dev/urandom | head -c8)`)

There is no opt-out. Unnamed or persistent build/test containers are a spec violation.

### Portability Rule

No hardcoded org, project name, official site, or registry value may appear in any Dockerfile, workflow, or Makefile. Use build-time variables:

| Context | Reference |
|---------|-----------|
| Dockerfile | `ARG PROJECT_ORG` / `ARG PROJECT_NAME` (passed via `--build-arg`) |
| GitHub Actions | `${{ github.repository_owner }}` / `${{ github.event.repository.name }}` / `${{ github.event.repository.html_url }}` |
| GitLab CI | `$CI_REGISTRY_IMAGE`, `$CI_PROJECT_NAMESPACE`, `$CI_PROJECT_NAME` |
| Gitea / Forgejo | provider-supplied equivalents to the GitHub variables |
| Jenkinsfile | `${env.JOB_NAME}`, `${env.GIT_URL}` (parse org/name) |
| Makefile | `PROJECT_ORG ?= $(shell git remote get-url origin \| ...)` |

This rule ensures every workflow, Dockerfile, and Makefile keeps working after a fork without editing values.

### Build-Time vs Runtime Linkage of Display Libraries

`CGO_ENABLED=0` is the default for Go builds (PART 5 → "Toolchain Image"), so the produced binary is a pure static Go binary with no C link-time dependency of any kind — including on GUI-stack libraries. Display-related C libraries, when needed at all, reach the binary only via runtime `dlopen`/`cgo` opt-in, on demand, when a real GUI session is started; the default build MUST NOT introduce a link-time dependency on them.

- **X11**: prefer a pure-Go X11 client (e.g. `github.com/jezek/xgb` / `github.com/jezek/xgbutil`) — no `libX11` involvement at any stage, and no CGO required. If a CGO-based binding is unavoidable, it MUST be gated behind a build tag that is off by default, and the default release build MUST stay `CGO_ENABLED=0`.
- **Wayland**: prefer a pure-Go Wayland client (e.g. `github.com/rajveermalviya/go-wayland`) that talks the wire protocol directly rather than linking `libwayland-client`.
- **OpenGL / EGL / Vulkan**: when in scope, use Go loaders that resolve symbols at runtime (e.g. `github.com/go-gl/glow`-style dynamic loading) so the binary stays portable across hosts with different GPU stacks; do not statically link `libGL`/`libEGL`/`libVulkan`.
- After every release build, the static-linkage check (`ldd` / `otool -L` / `dumpbin /dependents`, PART 5 → "Release Artifacts") MUST confirm there is no link-time dependency on `libX11` / `libwayland-client` / `libGL` / `libEGL` / `libVulkan`. With `CGO_ENABLED=0` this is automatic; any exception that re-enables CGO for GUI support MUST re-run and document this check.

### CI/CD Workflow Pattern

All Go CI jobs run inside `casjaysdev/go:latest`. Never `apk add` or `go install` in a workflow `run:` step — every tool is already in the image. No `ensure-build-image` pre-flight, no `build-toolchain.yml`.

Canonical job pattern for `ci.yml` / `release.yml`:

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    container:
      image: casjaysdev/go:latest
      options: "--user 0:0"
    steps:
      - uses: actions/checkout@9c091bb21b7c1c1d1991bb908d89e4e9dddfe3e0  # v7.0.0
      - name: Capture build epoch
        run: |
          # Captured ONCE per job - every other time value derives from this
          BUILD_EPOCH=$(date -u +%s)
          echo "BUILD_EPOCH=$BUILD_EPOCH" >> $GITHUB_ENV
      - run: go build -buildvcs=false -trimpath -ldflags="-X 'main.BuildEpoch=${BUILD_EPOCH}'" ./src
```

### X11 and Wayland Forwarding (Mandatory for GUI/Display Testing)

GUI and display-aware test runs use the `gui` compose service (or equivalent `docker run` flags). Both X11 and Wayland forwarding MUST be supported; the spec does not pick one — the container detects what the host provides and forwards accordingly.

**X11 forwarding (host running Xorg or XWayland):**

```bash
# grant access to current user only; revoke when done
xhost +SI:localuser:$(id -un)

docker run --rm \
  --name "${PROJECT_NAME}-$(tr -dc 'a-z0-9' </dev/urandom | head -c8)" \
  -e DISPLAY="$DISPLAY" \
  -e XAUTHORITY=/tmp/.docker.xauth \
  -v /tmp/.X11-unix:/tmp/.X11-unix:rw \
  -v "$XAUTHORITY":/tmp/.docker.xauth:ro \
  --device /dev/dri \
  -v "$PWD":/work -w /work \
  "$PROJECT_IMAGE" go run . -- --ui gui

# revoke after the session
xhost -SI:localuser:$(id -un)
```

**Wayland forwarding (host running a Wayland compositor):**

```bash
docker run --rm \
  --name "${PROJECT_NAME}-$(tr -dc 'a-z0-9' </dev/urandom | head -c8)" \
  -e WAYLAND_DISPLAY="$WAYLAND_DISPLAY" \
  -e XDG_RUNTIME_DIR=/tmp/xdg \
  -v "$XDG_RUNTIME_DIR/$WAYLAND_DISPLAY":"/tmp/xdg/$WAYLAND_DISPLAY" \
  --device /dev/dri \
  -v "$PWD":/work -w /work \
  "$PROJECT_IMAGE" go run . -- --ui gui
```

**Rules:**
- Both forwarding modes MUST be documented and working — pick one based on what the host session exposes
- `xhost +` (wide open) is forbidden; use scoped grants (`+SI:localuser:…`) and revoke after
- Do not run the container as root just to make sockets work; map UID/GID instead
- GPU access (`/dev/dri`) is allowed when the GUI stack benefits from it; keep it opt-in via a compose profile when not always needed
- Sound/MIDI/input device forwarding is allowed only when IDEA.md justifies it
- Audio (PulseAudio/PipeWire) forwarding follows the same scoped-socket pattern when needed

### Forbidden

- Running `go`, `go build`, `go test`, or any built binary on the host
- Treating Docker as a "CI-only" or "release-only" concern
- Recommending `go install` on the host to "just try it quickly"
- `--privileged`, `--net=host` outside of clearly justified, documented profiles
- `xhost +` without a scoped target

---

# PART 6: VERSION, SITE, AND BUILD METADATA

## `release.txt` (Canonical Version When Present)

- `release.txt` is a single-line canonical version source when present
- It overrides derived/tag-based defaults and should stay in sync with `go.mod` module version if applicable
- Use semantic versions unless the project intentionally uses another documented scheme
- If `release.txt` exists, CI/CD and release scripts must use it rather than inventing a version

## `site.txt` (Optional Official Site)

- Optional single-line URL file in project root
- Use it for the official site, homepage, docs root, support page, or update landing page
- Never guess it
- If `site.txt` exists, it wins over `IDEA.md`, README, environment variables, and CI secrets
- If `site.txt` and `IDEA.md` disagree, `site.txt` wins; update IDEA.md to match or remove the stale variable

## Metadata Priority Rules

### Version
1. `release.txt` (when present)
2. `go.mod` module version / git tag
3. explicit documented fallback for bootstrap/dev builds

### Official Site
1. `site.txt` (when present)
2. `IDEA.md ## Project variables` (`official_site`)
3. explicit environment/CI override used only where documented
4. empty

## Build Metadata

Inject at build time via `-ldflags` `-X` flags:
- version
- commit ID
- build epoch (Unix seconds, UTC — the single captured build time)
- official site (optional)

`BuildDate` is not an ldflag — it is derived from `BuildEpoch` at process start (see PART 5 → "Build Metadata" for the `buildEpoch()` helper and `init()` derivation).

**Example variable declarations:**
```go
// Build info - BuildEpoch is set via -ldflags at build time; BuildDate is derived from it
var (
    Version      = "devel"
    CommitID     = "N/A"
    // BuildDate is derived from BuildEpoch in init(); "N/A" when BuildEpoch is unset
    BuildDate    = "N/A"
    // BuildEpoch is the Unix build timestamp (seconds, UTC) set via -ldflags; "0" when unset
    BuildEpoch   = "0"
    OfficialSite = ""
)
```

**Makefile pattern for version injection and Docker build macro:**
```makefile
PROJECTNAME := {project_name}
PROJECTORG  := {project_org}
BINDIR      := binaries
RELDIR      := releases

VERSION     ?= $(shell cat release.txt 2>/dev/null || echo "devel")
COMMIT_ID   := $(shell git rev-parse --short=7 HEAD 2>/dev/null || echo "N/A")
# BUILD_EPOCH is the single captured time source - captured once per build
BUILD_EPOCH := $(shell date -u +%s)
# Derived from BUILD_EPOCH - used only for the Docker OCI created annotation; not an ldflag
BUILD_DATE  := $(shell date -u -d @$(BUILD_EPOCH) +"%Y-%m-%dT%H:%M:%SZ")
SITE        := $(shell cat site.txt 2>/dev/null || true)

# BuildDate is NOT embedded - it is derived from BuildEpoch at process start
LDFLAGS := -s -w \
  -X main.Version=$(VERSION) \
  -X main.CommitID=$(COMMIT_ID) \
  -X 'main.BuildEpoch=$(BUILD_EPOCH)' \
  -X main.OfficialSite=$(SITE)

GO_CACHE  ?= $(HOME)/go/pkg/mod
GO_BUILD  ?= $(HOME)/.cache/go-build/$(PROJECTNAME)

DOCKER_MEM  ?= 4g
DOCKER_CPUS ?= 2

GO_DOCKER := docker run --rm \
	--name $(PROJECTNAME)-$$(tr -dc 'a-z0-9' </dev/urandom | head -c8) \
	--memory=$(DOCKER_MEM) --cpus=$(DOCKER_CPUS) \
	-v $(PWD):/app \
	-v $(GO_CACHE):/usr/local/share/go/pkg/mod \
	-v $(GO_BUILD):/usr/local/share/go/cache \
	-w /app \
	-e CGO_ENABLED=0 \
	-e GOFLAGS=-buildvcs=false \
	casjaysdev/go:latest

PLATFORMS        ?= linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64 freebsd/amd64 freebsd/arm64
DOCKER_PLATFORMS ?= linux/amd64,linux/arm64
REGISTRY         ?= ghcr.io/$(PROJECTORG)/$(PROJECTNAME)

.PHONY: build release docker test dev clean

build:
	@mkdir -p $(BINDIR) $(GO_CACHE) $(GO_BUILD)
	$(GO_DOCKER) go build -buildvcs=false -trimpath -ldflags "$(LDFLAGS)" -o $(BINDIR)/$(PROJECTNAME) ./src

test:
	@mkdir -p $(GO_CACHE) $(GO_BUILD)
	@$(GO_DOCKER) sh -c " \
		mkdir -p \"\$${TMPDIR:-/tmp}/$(PROJECTORG)\" && \
		COVDIR=\$$(mktemp -d \"\$${TMPDIR:-/tmp}/$(PROJECTORG)/$(PROJECTNAME)-XXXXXX\") && \
		go test -v -cover -coverprofile=\$$COVDIR/coverage.out ./... && \
		COVERAGE=\$$(go tool cover -func=\$$COVDIR/coverage.out | grep total | awk '{print \$$3}' | sed 's/%//') && \
		echo \"Coverage: \$$COVERAGE%\" && \
		if [ \$$(echo \"\$$COVERAGE < 60\" | bc -l) -eq 1 ]; then \
			echo \"ERROR: Coverage is \$$COVERAGE%, must be >= 60%\"; exit 1; \
		fi && \
		echo \"Tests complete - Coverage: \$$COVERAGE% (>= 60% required) ✓\""

release:
	@mkdir -p $(BINDIR) $(RELDIR) $(GO_CACHE) $(GO_BUILD)
	@for platform in $(PLATFORMS); do \
		OS=$${platform%/*}; ARCH=$${platform#*/}; \
		OUTPUT=$(BINDIR)/$(PROJECTNAME)-$$OS-$$ARCH; \
		[ "$$OS" = "windows" ] && OUTPUT=$$OUTPUT.exe; \
		echo "Building $$OS/$$ARCH..."; \
		$(GO_DOCKER) sh -c "GOOS=$$OS GOARCH=$$ARCH \
			go build -buildvcs=false -trimpath -ldflags \"$(LDFLAGS)\" \
			-o $$OUTPUT ./src" || exit 1; \
	done
	@cp $(BINDIR)/$(PROJECTNAME)-* $(RELDIR)/
	@echo "$(VERSION)" > $(RELDIR)/version.txt
	@cd $(RELDIR) && FILES="$$(ls)" && sha256sum $$FILES > sha256.txt && sha512sum $$FILES > sha512.txt

docker:
	docker buildx build \
		--platform $(DOCKER_PLATFORMS) \
		--build-arg VERSION=$(VERSION) \
		--build-arg COMMIT_ID=$(COMMIT_ID) \
		--build-arg BUILD_DATE=$(BUILD_DATE) \
		--build-arg BUILD_EPOCH=$(BUILD_EPOCH) \
		-t $(REGISTRY):$(VERSION) -t $(REGISTRY):latest \
		-f docker/Dockerfile .

dev: build
	docker run --rm -it \
		--name $(PROJECTNAME)-dev-$$(tr -dc 'a-z0-9' </dev/urandom | head -c8) \
		-v $(PWD):/app -w /app \
		casjaysdev/go:latest \
		$(BINDIR)/$(PROJECTNAME)

clean:
	rm -rf $(BINDIR) $(RELDIR)
```

---

# PART 7: RUNTIME DETECTION, CONFIGURATION, AND OUTPUT

## Runtime Detection Rules

All machine-dependent settings MUST be detected at runtime on the target machine, never copied from the developer machine.

| Setting | Detection Method | NEVER Do |
|---------|------------------|----------|
| OS / architecture | `runtime.GOOS`, `runtime.GOARCH` | hardcode dev machine values |
| Home/config/data dirs | `os.UserHomeDir()`, `os.UserConfigDir()`, `os.UserCacheDir()` | hardcode absolute dev paths |
| Locale | `LANG`, platform locale APIs | assume English-only unless documented |
| Theme preference | OS theme APIs / config | assume dark/light universally |
| Terminal capability | TTY + TERM + negotiated features | assume ANSI/alt-screen support |
| Display stack | X11 (`DISPLAY`) and Wayland (`WAYLAND_DISPLAY`) detection — both backends supported, runtime-selected by GUI toolkit | assume GUI exists; ship X11-only or Wayland-only |
| CPU / memory | `runtime.NumCPU()`, runtime detection if needed | tune only for dev hardware |

## Configuration Rules

- Config files store user choices, not auto-detected machine facts
- Defaults must be sensible for a fresh per-user install
- Support config file + environment variable + CLI override layering
- CLI override wins over env; env wins over config; config wins over defaults
- Secrets must not be stored in world-readable files

## Logging & Log Rotation

Applications write `app.log` and `error.log` to the platform log directory. Rotation is built in — no external logrotate needed. Same `rotate`/`keep` schema as the SERVER specification.

### Rotation Options

| Option | Description |
|--------|-------------|
| `never` | Never rotate |
| `daily` | Rotate daily |
| `weekly` | Rotate weekly |
| `monthly` | Rotate monthly |
| `yearly` | Rotate yearly |
| `NMB` | Rotate at N megabytes (e.g., `50MB`) |
| `NGB` | Rotate at N gigabytes (e.g., `1GB`) |
| Combined | Time + size, whichever first (e.g., `weekly,50MB`) |

### Retention Options

| Option | Description |
|--------|-------------|
| `none` | Do not keep old logs (delete after rotation) |
| `N` | Keep N old log files |
| `Nd` | Keep logs for N days |
| `Nw` | Keep logs for N weeks |
| `Nm` | Keep logs for N months |
| `forever` | Keep forever (no automatic deletion) |

### Configuration

```yaml
logging:
  # Global log level: debug, info, warn, error
  level: warn

  # All log types share these options:
  #   filename: name of log file
  #   format: output format (text, json)
  #   rotate: rotation policy
  #   keep: retention policy
  #   compress: compress rotated logs (only useful if keep > 0)

  app:
    filename: app.log
    format: text
    rotate: weekly,50MB
    keep: none
    compress: false

  error:
    filename: error.log
    format: text
    rotate: weekly,50MB
    keep: none
    compress: false
```

### Defaults

| Log Type | Rotation | Keep |
|----------|----------|------|
| app.log | weekly,50MB | none |
| error.log | weekly,50MB | none |

- `weekly,50MB` = rotate on weekly OR 50MB, whichever comes first
- `keep: none` = do not retain old logs (default) — deleted immediately after rotation

## Standard CLI Flags

### Universal Flags (ALL Binaries)

| Flag | Short | Values | Description |
|------|-------|--------|-------------|
| `--help` | `-h` | — | Show help and exit 0 |
| `--version` | `-v` | — | Show version and exit 0 |
| `--debug` | — | — | Enable debug output |
| `--color` | — | `auto` (default) / `yes` / `no` | Color output — `auto`: TTY detect; `yes`: force on; `no`: force off and remove emojis |

- `--color auto` and `--color=auto` (space and `=` forms) must both work — stdlib `flag` and `cobra`/`pflag` both handle this natively; no extra parsing needed either way.
- **No escalation** — `--help` and `--version` must never be gated behind privilege checks; never call `sudo`, require root/admin, or check `os.Getuid()` before help/version output.

## Output Rules

- Respect `NO_COLOR`
- Respect `TERM=dumb`
- Respect non-TTY stdout/stderr
- Provide machine-readable output when the app advertises it (`--json`, `--plain`, etc.)
- Help/version output must show the **actual invoked binary name**
- **No escalation** — help at every level (main, subcommand, nested) must never call `sudo`, require root/admin, or check privilege state; exit immediately with the help text.

## Human-Readable Values (User-Facing Output)

**Every value shown on a user-facing surface — GUI windows and dialogs, TUI screens, and pretty CLI output — MUST be human-readable. Raw machine values belong to machine-readable output (`--json`, `--plain`) and log files only.**

| Kind | Rule | Examples |
|------|------|----------|
| **Durations** | Largest fitting unit, at most two units, correct singular/plural: <60 s → seconds · ≥60 s → minutes · ≥60 min → hours · ≥24 h → days | `1 second` · `45 seconds` · `3 minutes` · `2 minutes 5 seconds` · `2 hours` · `1 hour 30 minutes` · `3 days 4 hours` |
| **Sizes** | 1024 boundaries, full unit names, singular/plural, at most one decimal (drop `.0`): bytes → kilobytes → megabytes → gigabytes → terabytes | `1 byte` · `512 bytes` · `1 kilobyte` · `2.5 megabytes` · `5 gigabytes` · `1.2 terabytes` |
| **Counts** | Locale-aware thousands separators | `12,847` |
| **Timestamps** | Display format `%B %d, %Y at %H:%M:%S %Z` (zero-padded day) | `January 05, 2026 at 14:03:07 UTC` |

| Rule | Detail |
|------|--------|
| **Shared helpers** | One implementation: `format.Duration()` / `format.Size()` / `format.Count()` in the shared/common formatting package — never per-surface ad-hoc formatting |
| **i18n** | Unit names go through translation keys (`format.seconds_one`, `format.seconds_other`, …) with per-language plural rules — never hardcoded English unit strings |
| **Raw value preserved** | GUI/TUI surfaces MAY carry the machine value in a tooltip or detail view; the visible text is always the human form |
| **Machine surfaces unchanged** | `--json`/`--plain` output fields and log files keep raw base units (seconds, bytes) — formatting is a presentation concern only |

## NO_COLOR Support

When `NO_COLOR` is set and non-empty, disable ANSI color output. If the TUI depends on richer formatting, fall back to CLI/plain output instead of forcing a degraded pseudo-TUI.

| Condition | GUI | TUI | CLI |
|-----------|-----|-----|-----|
| local desktop with display | allowed | fallback | fallback |
| SSH with TTY | no auto-GUI | allowed if terminal capable | fallback |
| `TERM=dumb` | no auto-GUI | avoid | use CLI |
| `NO_COLOR=1` | GUI unaffected unless project says otherwise | avoid color-dependent TUI | prefer CLI/plain |
| stdout piped | avoid GUI | avoid TUI | use CLI |

### Color Enablement Precedence

```go
// ColorEnabled checks if color output should be used
func ColorEnabled(forceColor *bool) bool {
    // 1. CLI flag overrides everything
    if forceColor != nil {
        return *forceColor
    }
    // 2. Config file (if applicable)
    if cfg := GetConfig(); cfg != nil && cfg.Output.ColorSet {
        return cfg.Output.Color
    }
    // 3. NO_COLOR env var (non-empty = disable)
    if os.Getenv("NO_COLOR") != "" {
        return false
    }
    // 4. Auto-detect: TTY + TERM support
    if !term.IsTerminal(int(os.Stdout.Fd())) {
        return false
    }
    if os.Getenv("TERM") == "dumb" {
        return false
    }
    return true
}
```

CLI and TUI output MUST gate on `ColorEnabled(nil)` — never a separate ad
hoc `NO_COLOR` check. `lipgloss`/`termenv` (TUI) auto-detect `NO_COLOR`
on their own; raw `fmt.Printf` ANSI escapes in the CLI path are not
NO_COLOR-aware by themselves and must check `ColorEnabled()` explicitly.

## Color Palette (TUI/CLI/GUI)

**This is a single native binary — there is no Web CSS palette. TUI/CLI
and GUI theming are defined separately, and neither uses literal hex:**

### TUI/CLI — ANSI-mapped

Terminals render a fixed, user-configured 16/256-color set, so TUI/CLI
map semantic roles to the nearest ANSI color instead of literal hex:

| Role | Dark ANSI | Light ANSI |
|------|-----------|------------|
| `Foreground` | `BrightWhite` | `Black` |
| `Muted` | `White` | `DarkGray` |
| `Primary` / `Accent` | `BrightMagenta` | `Blue` |
| `Secondary` / `Success` | `BrightGreen` | `Green` |
| `Warning` | `BrightYellow` | `Yellow` |
| `Error` | `BrightRed` | `Red` |
| `Info` | `BrightBlue` | `Blue` |

```go
// TerminalPalette holds ANSI 16-color indices (0-15) for TUI/CLI.
// lipgloss.Color() and the ESC[38;5;{n}m escape both accept these
// indices directly.
type TerminalPalette struct {
    Foreground string `json:"foreground"`
    Muted      string `json:"muted"`
    Primary    string `json:"primary"`
    Success    string `json:"success"`
    Warning    string `json:"warning"`
    Error      string `json:"error"`
    Info       string `json:"info"`
    Border     string `json:"border"`
}

var TerminalPaletteDark = TerminalPalette{
    Foreground: "15", Muted: "7", Primary: "13",
    Success: "10", Warning: "11", Error: "9", Info: "12", Border: "13",
}

var TerminalPaletteLight = TerminalPalette{
    Foreground: "0", Muted: "8", Primary: "4",
    Success: "2", Warning: "3", Error: "1", Info: "4", Border: "4",
}
```

### GUI — native theming only

GUI never consumes `TerminalPalette` or any literal hex palette. It
detects light/dark only (`github.com/adrg/xdg` / OS theme APIs — see
"Theme detection" above) and lets the native toolkit (GTK/Cocoa/Win32)
apply its own light/dark widget theme.

---

## Reuse Before Creating

**Before writing new code for anything — a function, a variable/constant, or a UI component/style — check whether an equivalent already exists in the project and reuse or extend it. Only create something new when nothing existing covers the need.** This is a general project-wide rule; it governs every artifact type and is not restricted to any one feature.

### Functions

Before writing a new function, search for an existing one with the same or similar behavior — in the same package, in `helpers.go`, in existing handlers/validators/detection logic (e.g. the UI-mode detection and theme-detection functions above) — and call or extend it instead of re-implementing the logic. Two near-identical functions that differ only in a hardcoded value are a sign the existing function should take that value as a parameter instead of being copy-pasted.

### Variables & Constants

Before adding a new constant or config key, check existing constants (e.g. `UIModeTUI`/`UIModeGUI`/`UIModeCLI`) and the config schema for one that already means the same thing. Two names for the same underlying value is a bug waiting to happen, not two separate settings.

### Components

Before building a new TUI view, GUI widget, or CLI output helper, check for an existing one and reuse or extend it instead of building a near-duplicate with a different name.

### Styling (TUI/CLI/GUI — not CSS)

**This is a single native binary — there is no Web CSS to reuse.** Styling reuse instead means: never invent a second color palette, a second set of semantic role names, or a second theme-detection mechanism alongside the ones already defined above.

- **TUI/CLI:** always style through the existing `TerminalPalette` struct and its semantic roles (`Foreground`, `Muted`, `Primary`, `Success`, `Warning`, `Error`, `Info`, `Border`) and the existing `TerminalPaletteDark`/`TerminalPaletteLight` instances — never a new hardcoded ANSI index or a second palette struct. New TUI/CLI output reuses existing `lipgloss` styles built from the palette; a genuinely new style still derives from a `TerminalPalette` role, never a literal color value.
- **GUI:** never invent a custom color palette or literal hex values. Reuse the existing native-toolkit theming path (`gioui.org` / `fyne.io/fyne/v2`, OS light/dark detection) so widgets automatically match the user's OS theme — the same "no literal hex, detect and defer to the platform" rule already established in Color Palette (TUI/CLI/GUI) above.

```go
// CORRECT — new TUI element styled from the existing palette, no new
// hardcoded ANSI index or duplicate palette struct
style := lipgloss.NewStyle().Foreground(lipgloss.Color(palette.Warning))
```

```go
// WRONG — hardcodes a new ANSI index instead of reusing the existing
// TerminalPalette.Warning role
style := lipgloss.NewStyle().Foreground(lipgloss.Color("11"))
```

This rule governs the entire spec — it is not restricted to the Color Palette section, and it is not restricted to whatever the reader might assume is the "main" UI surface (TUI vs. GUI vs. CLI). Every section describing a function, variable, or UI surface inherits this rule automatically; sections do not need to restate it.

---

# PART 8: TESTING, QUALITY, AND DEBUGGING

## Required Quality Gates

All gates execute inside the project Docker container (PART 5 → "Docker Rule"). The "Logical Command" column shows the Go invocation; the actual command is Docker-wrapped.

| Gate | Logical Command | Wrapped Form (example) |
|------|-----------------|------------------------|
| Formatting | `gofmt -l .` | `$(GO_DOCKER) gofmt -l .` |
| Linting | `golangci-lint run ./...` | `$(GO_DOCKER) golangci-lint run ./...` |
| Tests | `go test ./...` | `$(GO_DOCKER) go test ./...` |
| Vet | `go vet ./...` | `$(GO_DOCKER) go vet ./...` |
| Vulnerability scan | `govulncheck ./...` | `$(GO_DOCKER) govulncheck ./...` |
| License enumeration | `go-licenses report ./...` | see PART 10 → "Suggested CI Steps" |
| Attribution drift | `go-licenses report ./...` (output diffed against the GENERATED region of `LICENSE.md`) | see PART 10 → "Suggested CI Steps" |
| Coverage | `go test -cover ./...` | see "Coverage Gate" below |
| GUI smoke (X11) | `go run . -- --ui gui` against an X11 socket | see PART 5 → "X11 and Wayland Forwarding" |
| GUI smoke (Wayland) | `go run . -- --ui gui` against a Wayland socket | see PART 5 → "X11 and Wayland Forwarding" |

## Coverage Gate

`ci.yml` MUST enforce a minimum test coverage threshold. The threshold is declared in `IDEA.md ## Business logic` (free-form prose — e.g., "minimum test coverage: 75%"). If `IDEA.md` does not specify a value, the **default is 60%**. The `test:` Makefile target already enforces a hardcoded 60% floor (PART 6 → "Build Metadata"); when `IDEA.md` declares a different threshold, update that hardcoded value to match — the Makefile check and the CI gate below must never disagree.

- Use `go test -cover ./...` to compute per-package coverage and `go tool cover -func=coverage.out` to get a total percentage — both pre-installed in `casjaysdev/go:latest`
- CI MUST fail when coverage drops below the threshold; a passing build with uncovered code is a silent regression
- Coverage is computed against `go test ./...` output, run inside the toolchain image

```bash
# Example (Docker-wrapped) — fails if total coverage < threshold
THRESHOLD="${COVERAGE_MIN:-60}"
docker run --rm \
  --name "${PROJECT_NAME}-$(tr -dc 'a-z0-9' </dev/urandom | head -c8)" \
  -v "$PWD":/work -w /work "$PROJECT_IMAGE" \
  sh -c "go test -coverprofile=coverage.out ./... && \
    COVERAGE=\$(go tool cover -func=coverage.out | tail -1 | awk '{print \$3}' | tr -d '%') && \
    awk -v c=\"\$COVERAGE\" -v t=\"$THRESHOLD\" 'BEGIN { exit !(c >= t) }'"
```

## Testing Rules

- Core business logic must have unit tests
- Integration tests go in `tests/`
- UI-specific code should keep as much logic as possible outside the UI layer so it is unit-testable
- GUI/TUI smoke tests are required when those surfaces exist
- Snapshot/golden tests are allowed for stable text output and rendering state, but must be intentionally reviewed
- Fixes must include a test when the behavior is testable

## Debug Rules

- Release builds default to user-friendly error messages
- Debug/development mode may include verbose logs and stack traces
- Do not expose sensitive data in logs
- Panic behavior must be intentional and documented (`recover` used deliberately, not to swallow all errors)
- GUI/TUI debug tooling must not leak into normal production UX by default

**Standard exit codes (MUST):** `0` success · `1` general/runtime error (includes a recovered panic) · `2` usage error (bad flags/args — matches Go's own `flag` package behavior). `--help`/`--version` always exit `0` (see "Standard CLI Flags"). Never invent additional exit codes without documenting them in `--help` output.

**Top-level panic recovery (MUST, `main()`):** every binary's `main()` MUST install a top-level `recover` so an unhandled panic prints a clean, user-facing error on stderr and exits `1` instead of dumping a raw Go stack trace to a normal user — the stack trace is only for `--debug`/development mode (see above).

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            if debugMode {
                fmt.Fprintf(os.Stderr, "panic: %v\n\n%s\n", r, debug.Stack())
            } else {
                fmt.Fprintf(os.Stderr, "error: %v\n", r)
            }
            os.Exit(1)
        }
    }()
    run()
}
```

**Signal handling outside daemon mode (MUST):** a plain CLI/GUI/TUI binary (i.e., not running in PART 14's RFC daemon mode) MUST still install `signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)` around its main operation so `Ctrl+C`/`SIGTERM` during a long-running command (a download, a build, an IPC session) cancels the context, lets in-flight cleanup (temp files, open IPC sockets/handles, partial output) run, and then exits `1` — never leave the process to die mid-write via Go's default signal behavior. Daemon mode's own richer signal contract (SIGHUP reload, graceful drain) is defined separately in PART 14 and takes precedence when that mode is active.

## Directory Naming

**Singular** — Go source directories match package names, and Go package names are singular by convention (`internal/handler/`, `internal/model/`, `internal/middleware/`) — not `handlers/`, `models/`. Tooling dirs are always plural regardless of language (`scripts/`, `tests/`, `completions/`).

## Performance Rules

- Optimize based on measured behavior, not guesses
- Scale caches/worker counts from `runtime.NumCPU()` or documented defaults when relevant
- Keep minimal-system defaults sensible
- Avoid unnecessary allocations and duplicate state copies in hot paths
- Never spawn unbounded goroutines — always cap with a semaphore, worker pool, or context cancellation

---

# PART 9: SECURITY & PRIVACY

## Security-First Design

- Least privilege by default
- Per-user storage by default
- Explicit consent for networked, destructive, or privileged operations
- Validate all untrusted input
- Prefer pure-Go packages and audited dependencies
- Use `crypto/tls` (standard library) for TLS on all platforms. On macOS and Windows, the platform-native TLS APIs are also accessible through Go's standard library — these are acceptable. On Linux/BSD, `crypto/tls` covers all use cases without any C dependency
- For in-memory secret handling, avoid holding secrets in plain `string` or `[]byte` longer than necessary; zero memory after use where practical (`defer runtime.KeepAlive` patterns, explicit zeroing before `return`)
- For OS-level secret storage, use `github.com/zalando/go-keyring` (pure Go; speaks the platform protocol on Linux/macOS/Windows). Do not link to `libsecret` / `libgnome-keyring` / `kwallet` directly — those would require CGO and violate the `CGO_ENABLED=0` rule

## Secrets and Sensitive Data

- Never log raw secrets
- Redact tokens, passwords, API keys, auth headers, and personal data from logs and diagnostics
- Config files containing secrets must use restrictive permissions
- Export/import features must clearly label sensitive content

## Update / Download Safety

If the app downloads updates or remote content:

- require explicit user intent or documented auto-update policy
- verify integrity/signatures where supported
- never run downloaded code implicitly
- document trust assumptions in IDEA.md

Plugin downloads are an additional case and apply only when IDEA.md defines a hardened plugin contract per PART 0 → "Single Static Binary"; the four rules above apply to them as well.

## Dependency Governance

- Keep dependencies minimal
- Remove unused packages promptly
- **Renovate is the only supported dependency-update tool** — covers Go deps, GitHub Actions SHAs, and Docker image digests from a single `renovate.json` at the repo root. Works on GitHub, GitLab, Gitea, Forgejo, and Jenkins.
- **Dependabot is forbidden** — GitHub-only, duplicates Renovate's work on GitHub, and cannot serve the other four providers. Never enable both.
- Public repos MUST ship `renovate.json` so Go / Actions / Docker updates land as PRs automatically; Renovate uses `pinDigests: true` to keep all `uses:` lines pinned to immutable SHAs
- Renovate only updates the SHA; the **runtime-still-supported** verification (e.g., node24 vs deprecated runtimes) remains a manual check on every SHA bump (PART 10 → "Third-party Action Pinning")
- Security advisories are blockers until triaged
- Run `govulncheck ./...` as part of CI to catch known vulnerabilities

## Telemetry Rule

No hidden telemetry. Any analytics, crash reporting, or update pings must be documented and controllable.

---

# PART 10: CI/CD, RELEASES, AND AUTOMATION

## CI/CD Rules

| Rule | Description |
|------|-------------|
| Explicit commands | CI must run explicit `go` / `make` commands, not hide core release logic behind undocumented wrappers |
| Least privilege | CI tokens/permissions default to read-only |
| Pinned actions | Third-party actions pinned to full commit SHAs; verify runtime and maintenance status on every SHA update |
| No unsafe fork secrets | Fork PRs do not receive secrets or publish permissions |
| Version precedence | `release.txt` wins when present |
| Site precedence | `site.txt` wins when present |
| Verifiable outputs | Releases publish aggregate `sha256.txt` / `sha512.txt` checksum files and an SBOM (always); provenance/attestation when the platform supports it |
| Single build time source | Every build job — on every provider (GitHub Actions, Gitea/Forgejo, GitLab CI, Jenkins) — captures `BUILD_EPOCH="$(date -u +%s)"` exactly once and exports it to the build environment so the `go build` `-ldflags` `-X 'main.BuildEpoch=…'` injection embeds it (PART 6). `BUILD_DATE` is never independently captured; where a Docker build needs the OCI `image.created` label, derive it (`date -u -d "@${BUILD_EPOCH}" +%Y-%m-%dT%H:%M:%SZ`) and pass both via `--build-arg BUILD_EPOCH=` / `--build-arg BUILD_DATE=` |
| No Makefile in CI | Workflow `run:` steps invoke explicit commands with all environment variables inlined — never `make {target}`. The Makefile is for local developer convenience only. CI MUST NOT depend on Makefile targets that could drift silently. |
| Portability | No hardcoded org, project name, official site, or registry value anywhere in workflows. Use `${{ github.repository_owner }}` / `${{ github.event.repository.name }}` (and provider equivalents). Workflows must keep working after a fork without editing values. |
| Renovate only | `renovate.json` at repo root is the only supported dependency-update tool — it covers GitHub Actions SHAs, Docker image digests, Go module versions, and works across all five providers from a single config. Dependabot is **forbidden** (GitHub-only; duplicates Renovate on GitHub; cannot serve the other four providers). |
| `act` pre-commit validation | Before committing any change to `.github/workflows/*.yml`, run `act --list -W {file}` on each changed file. Fix all errors before committing. The `validate-workflows.sh` PreToolUse hook enforces this automatically. |
| Concurrency groups | Every push/PR workflow declares `concurrency: { group: ${{ github.workflow }}-${{ github.ref }}, cancel-in-progress: true }`. Release workflows also use `cancel-in-progress: true` — a newer tag push supersedes the in-flight release build. |
| Artifact retention | Every `actions/upload-artifact` step sets `retention-days: 7` (or shorter) — no infinite retention of build outputs. |

## Workflow Permissions

Set `contents: read` at the workflow level as the read-only baseline. Grant write permissions only on the specific job that performs the release or publish step — never workflow-wide.

| Permission | Scope | Why |
|------------|-------|-----|
| `contents: read` | All jobs (baseline) | Checkout |
| `contents: write` | Release job only | Create GitHub release, upload assets |
| `packages: write` | Release job only | Push images to `ghcr.io` |
| `id-token: write` | Release job only | OIDC token for Sigstore/cosign artifact signing |
| `attestations: write` | Release job only | GitHub artifact attestation (SBOM, provenance) |

```yaml
# Workflow-level: read-only baseline
permissions:
  contents: read

jobs:
  build:
    # Inherits read-only — no overrides needed
    runs-on: ubuntu-latest
    ...

  release:
    needs: build
    permissions:
      # create GitHub release + upload assets
      contents: write
      # push to ghcr.io
      packages: write
      # OIDC token for cosign signing
      id-token: write
      # GitHub artifact attestations (SBOM, provenance)
      attestations: write
    ...
```

Third-party registry publishing uses repository secrets, not GitHub token permissions (e.g. a private Go module proxy auth token for a non-public `GOPROXY`).

## Third-party Action Pinning

Every external action (`uses: owner/action@...`) MUST be pinned to a full commit SHA — never a mutable tag or branch:

```yaml
# Wrong — tag can silently change or be deleted
- uses: actions/checkout@v4

# Correct — SHA is immutable
- uses: actions/checkout@9c091bb21b7c1c1d1991bb908d89e4e9dddfe3e0  # v7.0.0
```

**When updating a pinned SHA**, verify three things:

1. **Action is still maintained** — check the upstream repo is not archived, deprecated, or abandoned
2. **Runtime is still supported** — open the action's `action.yml` at the new SHA and check `runs.using`; if it names a runtime that GitHub has deprecated or scheduled for removal, the action will silently fail after that date. Example: `node20` is removed from GitHub-hosted runners on **2026-09-16** — any action still on `node20` must be updated to a SHA where it has migrated to `node24` — all common `actions/*` and `docker/*` actions have already done so
3. **No supply-chain change** — skim the diff between the old and new SHA; unexpected new dependencies, changed entrypoints, or network calls added to setup steps are red flags

**Renovate** (`renovate.json`) is the only supported dependency update tool — it updates GitHub Actions SHAs, Docker image digests, and Go module versions across all providers in a single config (see `cicd_conventions.md`). Renovate does not do the runtime verification — that is always a manual check.

### Provider CLIs and Local Runner

**Internet access is available and must be used** — never say "I don't have internet access". Fetch docs, versions, READMEs, and any fact that changes over time rather than guessing from training data.

**User-provided URLs:** always fetch with `\curl -q -LSs {url}` — never WebFetch. WebFetch is for AI-initiated lookups only.

**Tool preference for lookups:** `WebSearch` (open-ended query) → `WebFetch` (known URL, no pipe) → `\curl -q -LSs` (pipe to `jq`/`grep`/file) → provider CLI for provider API ops.

**Provider CLIs** (prefer over raw `curl` when installed):
- `gh` — GitHub (Apache-2.0)
- `glab` — GitLab (MIT)
- `tea` — Gitea and Forgejo (MIT, compatible API)

**`act`** (nektos/act, MIT) — run GitHub Actions locally with `act -j {job}` to verify before pushing. Not a CI replacement — always let real CI run.

## Minimum Public Repo Workflows

Every project ships workflow files for all five CI/CD providers. Same gates, different syntax — no vendor lock-in.

**Workflow creation order — not all workflows carry the same risk:**
1. **Security-only workflows** (secret scan, SHA/digest policy, dependency audit) — no build dependency; safe to add anytime
2. **`ci.yml` and `release.yml`** — add **last**, only after all code is complete, `make test` passes, and the lint gate is clean; these trigger a full build on push and will fail immediately if the code is not ready

Go projects never have `build-toolchain.yml` — `casjaysdev/go:latest` is maintained externally and needs no per-project rebuild workflow.

| Provider | Workflow location | Syntax |
|----------|------------------|--------|
| GitHub | `.github/workflows/ci.yml` / `release.yml` | GitHub Actions |
| GitLab | `.gitlab-ci.yml` (stages: build, test, security, release) | GitLab CI |
| Gitea | `.gitea/workflows/ci.yml` / `release.yml` | GitHub Actions (act runner) |
| Forgejo | `.forgejo/workflows/ci.yml` / `release.yml` | GitHub Actions (act runner) |
| Jenkins | `Jenkinsfile` (parallel stages: Build / Test / Security / Release) | Declarative Pipeline |

**Workflow purposes:**

| Workflow | Trigger | Purpose |
|----------|---------|---------|
| `ci.yml` | push + pull_request + weekly schedule (security jobs only on schedule) | `gofmt -l .`, `golangci-lint run`, `go test ./...`, coverage gate, `go build`, truffleHog secret scan, workflow-policy SHA check, `govulncheck ./...`, Trivy image scan |
| `release.yml` | tag push (`v*`) + workflow_dispatch | Build statically linked artifacts for the supported target matrix, sign, upload to GitHub Releases, publish SBOM and checksums |

In addition to the workflows, every repository ships:

- `renovate.json` — single dependency-update config covering GitHub Actions SHAs, Docker digests, Go module versions; works on all five providers (see PART 9 → "Dependency Governance")

**`security` job conditionality (applies to all providers):**
- Secret scan (truffleHog) — always runs; full git history required
- Workflow policy (SHA/digest pinning check) — always runs
- `vuln-scan` (`govulncheck`) — conditional on `go.sum` present
- `image-scan` (Trivy) — conditional on Dockerfile present; runs after image build

**Container-job user rule:** every `container:` job MUST set `options: "--user 0:0"`. The runner (and actions/checkout's post-job cleanup) execs into the job container — e.g. `cat /etc/*release` for OS diagnostics — as a user the image's `/etc/passwd` may not define, which fails or flakes the job after all real work already passed, wasting the entire run. Numeric `0:0` needs no `/etc/passwd` lookup at all, so it is immune regardless of the image's user table.

**GitHub Actions job ordering (`needs:`):**

- `ci.yml`: `lint` and `test` run in parallel → `build` needs: test → `upload-artifacts` needs: build; security jobs (`secret-scan`, `workflow-policy`, `vuln-scan`, `image-scan`) run in parallel with each other
- `release.yml`: `build` → `release` (needs: build); release job always re-runs its own build inline
- Cross-workflow ordering via branch protection; never `workflow_run`

**GitLab CI**: security jobs run in the `security` stage (parallel by default). Release stage triggered by `$CI_COMMIT_TAG`.

**Jenkins**: `Security` stage uses `parallel {}`. `Release` stage uses `when { tag 'v*' }`.

CI MUST fail on all providers when tests, coverage gates, secret scans, dependency checks, or release validation fail. Never accept a weaker gate on one provider than another.

### Security Jobs in `ci.yml` Example (GitHub / Gitea / Forgejo)

truffleHog is the mandatory secret scanner — **never** `gitleaks` (requires a commercial license for org repos). Trivy is the mandatory image scanner. Both pin to full commit SHAs. These jobs live in `ci.yml` and also run on the weekly schedule trigger.

```yaml
name: CI

on:
  push:
  pull_request:
  schedule:
    # weekly Monday 06:00 UTC
    - cron: '0 6 * * 1'

permissions:
  contents: read

concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}-${{ github.event_name }}
  cancel-in-progress: true

jobs:
  secret-scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@9c091bb21b7c1c1d1991bb908d89e4e9dddfe3e0  # v7.0.0
        with:
          # required: truffleHog needs full history
          fetch-depth: 0

      # Empty base/head = full-history scan (schedule runs and new-branch pushes)
      - name: Determine scan range
        id: range
        run: |
          BASE=""; HEAD=""
          if [ "${{ github.event_name }}" = "push" ] && [ "${{ github.event.before }}" != "0000000000000000000000000000000000000000" ]; then
            BASE="${{ github.event.before }}"; HEAD="${{ github.sha }}"
          elif [ "${{ github.event_name }}" = "pull_request" ]; then
            BASE="${{ github.event.pull_request.base.sha }}"; HEAD="${{ github.event.pull_request.head.sha }}"
          fi
          echo "base=$BASE" >> "$GITHUB_OUTPUT"
          echo "head=$HEAD" >> "$GITHUB_OUTPUT"

      - name: TruffleHog secret scan
        uses: trufflesecurity/trufflehog@27b0417c16317ca9a472a9a8092acce143b49c55  # v3.95.9
        with:
          # NEVER use default_branch — it resolves to HEAD post-push and skips the scan
          base: ${{ steps.range.outputs.base }}
          head: ${{ steps.range.outputs.head }}
          extra_args: --results=verified,unknown

  workflow-policy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@9c091bb21b7c1c1d1991bb908d89e4e9dddfe3e0  # v7.0.0
      - name: Verify all third-party actions are pinned to a 40-char SHA
        run: |
          set -eo pipefail
          bad=$(grep -RnE '^[[:space:]]*uses:' .github/ .gitea/ .forgejo/ 2>/dev/null | grep -vE '@[0-9a-f]{40}([[:space:]]|$)' || true)
          if [[ -n "$bad" ]]; then
            echo "::error::Unpinned actions found (must be 40-char SHAs):"
            echo "$bad"
            exit 1
          fi

  vuln-scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@9c091bb21b7c1c1d1991bb908d89e4e9dddfe3e0  # v7.0.0
      # step-level gate: hashFiles() is not valid in a job-level if
      - name: govulncheck (inside casjaysdev/go:latest)
        if: hashFiles('go.sum') != ''
        run: |
          IMAGE="casjaysdev/go:latest"
          docker run --rm -i \
            --name "$(basename "$GITHUB_REPOSITORY")-$(tr -dc 'a-z0-9' </dev/urandom | head -c8)" \
            -v "$PWD":/work -w /work -e CGO_ENABLED=0 -e GOFLAGS=-buildvcs=false \
            "$IMAGE" govulncheck ./...

  image-scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@9c091bb21b7c1c1d1991bb908d89e4e9dddfe3e0  # v7.0.0
      # step-level gates: hashFiles() is not valid in a job-level if
      - uses: docker/setup-buildx-action@d7f5e7f509e45cec5c76c4d5afdd7de93d0b3df5  # v4.1.0
        if: hashFiles('docker/Dockerfile') != ''
      - name: Build local image for scanning
        if: hashFiles('docker/Dockerfile') != ''
        run: |
          docker build -f docker/Dockerfile -t scan-target:ci .
      - name: Trivy image scan
        if: hashFiles('docker/Dockerfile') != ''
        uses: aquasecurity/trivy-action@ed142fd0673e97e23eac54620cfb913e5ce36c25  # v0.36.0
        with:
          image-ref: scan-target:ci
          severity: CRITICAL,HIGH
          exit-code: '1'
```

**Per-provider notes:**

- GitLab: `secret-scan` runs as a docker job using `image: trufflesecurity/trufflehog:latest` with `GIT_DEPTH: 0`; `image-scan` uses `image: aquasec/trivy:0.70.0`.
- Jenkins: `Security` stage uses `parallel {}`; truffleHog and Trivy each run via `docker.image(...).inside { ... }`.
- All providers: same gates, same severities, same exit conditions — no weaker subset on any provider.

## Post-Push CI Verification

`act --list` and a local `make test` pass only prove the workflow's syntax/job graph is valid and the code works in the local environment — they are not the real CI build. Every push (normal feature-branch push or an emergency direct push to the default branch) triggers a real CI run on the provider's infrastructure with real secrets, real matrix jobs, and the real `casjaysdev/go:latest` toolchain image; any of those can fail even when every local check passed. Treating "local checks passed" as equivalent to "the build is green" is itself a bug.

After every push, check the triggered run's status:
- **GitHub**: `gh run list --branch {branch} --commit {sha} --limit 1` then `gh run watch {run-id}` (or `gh run view {run-id} --json status,conclusion`)
- **GitLab**: `curl -qsSf -H "PRIVATE-TOKEN: $GITLAB_TOKEN" ".../repository/commits/{sha}/statuses"`
- **Gitea / Forgejo**: `curl -qsSf -H "Authorization: token $TOKEN" ".../commits/{sha}/status"`
- **Jenkins**: poll the job's `lastBuild/api/json` for `result`

Build failed → this is a bug, not a note for later; diagnose the root cause and fix it with a follow-up commit — never leave the default branch red. Build pending/running → the task is not done yet; wait and re-check. No CI config in the project → this step is a no-op.

## Suggested CI Steps

CI runs every Go step inside `casjaysdev/go:latest`. CI MUST NOT install a Go toolchain on the runner and call `go` directly — and MUST NOT run quality-gate commands inside the runtime image (`docker/Dockerfile`), which contains only the final binary. Use `casjaysdev/go:latest` for all build, test, lint, and vet steps.

**Single time source (all providers — GitHub Actions, Gitea/Forgejo, GitLab CI, Jenkins):** every CI job that builds binaries or images captures `BUILD_EPOCH="$(date -u +%s)"` exactly once (GitHub / Gitea / Forgejo: `echo "BUILD_EPOCH=${BUILD_EPOCH}" >> "$GITHUB_ENV"`; GitLab CI and Jenkins: a shell/environment variable set at the top of the stage). Every ldflags string passes `-X 'main.BuildEpoch=${BUILD_EPOCH}'`, every image build passes `--build-arg BUILD_EPOCH="${BUILD_EPOCH}"`, and `BUILD_DATE` is only ever derived from it — `BUILD_DATE="$(date -u -d "@${BUILD_EPOCH}" +%Y-%m-%dT%H:%M:%SZ)"` — where an OCI `org.opencontainers.image.created` annotation needs it. `BUILD_DATE` is never independently captured and never embedded via ldflags.

### Required Concurrency and Retention Headers

Every push/PR workflow (`ci.yml`) MUST declare:

```yaml
concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true
```

Release workflows (`release.yml`) MUST use `cancel-in-progress: true` — a newer tag push supersedes the in-flight release build, so the superseded run should be cancelled rather than left to finish.

Every `actions/upload-artifact` step MUST set a finite `retention-days`:

```yaml
- uses: actions/upload-artifact@043fb46d1a93c77aae656e7c1c64a875d1fc6a0a  # v7.0.1
  with:
    name: {project_name}-${{ matrix.target }}
    path: binaries/
    # release-job artifacts may use up to 30; build-job CI artifacts use 7
    retention-days: 7
```

```bash
# Prepare output directory for release artifacts (binaries, checksums, SBOM)
mkdir -p binaries

# BUILD_EPOCH is the single captured time source - captured once per job.
# BUILD_DATE is derived from it and used only for Docker OCI created annotations - never an ldflag
BUILD_EPOCH="$(date -u +%s)"
BUILD_DATE="$(date -u -d "@${BUILD_EPOCH}" +%Y-%m-%dT%H:%M:%SZ)"

# Run quality gates inside casjaysdev/go:latest
docker run --rm --name "${PROJECT_NAME}-$(tr -dc 'a-z0-9' </dev/urandom | head -c8)" \
  -v "$PWD":/app -w /app -e CGO_ENABLED=0 -e GOFLAGS=-buildvcs=false \
  casjaysdev/go:latest gofmt -l .
docker run --rm --name "${PROJECT_NAME}-$(tr -dc 'a-z0-9' </dev/urandom | head -c8)" \
  -v "$PWD":/app -w /app -e CGO_ENABLED=0 -e GOFLAGS=-buildvcs=false \
  casjaysdev/go:latest golangci-lint run ./...
docker run --rm --name "${PROJECT_NAME}-$(tr -dc 'a-z0-9' </dev/urandom | head -c8)" \
  -v "$PWD":/app -w /app -e CGO_ENABLED=0 -e GOFLAGS=-buildvcs=false \
  casjaysdev/go:latest go test ./...
docker run --rm --name "${PROJECT_NAME}-$(tr -dc 'a-z0-9' </dev/urandom | head -c8)" \
  -v "$PWD":/app -w /app -e CGO_ENABLED=0 -e GOFLAGS=-buildvcs=false \
  casjaysdev/go:latest go vet ./...
docker run --rm --name "${PROJECT_NAME}-$(tr -dc 'a-z0-9' </dev/urandom | head -c8)" \
  -v "$PWD":/app -w /app -e CGO_ENABLED=0 -e GOFLAGS=-buildvcs=false \
  casjaysdev/go:latest govulncheck ./...

# License enumeration and attribution-drift check:
# Regenerate the GENERATED region and compare to committed.
docker run --rm --name "${PROJECT_NAME}-$(tr -dc 'a-z0-9' </dev/urandom | head -c8)" \
  -v "$PWD":/app -w /app -e CGO_ENABLED=0 -e GOFLAGS=-buildvcs=false \
  casjaysdev/go:latest go-licenses report ./... > LICENSE.generated.md
sed -n '/<!-- GENERATED:/,$p' LICENSE.md > LICENSE.committed-generated.md
diff LICENSE.committed-generated.md LICENSE.generated.md

# Build statically linked release binaries (one per supported target).
for TARGET in "linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64" "windows/amd64" "windows/arm64" "freebsd/amd64" "freebsd/arm64"; do
  GOOS="${TARGET%/*}"
  GOARCH="${TARGET#*/}"
  EXT=""
  [ "$GOOS" = "windows" ] && EXT=".exe"
  ARTIFACT="{project_name}-${GOOS}-${GOARCH}${EXT}"

  docker run --rm \
    --name "${PROJECT_NAME}-$(tr -dc 'a-z0-9' </dev/urandom | head -c8)" \
    -v "$PWD":/app -w /app \
    -e CGO_ENABLED=0 -e GOFLAGS=-buildvcs=false -e GOOS="$GOOS" -e GOARCH="$GOARCH" \
    casjaysdev/go:latest \
    go build -buildvcs=false -trimpath \
      -ldflags="-s -w -X main.Version=$(cat release.txt) -X main.CommitID=$(git rev-parse --short=7 HEAD) -X 'main.BuildEpoch=${BUILD_EPOCH}'" \
      -o "binaries/$ARTIFACT" ./src

  # Verify static linkage on Linux
  if [ "$GOOS" = "linux" ]; then
    docker run --rm --name "${PROJECT_NAME}-$(tr -dc 'a-z0-9' </dev/urandom | head -c8)" \
      -v "$PWD":/app -w /app casjaysdev/go:latest \
      sh -c "ldd binaries/$ARTIFACT 2>&1 | grep -qE 'not a dynamic executable|statically linked'"
  fi
done

# Generate the SBOM (CycloneDX JSON) — published alongside the release artifacts.
docker run --rm --name "${PROJECT_NAME}-$(tr -dc 'a-z0-9' </dev/urandom | head -c8)" \
  -v "$PWD":/app -w /app -e CGO_ENABLED=0 -e GOFLAGS=-buildvcs=false \
  casjaysdev/go:latest cyclonedx-gomod app -json -output "binaries/{project_name}-bom.json"

# Generate aggregate checksums over every release artifact - the file list is
# pre-captured so sha256.txt / sha512.txt never hash themselves; both files are
# uploaded as release assets alongside the binaries
(cd binaries && FILES="$(ls)" && sha256sum $FILES > sha256.txt && sha512sum $FILES > sha512.txt)
```

For GUI smoke tests in CI, use a virtual X server (e.g., `Xvfb`) and a headless Wayland compositor (e.g., `cage`, `weston --backend=headless`) **inside** the container or as a sidecar service — both backends MUST be exercised, not just one.

## Workflow Error Messaging

Use `::error::` workflow commands for validation failures — they appear as red annotations on the Actions summary page, not just buried in logs:

```bash
echo "::error::Tag 'foo' does not exist in this repository"
echo "::error file=.github/workflows/release.yml,line=12::message tied to a source location"
```

Always write messages so a developer reading only the step name + message understands what failed and what to do next.

## Release Pre-flight Validation

The GitHub Releases API returns HTTP 422 `"tag_name is not a valid tag"` when the tag does not exist at API call time or is malformed. The correct fix is for the **release job to own the tag** — delete it if it exists, then recreate it at the current HEAD. This ensures the tag always exists and points to the right commit, and makes the workflow idempotent.

The `release` job already has `contents: write` to push assets — this covers tag push as well.

```yaml
- uses: actions/checkout@9c091bb21b7c1c1d1991bb908d89e4e9dddfe3e0  # v7.0.0
  with:
    # required: full history needed to inspect and push tags
    fetch-depth: 0

- name: Ensure release tag
  run: |
    ref="${{ github.ref }}"
    if [[ "$ref" != refs/tags/* ]]; then
      echo "::error::release.yml triggered on non-tag ref '$ref'. Releases require a tag push (refs/tags/v...)."
      exit 1
    fi
    tag="${ref#refs/tags/}"
    if printf '%s' "$tag" | grep -qP '[[:space:][:cntrl:]]'; then
      echo "::error::Tag '$tag' contains whitespace or control characters and is not a valid GitHub tag name."
      exit 1
    fi
    # Delete existing tag (local + remote) then recreate at current HEAD
    git tag -d "$tag" 2>/dev/null || true
    git push origin ":refs/tags/$tag" 2>/dev/null || true
    git tag "$tag"
    git push origin "refs/tags/$tag"
    echo "Tag '$tag' ensured at $(git rev-parse HEAD)"
```

## Release Integrity

Tagged/manual releases should publish:
- binaries/packages
- aggregate `sha256.txt` and `sha512.txt` checksum files covering every artifact
- release notes
- SBOM (`CycloneDX` from `cyclonedx-gomod`; `SPDX JSON` is acceptable if a project chooses that format instead) — always
- provenance / attestation — when the release platform supports it

If signing or attestation is required but keys/permissions are unavailable, stop and ask instead of faking compliance.

---

# PART 11: DOCUMENTATION & LICENSE

## README Minimum Sections

- Badges (linked — see Badges rule below)
- About
- Features
- Installation
- Usage
- GUI/TUI/CLI mode behavior (GUI section MUST cover both X11 and Wayland)
- Configuration
- Development (MUST document the Docker-based workflow — no host `go` invocations)
- Testing (MUST show Docker-wrapped commands; GUI tests show X11 and Wayland forwarding)
- License

### Badges

**Every badge MUST be a linked badge** — `[![alt](image_url)](link_url)` is the only valid form. A bare `![alt](image_url)` with no wrapping link is never acceptable. Each badge links to the resource it represents:

| Badge | Links to |
|-------|----------|
| CI / build status | CI runs page — detect platform: `.github/workflows/*.yml` → GitHub Actions; `.gitea/workflows/*.yml` → Gitea/Forgejo; `.gitlab-ci.yml` → GitLab CI; `Jenkinsfile` → Jenkins |
| Release / version | Releases page |
| License | `LICENSE.md` |
| Docs | Documentation site URL |

Never use a GitHub Actions badge for a GitLab or Gitea project — the CI badge must match the actual hosting platform.

**Badge Templates by Platform:**

```markdown
# GitHub Actions
[![CI](https://github.com/{project_org}/{project_name}/actions/workflows/ci.yml/badge.svg)](https://github.com/{project_org}/{project_name}/actions/workflows/ci.yml)

# Gitea/Forgejo Actions
[![CI](https://git.example.com/{project_org}/{project_name}/actions/workflows/ci.yml/badge.svg)](https://git.example.com/{project_org}/{project_name}/actions)

# GitLab CI
[![Build](https://gitlab.com/{project_org}/{project_name}/badges/main/pipeline.svg)](https://gitlab.com/{project_org}/{project_name}/-/pipelines)

# Jenkins
[![Build](https://jenkins.example.com/buildStatus/icon?job={project_org}/{project_name})](https://jenkins.example.com/job/{project_org}/job/{project_name}/)
```

**Release/License/Docs badges also adapt to platform:**

```markdown
# GitHub
[![Release](https://img.shields.io/github/v/release/{project_org}/{project_name})](https://github.com/{project_org}/{project_name}/releases)
[![License](https://img.shields.io/github/license/{project_org}/{project_name})](LICENSE.md)
[![Docs](https://readthedocs.org/projects/{rtd_project}/badge/?version=latest)](https://{rtd_url})

# GitLab
[![Release](https://gitlab.com/{project_org}/{project_name}/-/badges/release.svg)](https://gitlab.com/{project_org}/{project_name}/-/releases)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE.md)
[![Docs](https://readthedocs.org/projects/{rtd_project}/badge/?version=latest)](https://{rtd_url})

# Gitea/Forgejo (use shields.io with custom endpoint or static badge)
[![Release](https://img.shields.io/badge/dynamic/json?url=https://git.example.com/api/v1/repos/{project_org}/{project_name}/releases/latest&query=$.tag_name&label=release)](https://git.example.com/{project_org}/{project_name}/releases)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE.md)
[![Docs](https://readthedocs.org/projects/{rtd_project}/badge/?version=latest)](https://{rtd_url})

# {rtd_project} and {rtd_url} - Use one of:
#   {project_org}-{project_name} / {project_org}-{project_name}.readthedocs.io
#   {project_name} / {project_name}.readthedocs.io
#   Custom project name / {custom_rtd_address}
```

**License badge — use the GitHub API endpoint, not a static badge:**

```markdown
# ✅ CORRECT - GitHub can detect license
[![License](https://img.shields.io/github/license/{project_org}/{project_name})](LICENSE.md)

# ❌ WRONG - Static badge, GitHub cannot detect
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE.md)
```

**Docs badge — only add when documentation is actually deployed:**

| Docs Platform | Badge |
|---------------|-------|
| ReadTheDocs | `[![Docs](https://readthedocs.org/projects/{project_org}-{project_name}/badge/?version=latest)](https://{project_org}-{project_name}.readthedocs.io)` |
| GitHub Pages | `[![Docs](https://img.shields.io/badge/docs-GitHub%20Pages-blue)](https://{project_org}.github.io/{project_name})` |
| Self-hosted | `[![Docs](https://img.shields.io/badge/docs-online-blue)]({docs_url})` |
| None | **Do not add docs badge** |

## LICENSE.md Requirements

`LICENSE.md` is the canonical attribution document. It MUST contain:

1. **Project license text** — verbatim copy of the MIT license (or whatever IDEA.md declares), with the project's copyright line
2. **Third-party dependency attributions** — for every direct and transitive module in `go.sum` whose license requires attribution (MIT, BSD-2/3-Clause, Apache-2.0, ISC, MPL-2.0, Zlib, BSL-1.0, OpenSSL, Unicode-DFS-2016, etc. — i.e., everything except true public domain). Public-domain / CC0-1.0 packages SHOULD still be listed for transparency
3. **Verbatim NOTICE / COPYING / LICENSE preservation** — preserve upstream notice files verbatim whenever the upstream license requires it. Common triggers: (a) packages that vendor C code — preserve the upstream C project's NOTICE / COPYING / LICENSE alongside the package's own license; (b) Apache-2.0 dependencies that ship a NOTICE file — preserve verbatim
4. **Embedded asset notices** — fonts, icons, theme data, sample data, locale catalogs etc. that are embedded in the binary; each shipped asset's license/attribution as required by its source
5. **Generation banner** — a header comment recording which tool generated the third-party section and from which `go.sum` revision, so reviewers can tell at a glance whether the file is current

**Rule:** `LICENSE.md` is generated, not hand-written, for the third-party section (see "License Compliance" below). The project license text and upstream NOTICE preservation pieces are hand-written and MUST be merged with the generated section, not overwritten by it.

## License Compliance

License compliance is enforced by tooling and gated in CI; it is not left to manual review.

### Required Tooling

| Tool | Purpose |
|------|---------|
| `govulncheck` | Security advisory scanning for known Go vulnerabilities |
| `go-licenses` | License enumeration and generation of the third-party attribution section of `LICENSE.md` from `go.sum` |
| `cyclonedx-gomod` | Generate the SBOM (CycloneDX format) referenced in PART 5 → "Release Artifacts" and PART 10 → "Release Integrity" |
| `golangci-lint` | Lint enforcement; configured to ban import patterns violating the pure-Go rule |

All tools run inside the project Docker image. The image MUST have them pre-installed.

### License Allowlist (default)

| Allowed | Notes |
|---------|-------|
| `MIT` | Project default; baseline |
| `BSD-2-Clause`, `BSD-3-Clause` | Permissive; require attribution |
| `Apache-2.0`, `Apache-2.0 WITH LLVM-exception` | Require NOTICE preservation when one is shipped |
| `ISC` | Permissive; require attribution |
| `MPL-2.0` | File-level copyleft only — compatible with static linking into MIT binaries |
| `Zlib`, `BSL-1.0`, `Unicode-DFS-2016`, `Unicode-3.0` | Permissive |
| `CC0-1.0`, `Unlicense` | Public-domain equivalents |

Compound SPDX expressions like `MIT OR Apache-2.0` are accepted automatically when any alternative in the expression is on the allowlist. No special handling required.

### License Denylist (default — deny without IDEA.md exception)

| Denied | Reason |
|--------|--------|
| `GPL-2.0-*`, `GPL-3.0-*` | Strong copyleft; static linking forces the entire distributed binary to GPL, silently relicensing it away from MIT |
| `AGPL-1.0-*`, `AGPL-3.0-*` | Strong copyleft + network-use clause; even non-distributed deployments incur source-release obligations |
| `LGPL-2.0-*`, `LGPL-2.1-*`, `LGPL-3.0-*` | LGPL §6 under static linking requires shipping object files / a written relink offer alongside the binary, breaking the single-binary deliverable |
| `SSPL-1.0` | Service-source copyleft; treated like AGPL for our purposes |
| `Commons-Clause`, `BUSL-*`, `Elastic-2.0`, `Confluent-Community-*` | Source-available with commercial-use restrictions; not OSI-approved; incompatible with the "Free & open source" rule (PART 0 → "Licensing & Features") |

### IDEA.md Exception Path

A project MAY use a denylisted module only if **all** of:

1. `IDEA.md` adds (or extends) the `**License exceptions:**` block under `## Business logic`, naming the module, the upstream license, and the rationale (IDEA.md keeps exactly three top-level sections — this is a block, never a new `##` section)
2. The exception explicitly accepts the consequence — e.g., "this binary is distributed under GPL-3.0; the project's MIT claim applies only to the source we author, not the published binary" — and `IDEA.md ## Project variables` adds (or updates) the variable `distribution_license` accordingly. `distribution_license` is an exception-only variable: it is not part of the required-keys list and is present only in projects that have taken a license exception
3. README, `LICENSE.md`, and any user-visible "About" surface are updated to reflect the actual distribution license, not just the source license
4. Any tooling configuration (e.g., `golangci-lint` import ban config) is updated with a scoped allow entry for the specific module + version, not a blanket category unblock

There is no other path to use a denylisted module. Adding it without going through these steps is a spec violation.

### Required `go.mod` Metadata

Every `go.mod` MUST include a `module` path and a `go` version directive:

```
module github.com/{project_org}/{project_name}

go 1.23
```

And `README.md` / `LICENSE.md` MUST include:
- project license (SPDX expression or full text)
- module path
- maintainer info

### Generation Pattern

`LICENSE.md` has two regions:

```markdown
<!-- HAND-WRITTEN: project license, upstream NOTICE files, embedded asset notices -->

# Project License

(MIT text here)

# Embedded Assets

(font / icon / locale attributions here)

# Upstream NOTICE Files

(verbatim NOTICE / COPYING from any vendored C dependency)

<!-- GENERATED: third-party module attributions -->
<!-- Generated by go-licenses from go.sum @ <commit> on <date> — do not edit by hand -->

# Third-Party Module Attributions

(go-licenses output here)
```

The generated region is regenerated on every dependency change. The hand-written region is preserved across regenerations.

### User-Visible Attribution Surface

Embedded license data MUST be reachable by the user at runtime, not just shipped in the source repo:

- **CLI**: `--licenses` (alias: `--credits`) prints the full embedded `LICENSE.md` to stdout
- **TUI**: a "Licenses" / "Credits" entry in the help / about screen, scrollable
- **GUI**: an "About → Open Source Licenses" entry that opens a scrollable view of the full text

All three surfaces read the same `LICENSE.md` blob embedded at compile time via `//go:embed` (PART 0 → "Self-Contained Assets"). The version, commit ID, and build date (PART 6) are shown alongside.

### CI Gate (mandatory)

The canonical CI invocation lives in PART 10 → "Suggested CI Steps" (the `govulncheck` command and the `go-licenses report` + `sed`/`diff` drift check). This section defines the **policy**; PART 10 owns the script template. Do not duplicate the commands here when editing — keep PART 10 as the single source.

Drift between `go.sum` and the generated section of `LICENSE.md` is a CI failure, not a warning. The module graph is the source of truth for what licenses we ship, and `LICENSE.md` MUST match it.

## Documentation Style

- Keep docs precise and operational
- Prefer command snippets that actually match the Go project layout
- If GUI/TUI/CLI differ in behavior, document the differences explicitly

---

# PART 12: CHECKLISTS

## Bootstrap Checklist

- [ ] `IDEA.md` exists
- [ ] `IDEA.md` has `## Project description`
- [ ] `IDEA.md` has `## Project variables`
- [ ] `IDEA.md` has `## Business logic`
- [ ] `project_name`, `project_org`, `internal_name`, and `internal_org` exist
- [ ] If a pre-existing `CLAUDE.md` or `.claude/CLAUDE.md` existed, project-specific content was migrated into IDEA.md
- [ ] `CLAUDE.md` / `.claude/CLAUDE.md` are short loaders, not duplicate specs
- [ ] `.claude/memory/` directory exists (with an empty `MEMORY.md` index if no entries yet) and is committed, not gitignored
- [ ] `release.txt` exists if the project is using explicit release versioning
- [ ] `site.txt` exists only if there is a real official site URL
- [ ] `docker/Dockerfile`, `docker/docker-compose.yml`, `docker/docker-compose.dev.yml`, `docker/docker-compose.test.yml`, and `docker/rootfs/usr/local/bin/entrypoint.sh` exist; `Dockerfile` is the runtime image
- [ ] `.dockerignore` exists and excludes `binaries/`, `releases/`, and other build noise — but never `docker/`, `src/`, `go.mod`, `go.sum`, or `release.txt`
- [ ] `docker/Dockerfile.build` does not exist — Go projects always use `casjaysdev/go:latest` directly; no custom toolchain image is ever needed
- [ ] `CGO_ENABLED=0` is set as default in the Docker image environment
- [ ] X11 forwarding sample command is documented and works against a real Xorg/XWayland session
- [ ] Wayland forwarding sample command is documented and works against a real Wayland compositor
- [ ] `.go-version` pins the Go version; `go.mod` declares the same version
- [ ] `go.mod` has the correct `module` path and `go` directive
- [ ] Repo has `assets/` (build-time source) and a `//go:embed` directive wiring it into the binary
- [ ] No source files in any language other than Go contribute to the binary (small Docker shell helpers excepted)
- [ ] `Makefile` exists with at minimum `build`, `test`, `dev`, `clean` targets
- [ ] `go-licenses` produces the third-party attribution section; `LICENSE.md` is up to date
- [ ] `govulncheck` passes clean
- [ ] `cyclonedx-gomod` is pre-installed in the Docker image and produces a valid SBOM

## Implementation Checklist

- [ ] `ci.yml` and `release.yml` use `container: image: casjaysdev/go:latest` — no `ensure-build-image` pre-flight, no `build-toolchain.yml`
- [ ] Core logic is shared across GUI/TUI/CLI
- [ ] Runtime mode auto-detect follows GUI → TUI → CLI priority
- [ ] GUI is not auto-selected in SSH/MOSH/remote-shell contexts
- [ ] If GUI is in scope, both X11 and Wayland backends are supported and runtime-selected by the GUI toolkit
- [ ] TUI is avoided for `TERM=dumb`, non-TTY, and plain-output contexts
- [ ] CLI works for non-interactive/scripted use
- [ ] Privilege escalation is optional and explicit-request only
- [ ] Config/data/cache/log paths use per-user defaults unless system scope was explicitly requested
- [ ] All assets (fonts, icons, themes, default config, schemas, locales) are embedded at compile time via `//go:embed`
- [ ] Dependencies are pure Go where a viable pure-Go package exists (PART 5 → "Pure-Go Library Stack"); each C-dependent package is justified in IDEA.md
- [ ] `go mod graph` was reviewed — no surprise transitive C or CGO dependencies
- [ ] No GPL / AGPL / LGPL / SSPL / BUSL / source-available dep was added without an IDEA.md `**License exceptions:**` entry and an updated `distribution_license` (PART 11 → "License Compliance")
- [ ] User-visible licenses surface exists: CLI `--licenses`, TUI "Licenses" entry, GUI "About → Open Source Licenses" — all reading the same embedded `LICENSE.md`
- [ ] App runs end-to-end from the binary alone on an air-gapped machine — no first-run downloads
- [ ] No plugin or runtime extension loading from disk unless IDEA.md defines a hardened plugin contract
- [ ] `release.txt` / `site.txt` precedence is preserved
- [ ] Docs and examples use Go/make terminology — wrapped in Docker invocations
- [ ] No build/test/run instructions tell the user to invoke `go` on the host
- [ ] If IDEA.md declares an RFC protocol daemon: the PART 14 daemon checklist passes
- [ ] If IDEA.md declares a local client/server (IPC) split: sockets live in `$XDG_RUNTIME_DIR/{project_name}/` with `0700`/`0600` modes, the main-socket registry model is implemented, and no TCP listener exists (PART 2 → "Local Client/Server (IPC) Mode")

## Quality Checklist

All gates run inside the project Docker image — never on the host.

**Commit-time lint gate:** running `golangci-lint run ./...` directly satisfies the
quality bar shown here, but does NOT satisfy `gitcommit`'s pre-commit lint gate —
`lint-agent-mark.sh` only records the gate as passed when the `go-lint` agent itself
reports a clean result. Invoke the `go-lint` agent before running `gitcommit`; a raw
`golangci-lint` invocation leaves the commit permanently blocked.

- [ ] `gofmt -l .` reports no files needing formatting (Docker-wrapped)
- [ ] `golangci-lint run ./...` passes with no warnings (Docker-wrapped)
- [ ] `go test ./...` passes (Docker-wrapped)
- [ ] `go vet ./...` passes (Docker-wrapped)
- [ ] `govulncheck ./...` passes (Docker-wrapped)
- [ ] `go-licenses report ./...` output matches the generated region of `LICENSE.md` (Docker-wrapped diff check)
- [ ] If GUI is in scope: smoke test exercised under both X11 and Wayland forwarding
- [ ] New behavior includes tests where practical

## Security Checklist

- [ ] No hidden telemetry
- [ ] No secrets in logs
- [ ] No automatic privilege escalation
- [ ] No unsafe downloaded-code execution
- [ ] `govulncheck` and license checks remain enabled in CI

## Release Checklist

- [ ] Version comes from `release.txt` when present
- [ ] Official site comes from `site.txt` when present
- [ ] Release notes match actual changes
- [ ] Aggregate `sha256.txt` and `sha512.txt` checksum files are published, covering every artifact
- [ ] Each release artifact is a single statically linked binary — no companion `.so` / `.dylib` / `.dll` / asset bundle
- [ ] Static-linkage check (`ldd` / `otool -L` / `dumpbin /dependents`) was run and recorded for every published target
- [ ] `LICENSE.md` regenerated and committed if `go.sum` changed; CI license-drift check is green
- [ ] Distributed binary's actual license matches what README and "About" claim (no silent GPL/LGPL relicensing via a transitive module)
- [ ] Artifact filenames follow `{project_name}-{platform}-{arch}{.ext}` with no OS suffix tokens
- [ ] Artifacts cover the supported target matrix declared in IDEA.md (defaults: `{project_name}-linux-{amd64,arm64}`, `{project_name}-darwin-{amd64,arm64}`, `{project_name}-windows-{amd64,arm64}.exe`, `{project_name}-freebsd-{amd64,arm64}`)
- [ ] SBOM generated via `cyclonedx-gomod` and published with the release; provenance/attestation included when the platform supports it
- [ ] Packaging metadata matches supported GUI/TUI/CLI surfaces

## Red Flags (Stop and Ask Human)

- required project variables cannot be proven
- migration would discard meaningful project-specific content
- privileged/system-wide action is needed but was not explicitly requested
- licensing or security policy conflict cannot be resolved from the spec
- requested behavior contradicts documented product scope in IDEA.md
- a proposed dependency would force CGO, dynamic linking into the release binary, or a runtime asset alongside the binary
- a feature request implies shipping multiple files, an installer-only flow, or a first-run download
- a proposed GUI toolkit only supports X11 or only supports Wayland — not both
- a proposed dependency is GPL / AGPL / LGPL / SSPL / BUSL / source-available — flag the license-relicensing implication before adding (PART 11 → "License Compliance")

## Success Criteria

A compliant Go project following this specification:
- is driven by `IDEA.md` project variables while `AI.md` stays read-only
- preserves the governance/documentation discipline of this specification
- models a single-binary, fully self-contained Go application built around GUI / TUI / CLI surfaces (plus the conditional RFC protocol daemon mode of PART 14 and the conditional local client/server IPC mode of PART 2)
- ships exclusively Go source code (small Docker shell helpers excepted)
- produces one statically linked binary per target with all assets embedded
- runs end-to-end from the binary alone on an air-gapped machine
- supports both X11 and Wayland on Linux/BSD when GUI is in scope (via the chosen pure-Go GUI toolkit)
- builds, tests, and runs only inside the project's Docker image — never on the host
- auto-selects GUI, then TUI, then CLI using environment-aware detection
- uses optional explicit privilege escalation only when requested
- uses Go tooling (via `make` targets and Docker) consistently across development, testing, and CI/CD
- maintains `CGO_ENABLED=0` across every build target without exception

---

# PART 13: IDEA.md REFERENCE

**NEVER modify this PART. Read-only reference for IDEA.md structure.**

The canonical layout, variable rules, immutability constraints, and migration procedure are defined at the top of this file under "IDEA.md Required Layout" (the section preceding PART 0). PART 13 here is the worked-example reference — every example below uses the same three-section structure.

---

## IDEA.md Structure

Every IDEA.md has exactly three top-level sections, in this order:

1. `## Project description` — free-form prose: what the project is, who uses it, what problem it solves
2. `## Project variables` — `key: value` lines that provide the canonical values AI.md resolves for `project_name`, `project_org`, `internal_name`, `internal_org`, etc.
3. `## Business logic` — features, data models, user flows, trust boundaries, abuse cases, platform constraints, security assumptions (WHAT, not HOW)

See "IDEA.md Required Layout" at the top of this file for the authoritative rules: variable-key naming, the immutable `internal_name` / `internal_org` rule, the missing-value setup flow, and the migration procedure for legacy `CLAUDE.md` files.

---

## IDEA.md Template

```markdown
## Project description

{Brief description of what this project does, its primary users, and what problem it
solves. Free-form prose, 1–3 paragraphs.}

## Project variables

project_name:     {project_name}
project_org:      {project_org}
# FROZEN — equals project_name on first install, never changes
internal_name:    {project_name}
# FROZEN — equals project_org on first install, never changes
internal_org:     {project_org}
app_name:         {App Display Name}
module_path:      github.com/{project_org}/{project_name}
official_site:    {full URL with scheme, e.g., https://example.com — or empty}
maintainer_name:  {Maintainer Name — defaults to {project_org} if unset}
maintainer_email: {maintainer@example.com — or empty; used only if set}

## Business logic

**Target users:**
- {User type 1}
- {User type 2}

**Surfaces (declare which apply — see PART 2 → "GUI/TUI/CLI Capability Rule"):**
- GUI: {yes / no — and on which platforms}
- TUI: {yes / no}
- CLI: {yes / no}

**Features:**
- **{Feature category 1}**: {brief description}
- **{Feature category 2}**: {brief description}

**Data models:**
- {Entity}: {fields and meaning, where stored}

**User flows:**
- {Flow}: {step 1 → step 2 → outcome}

**Business rules:**
- {Validation, constraint, invariant}

**Platform constraints:**
- {OS support, hardware requirements, display-server requirements}

**Trust boundaries & abuse cases:**
- {Untrusted input source → where it is validated / what trusts what}
- {Abuse case: what a hostile input or user could attempt → expected mitigation posture}

**Security exceptions:** {intentionally allowed security-sensitive choice + documented reason — or `none`}

**Outbound network use (if any — see PART 9 → "Security-First Design"):**
- {Remote service consumed, purpose, auth model}

**Stored data location (per-user — see PART 4 → "Path Rule"):**
- {What is stored at the per-user paths}

**License exceptions (always present — write `none` when no denylisted or C-dependent module is used; see PART 11 → "License Compliance"):**
- {module path, upstream license, rationale, `distribution_license` consequence if applicable}
```

**Rules for the example contents above:**

- No implementation details — describe behavior, not algorithms or libraries. AI.md PARTs 0–11 define HOW; PART 12 verifies compliance.
- This specification targets GUI / TUI / CLI applications (PART 0 → "One Coherent Product").
- Cross-reference AI.md PARTs by number for any pattern that already exists there (Docker → PART 5, security → PART 9, license exceptions → PART 11, etc.).
- `internal_name` and `internal_org` are immutable after first set (see "IDEA.md Required Layout" → Project variables rules).

---

## IDEA.md Examples

**Example 1: Notes (GUI + CLI, fully offline)**

```markdown
## Project description

A local-first notes application with a native GUI for daily use and a CLI for scripting.
Notes are stored on the user's machine — no cloud sync, no telemetry, fully offline.

## Project variables

project_name:     notes
project_org:      casjay
internal_name:    notes
internal_org:     casjay
app_name:         Notes
module_path:      github.com/casjay/notes
official_site:    https://notes.example.com
maintainer_name:  Jane Doe
maintainer_email: jane@example.com

## Business logic

**Target users:**
- Individual users wanting a fast, offline notes app
- Power users scripting note creation / search via shell

**Surfaces:**
- GUI: yes (Linux X11 + Wayland via Gio toolkit, macOS, Windows)
- TUI: no
- CLI: yes (`notes new`, `notes search`, `notes ls`)

**Features:**
- **Note management**: create, edit, archive, delete with markdown support
- **Search**: full-text search across all notes
- **Organization**: tag-based grouping; no folders
- **Export**: dump notes to a directory of `.md` files

**Data models:**
- Note: id, title, body (markdown), tags[], created_at, updated_at

**User flows:**
- Quick capture: Ctrl+N opens compose → write → Ctrl+S saves → return to list
- Search: Ctrl+K opens overlay → type → live-filter results

**Business rules:**
- Archive is reversible; delete is permanent
- All edits are local; no network
- Title can be empty; body cannot

**Platform constraints:**
- GUI requires X11 or Wayland on Linux/BSD (handled by Gio toolkit — supports both)
- No internet access required; works fully offline (PART 0 → "Self-Contained Assets")

**Trust boundaries & abuse cases:**
- All data is local and owned by the user; the only untrusted input is imported markdown, which is parsed for rendering and never executed
- Abuse case: a crafted markdown file targeting the parser → memory-safe parsing, no HTML/script execution, size limit on import

**Security exceptions:** none

**Outbound network use:** none

**Stored data location (per-user):**
- Notes database: `~/.local/share/casjay/notes/notes.db`
- Config: `~/.config/casjay/notes/config.toml`
- Cache: `~/.cache/casjay/notes/`

**License exceptions:** none — `modernc.org/sqlite` (BSD-3-Clause, pure Go) powers the local notes database but is allowlisted; ordinary `LICENSE.md` attribution applies, no exception needed.
```

**Example 2: Feeds (TUI + CLI, consumes a remote API)**

```markdown
## Project description

A terminal-first RSS / Atom feed reader. Subscribes to feeds, fetches new items in the
background, and presents them in a keyboard-driven TUI. A `feeds` CLI exposes the same
operations for automation.

## Project variables

project_name:     feeds
project_org:      casjay
internal_name:    feeds
internal_org:     casjay
app_name:         Feeds
module_path:      github.com/casjay/feeds
official_site:
maintainer_name:  Jane Doe
maintainer_email: jane@example.com

## Business logic

**Target users:**
- Terminal-first users who follow many feeds
- Anyone wanting offline-friendly RSS reading without a hosted service

**Surfaces:**
- GUI: no
- TUI: yes (primary)
- CLI: yes (`feeds add`, `feeds sync`, `feeds export`)

**Features:**
- **Subscription**: add / remove / rename feeds; OPML import / export
- **Sync**: fetch all feeds in parallel; respect HTTP cache headers
- **Reading**: keyboard navigation; mark read / unread; star items
- **Filtering**: by feed, by tag, by unread state

**Data models:**
- Feed: id, url, title, last_etag, last_modified, last_synced
- Item: id, feed_id, guid, title, body, published_at, read, starred

**User flows:**
- Launch → background sync starts → pick a feed → arrow keys to navigate → Enter to read
- `feeds sync` from cron pulls new items without launching the TUI

**Business rules:**
- Items are deduplicated by GUID within a feed
- Sync respects ETag / Last-Modified to avoid re-fetching unchanged feeds
- Network failures degrade gracefully (read-only mode over cached items)

**Platform constraints:**
- TUI requires a TTY (PART 3 → "Smart Detect Rules"); for headless cron use the `feeds sync` CLI subcommand

**Trust boundaries & abuse cases:**
- Remote feed content (XML/HTML) is fully untrusted → hardened parsing, HTML sanitized to text, scripts and external resource loads never executed
- Feed URLs are user-supplied → fetched over HTTPS only; non-HTTP schemes refused; no credentials attached
- Abuse case: a hostile feed serving oversized or decompression-bomb responses → response size and time limits on sync

**Security exceptions:** none

**Outbound network use:**
- HTTPS GET to feed URLs the user adds. TLS via `crypto/tls` stdlib (PART 9 → "Security-First Design"). No telemetry, no tracking, no third-party analytics.

**Stored data location (per-user):**
- Feed / item database: `~/.local/share/casjay/feeds/feeds.db`
- Config: `~/.config/casjay/feeds/config.toml`
- Cache (article bodies, images): `~/.cache/casjay/feeds/`

**License exceptions:** none — `modernc.org/sqlite` (BSD-3-Clause, pure Go) powers the local feeds database but is allowlisted; ordinary `LICENSE.md` attribution applies, no exception needed.
```

**Example 3: dotctl (CLI only, manages dotfiles)**

```markdown
## Project description

A small CLI for managing user dotfiles: stow / unstow symlinks, version-pin individual
files, and verify that what is on disk matches what the user committed.

## Project variables

project_name:     dotctl
project_org:      casjay
internal_name:    dotctl
internal_org:     casjay
app_name:         dotctl
module_path:      github.com/casjay/dotctl
official_site:
maintainer_name:  Jane Doe
maintainer_email: jane@example.com

## Business logic

**Target users:**
- Developers managing dotfiles across multiple machines

**Surfaces:**
- GUI: no
- TUI: no
- CLI: yes (`dotctl link`, `dotctl unlink`, `dotctl status`, `dotctl diff`)

**Features:**
- **Link / unlink**: create or remove symlinks from a dotfiles repo into the user's home
- **Status**: report which managed files match, differ, or are missing
- **Diff**: show the diff between repo and on-disk versions
- **Atomic apply**: link operations are dry-run-able and reversible

**Data models:**
- ManagedFile: source_path (in repo), target_path (in home), state (linked / unlinked / divergent)

**User flows:**
- `dotctl status` shows the current state matrix
- `dotctl link --dry-run` previews what would change; `dotctl link` applies it

**Business rules:**
- Never overwrite a regular file without explicit `--force`
- Symlinks are created relative to the dotfiles repo so the home directory remains portable

**Platform constraints:**
- POSIX-style filesystem operations; on Windows, `os.Symlink` calls `CreateSymbolicLink`, producing real NTFS symlinks (requires Developer Mode or an elevated process) — junctions are a separate mechanism and are not used here

**Trust boundaries & abuse cases:**
- Repo contents are the user's own but treated as untrusted path input: every target path is validated to resolve inside `$HOME`
- Abuse case: a repo entry attempting path traversal (`../../etc/...`) or a symlink escape outside `$HOME` → refused; nothing outside `$HOME` is ever written

**Security exceptions:** none

**Outbound network use:** none

**Stored data location (per-user — PART 4 → "Path Rule"):**
- Config: `~/.config/casjay/dotctl/config.toml`
- State (last known link map): `~/.local/share/casjay/dotctl/state.json`

**License exceptions:** none
```

---

# PART 14: RFC PROTOCOL SERVER / DAEMON MODE (CONDITIONAL)

**This PART applies only when IDEA.md `## Business logic` declares the project an RFC protocol server (daemon). If it does not, this PART is inert and adds no requirements.**

**Local client/server (IPC) applications — tmux/nvim-style Unix-socket splits — are NOT protocol daemons and never use this PART**; they are governed by PART 2 → "Local Client/Server (IPC) Mode".

## What This Mode Is

A protocol daemon is an infrastructure server in the tradition of nginx/httpd/caddy (HTTP), proftpd/vsftpd (FTP), postfix/exim (SMTP), dovecot (IMAP/POP3), squid (HTTP proxy), or bind/unbound (DNS). It is still a single-binary application under this specification — it is NOT an API server and NOT a web frontend server:

| Type | Speaks | Governing spec |
|------|--------|----------------|
| API server | JSON/REST over HTTP for its own API | the API server specification |
| Full-stack web server | its own HTML frontend + optional REST | the web server specification |
| **Protocol daemon (this PART)** | **a standard wire protocol, exactly as its RFC defines** | this specification + this PART |

**Every listening socket speaks a standard, RFC-governed wire protocol. The daemon never exposes an invented protocol, and never mixes non-protocol traffic (dashboards, JSON APIs, health endpoints) into protocol sockets.**

## IDEA.md Declaration

IDEA.md `## Business logic` MUST declare:
- that the project is a protocol server (daemon)
- every protocol implemented, with its governing RFC numbers
- which optional protocol extensions are in scope

Example:

```markdown
### Protocols

| Protocol | RFCs | Extensions |
|----------|------|------------|
| SMTP (MTA) | RFC 5321, RFC 5322 | STARTTLS (RFC 3207), AUTH (RFC 4954), 8BITMIME (RFC 6152) |
| Submission | RFC 6409 | — |
```

## RFC Conformance Rules

| Rule | Detail |
|------|--------|
| **The RFC is the contract** | Wire behavior implements the RFCs IDEA.md declares. RFC 2119/8174 key words apply: RFC MUST requirements are hard gates; any SHOULD-level deviation requires an IDEA.md entry with reasoning |
| **Never invent wire protocols** | No proprietary commands, replies, or framing on protocol sockets |
| **Extensions via the protocol's own mechanism** | ESMTP EHLO keywords, IMAP CAPABILITY, FTP FEAT, HTTP headers/upgrade — advertised through the protocol's negotiation, declared in IDEA.md |
| **Protocol constants from the RFC** | Reply codes, status codes, and grammar come from the RFC — never invented values |
| **Interop is the ground truth** | Conformance is verified against real-world clients, not only unit tests |

## Common Protocol → RFC Map (Reference Only — IDEA.md's Declared List Governs)

| Protocol | Core RFCs | Example daemons |
|----------|-----------|-----------------|
| HTTP/1.1 | 9110, 9111, 9112 | nginx, httpd, caddy |
| HTTP/2 · HTTP/3 | 9113 · 9114 | nginx, caddy |
| HTTP proxy (CONNECT) | 9110 | squid |
| FTP | 959, 2228, 3659 | proftpd, vsftpd |
| SMTP / submission | 5321, 5322, 6409 | postfix, exim |
| IMAP4rev2 / POP3 | 9051 (3501), 1939 | dovecot |
| DNS | 1034, 1035 | bind, unbound |
| TLS / STARTTLS | 8446, 3207 | — |
| WebSocket | 6455 | — |

## Runtime Model

| Rule | Detail |
|------|--------|
| **Serve entrypoint** | `{project_name} serve` starts the daemon; IDEA.md may define aliases |
| **Foreground only** | The process runs in the foreground and never self-daemonizes (no fork/detach) — systemd or the container runtime supervises it |
| **UI mode** | `serve` forces CLI mode; the GUI/TUI auto-detect of PART 3 never applies to a service process. Optional GUI/TUI surfaces are management consoles over the same shared core |
| **Console output** | Startup banner, then console silence in normal operation — protocol traffic and periodic work log to files only; the console shows WARN/ERROR |

## Signals & Lifecycle

| Signal | Action |
|--------|--------|
| `SIGTERM` / `SIGINT` | Graceful shutdown: stop accepting, drain active sessions up to `daemon.shutdown_timeout` (default 30 seconds), then exit 0 |
| `SIGHUP` | Reload config and TLS certificates without dropping established sessions; if the new config is invalid, keep running on the old config and log ERROR |
| `SIGUSR1` | Reopen log files (for external rotation tooling; built-in rotation per PART 7 still applies) |

Signal handling uses `os/signal` (`signal.NotifyContext` for shutdown, a dedicated channel for `SIGHUP`/`SIGUSR1`).

**`--config-test` (`-t`) is REQUIRED:** parse and validate the full config, print errors to stderr, exit 0 (valid) or 1 (invalid) — mirrors `nginx -t`. It never touches sockets or running state.

## Sockets, Ports & Privileges

- IPv4 + IPv6 listeners (dual-stack by default; per-listener config)
- IANA standard port defaults for the protocol; every listener configurable
- **Privileged ports (<1024): bind first, then drop privileges** — target credentials are resolved before the drop, and the daemon never continues running as root after binding (`golang.org/x/sys/unix` setuid/setgid pattern)
- TLS: cert/key paths in config, reloaded on SIGHUP; implicit TLS vs STARTTLS per the protocol's RFC

## Health & Metrics Sidecar (Optional — OFF by Default)

```yaml
daemon:
  # Seconds to drain active sessions on SIGTERM
  shutdown_timeout: 30
  sidecar:
    # OFF by default — protocol sockets stay pure protocol
    enabled: false
    # Loopback only unless explicitly configured otherwise
    address: 127.0.0.1
    # REQUIRED when enabled — no implicit default port
    port: 0
```

| Rule | Detail |
|------|--------|
| **Disabled = no HTTP listener** | When `sidecar.enabled: false` (default), the daemon opens no HTTP socket at all |
| **Same contract as the server specs** | When enabled, the sidecar serves `/server/healthz` and `/server/metrics[/{service}]` with the exact same JSON shape, status semantics, and per-service bearer-token auth defined by the API/SERVER specifications |
| **Loopback bind** | `127.0.0.1` by default; binding a non-loopback address requires explicit config |
| **Never on protocol sockets** | Health/metrics never ride the protocol listeners — even when the daemon's protocol is HTTP itself, the sidecar is a separate listener |

## Protocol Logging

- Transaction logs use the protocol community's de-facto standard format where one exists (HTTP → Apache combined by default; FTP → xferlog; SMTP → per-message queue/delivery lines); otherwise structured text
- Same rotation/retention schema as PART 7 logging
- The wire is a machine surface: protocol bytes are exactly what the RFC prescribes — the human-readable formatting rules of PART 7 apply to CLI/status/TUI/GUI output only, never to wire output

## Testing

- Every implemented RFC MUST requirement has a conformance test; SHOULD-level behavior is tested when in scope
- Interop smoke tests run in `docker/docker-compose.test.yml` against real clients (`curl` for HTTP, `lftp` for FTP, `swaks` for SMTP, `openssl s_client` for TLS, `dig` for DNS)
- Malformed-input tests: negative/fuzz cases return the RFC-mandated error replies and never crash the daemon
- A failed RFC MUST requirement is a bug — never a documented quirk

## Daemon Checklist

- [ ] IDEA.md declares protocols, RFC numbers, and in-scope extensions
- [ ] All wire behavior traces to the declared RFCs; no invented commands, replies, or framing
- [ ] `serve` runs foreground in CLI mode; no self-daemonization
- [ ] SIGTERM drains gracefully within `shutdown_timeout`; SIGHUP reloads config/certs without dropping sessions; an invalid reload keeps the old config running
- [ ] `--config-test` validates the config and exits 0/1
- [ ] Privileged ports use bind-then-drop; the daemon never runs steady-state as root
- [ ] IPv4+IPv6 listeners with IANA defaults, all configurable
- [ ] Sidecar is OFF by default; when enabled it binds loopback and matches the server specs' healthz/metrics contract
- [ ] Transaction logs use the protocol's standard format; console is silent in normal operation
- [ ] Interop smoke tests pass against real clients in Docker
