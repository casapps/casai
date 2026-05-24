# 🧠 Project: `casai` — Unified AI Service Installer and Manager

A fully automated, GPU-aware, shell-based CLI project that installs, configures, manages, and integrates multiple LLM/AI tools using **binary-first** logic. Containers are used only when binaries are unavailable. The project adheres strictly to user-defined standards for path layout, system integration, model directories, shell behavior, permissions, configuration format, and tooling expectations. The CLI is a **single bash script**.

---

## 🧩 Core Intent

Build a complete LLM stack for development and runtime use, with support for:
- Ollama, LocalAI, Whisper.cpp, Piper, OpenDiffusion, OpenWebUI, and Opencode
- GPU auto detection and tiered model selection
- Model management, config management, full lifecycle support
- Shell completions, service management, secrets, permissions, and diagnostics

All configuration, logic, and CLI behavior follow user-defined architectural rules, with zero drift.

---

## 📦 OS and Distro Support

- Fully supports:
  - AlmaLinux 
  - Debian
  - Ubuntu
  - Alpine Linux
  - Arch Linux
  - And All derivatives
- Uses system tools, respects distro differences
- Root access required for service installation, config paths, system wiring

---

## 🧷 Directory Layout

- 🗂 Global config: `/etc/casai/casai.conf`
- 🧑 User config: `~/.config/casai/config.toml`
- 📓 MCP config: `/etc/casai/mcp.toml`
- 🔐 Secrets dir: `~/.local/share/casai/secrets`
- 🪵 Logs dir: `~/.local/log/casai`
- 🧠 Models dir: `/var/lib/casai/models`
- ⚙️ Per-service config: `/etc/{service}`
- ❌ No use of `~/bin`, `/opt`, or `/usr/local/etc`
- 📁 Binaries installed to: `/usr/local/bin`
- 🧪 Shell-completions enabled via: `/usr/local/bin` added to `PATH` (if not already)

---

## 🧠 GPU + Model Autodetect

- Smart GPU detection:
  - Handles passthrough, SR-IOV, multiple GPUs
  - Detects if usable, not just present
- Uses system specs to decide:
  - Which models to pull
  - Whether to install GPU-enabled binaries
  - Whether to fall back to CPU mode
- ✅ CLI: `casai --gpu` triggers full GPU system test
- 📊 CLI: `casai --why` explains model selection logic

---

## 🧠 Model Directory Standards

- Models stored in: `/var/lib/casai/models/{type}`
  - `gguf/`, `hf/`, `diffusion/`, `whisper/`, `tts/`, `ollama/`
- All paths configured per tool:
  - Ollama uses: `/var/lib/casai/models/ollama`
  - Piper uses: `/var/lib/casai/models/tts`
  - LocalAI supports multiple paths
- Permissions:
  - Directories: `chmod 755 using find -exec`
  - Files: `chmod 666 using find -exec`
  - Set before launching services using `smart-perms`
- Configurable via: `CASIA_MODELS_DIR`

---

## 🧠 Binary vs Container Rules

- All supported tools are installed as **binaries** if available:
  - Ollama
  - LocalAI
  - Whisper.cpp
  - Piper
  - Opencode (https://github.com/opencode-ai/opencode/releases)
- Containers are used only for:
  - OpenWebUI
  - OpenDiffusion

### Binary rules:
- Never use container if binary install is supported
- Use official release binary, not from source unless forced
- Installed into `/usr/local/bin`

### Container rules:
- Always use **official images**
- No Docker Desktop
- Use container `--name` (openwebui,opendiffusion)
- Volumes mounted:
  - Host config: `/etc/{service}` → `/etc/{service}`
  - Host models: `/var/lib/casai/models/{type}` → tool-specific mount point
  - Logs written to `~/.local/log/casai/{service}.log`
- No autostart — containers are controlled via `casai service`

---

## ⚙️ Service Control Rules

All tools (binary and container) are managed via:

```bash
casai service start|stop|restart|status <name>
```

- ❌ No systemd (except for us: casai runinit install,remove,disable,enable. exec casai service start|stop|restart|status all)
- ✅ Single source of control: `casai`

---

## 📥 Installation and Setup

- CLI: `casai install`
  - Installs all binaries and containers
  - Generates configs
  - Pulls models based on specs
  - Sets up permissions
- CLI: `casai update` updates everything
- CLI: `casai fix` resets permissions, regenerates config
- CLI: `casai uninstall` removes everything

---

## 🧩 AI Services

Each AI service is:

- Auto-configured
- Launched on a local port
- Logs stored in `~/.local/log/casai/{service}.log`
- Managed only via `casai`

| Service       | Type     | Managed Via     |
|---------------|----------|-----------------|
| Ollama        | Binary   | `casai`      |
| LocalAI       | Binary   | `casai`      |
| Piper         | Binary   | `casai`      |
| Whisper.cpp   | Binary   | `casai`      |
| Opencode      | Binary   | `casai`      |
| OpenWebUI     | Container| `docker --restart always --pull=always`      |
| OpenDiffusion | Container| `docker --restart always --pull=always`      |

---

## 🎛 MCP Tool Integration

- MCP tools are defined in `/etc/casai/mcp.toml`
- CLI: `casai mcp <tool>` to install/manage
- Follows authoritative MCP map
- All entries include:
  - Binary/container
  - Port
  - Healthcheck
  - Category
  - Autostart
  - Metadata (url, description)

---

## 🧠 CLI Commands

```bash
casai                   # Install all components if not init yet.
casai install           # Install all components
casai uninstall         # Remove all components
casai model <type>      # Manage models: add/remove/list
casai mcp <tool>        # Manage MCP tools
casai service <cmd> <name>    # Start/stop/restart/status
casai systemd <cmd>     # Install/remove/start/stop system service
casai shell <cmd>       # Shell integration (env/load/diff)
casai secrets <cmd>     # Encrypt/reset/audit secrets
casai update            # Update everything
casai fix               # Fix permissions, regenerate configs
casai summary           # Show system status
casai tts               # Speak with Piper
casai chat              # Chat using Ollama or LocalAI
casai version           # Show version
casai help              # Show help
```

---

## 🔐 Secrets

- Stored in `~/.local/share/casai/secrets`
- CLI: `casai secrets`
- Encrypted and auditable
- Can reset or wipe clean

---

## 🧠 Shell Integration

- Exports:
  - `CASIA_VERSION`
  - `CASIA_GLOBAL_CONFIG_DIR`
  - `CASIA_USER_CONFIG_DIR`
  - `CASIA_MODELS_DIR`
  - `CASIA_CACHE_DIR`
- Adds `~/.local/bin` to PATH
- CLI: `casai shell env|diff|load|reset`

---

## 🛡 Permissions

Before any operation that requires root or sudo:

```bash
if [ "$EUID" -ne 0 ]; then
  if sudo -n true 2>/dev/null; then
    sudo true || exit 1
  else
    echo "Root required"
    exit 1
  fi
fi
```

- After install: `chmod 755` for all dirs under `/var/lib/casai`, `chmod 666` for files

---

## 🧪 Health + Debug

- `/healthz` endpoints for web-based tools
- CLI: `casai --debug` shows logs
- CLI: `casai --why` explains model selection
- CLI: `casai summary` shows system overview

---

## 📄 Output Summary Example

```
📦 System initialized:
  • Hostname           → casai.local
  • OS                 → Ubuntu 22.04 (x86_64)
  • Kernel             → 5.15.0-88-generic
  • CPU                → AMD Ryzen 7 5800X
  • GPU                → NVIDIA RTX 3080 (CUDA 12.2)
  • RAM                → 32 GB
  • First run          → Yes (root)
  • Docker running     → ✅ Docker version: 24.0.5 🐳

🔧 Configuration:
  • Config file        → /etc/casai/casai.conf
  • MCP config         → /etc/casai/mcp.toml
  • Models dir         → /var/lib/casai/models
  • User config        → ~/.config/casai/config.toml
  • Secrets dir        → ~/.local/share/casai/secrets
  • Logs dir           → ~/.local/log/casai

🧠 System services:
  • Enabled            → Yes
  • Autostart          → Yes
  • Managed via        → systemd

🔐 Security:
  • Secrets stored     → Yes (hashed)
  • API keys encrypted → Yes
  • Certs generated    → No
  • Permission check   → OK (world-readable)

🧠 Models installed (based on system specs):
  • Coding Model       → codellama:7b-instruct (Quantized Q4)
  • Image Model        → stable-diffusion-xl-base-1.0
  • TTS Model          → en-us-kathleen-medium (Piper)

🐧 Binaries installed:
  • ollama             → /usr/local/bin/ollama (v0.1.27)
  • local-ai           → /usr/local/bin/local-ai (v2.2.0)
  • whisper-cpp        → /usr/local/bin/whisper (v1.5.1)
  • piper-tts          → /usr/local/bin/piper (v1.3.0)
  • opencode           → /usr/local/bin/opencode (v0.0.55)

🐳 Containers pulled and started:
  • open-webui         → {officialregistry}/open-webui (v0.3.5)
  • open-diffusion     → {officialregistry}/open-diffusion (v1.1.2) 🐳 Docker version: 24.0.5

🧠 AI Services:
  • Ollama             → http://localhost:11434
  • LocalAI            → http://localhost:64832
  • Open Web UI        → http://localhost:64666
  • Whisper            → http://localhost:64501
  • Piper TTS          → http://localhost:64502
  • OpenCode           → http://localhost:64503

🎉 You're all set!
```

---

## 📘 LICENSE

Use MIT License or Apache 2.0. All contributions must follow same license and coding standard.

---

## ✅ DONE

This is the entire definition of `casai`. Paste this into any chat, codegen tool, or LLM and it should produce the full working version including:
- CLI bash script
- README
- LICENSE
- Makefile
- All configs
- All directory scaffolds
- Full service control logic
- All model selection logic
- All binary/container logic

This description is 100% deterministic and sufficient to rebuild the project from scratch.

