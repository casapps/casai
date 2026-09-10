# Docker

casai ships as a single static binary; Docker is used to build, test, and
optionally run it in a container. There is no network listener in v1 (no
daemon/IPC mode — see `IDEA.md` → "App surfaces in scope"), so the
production compose service runs on demand rather than staying up as an
exposed service.

## Build

    make build      # debug build, host-arch, via casjaysdev/go:latest
    make release     # all platforms in PART 5's binary matrix
    make docker      # multi-arch runtime image via docker buildx

## Test

    make test        # go test ./... with coverage gate, inside Docker

## Run

    docker compose -f docker/docker-compose.yml run --rm casai --version

## GUI smoke test (X11 / Wayland)

    # X11
    xhost +SI:localuser:$(id -un)
    DISPLAY=$DISPLAY XAUTHORITY=$XAUTHORITY \
      docker compose -f docker/docker-compose.yml --profile gui run --rm gui
    xhost -SI:localuser:$(id -un)

    # Wayland
    WAYLAND_DISPLAY=$WAYLAND_DISPLAY XDG_RUNTIME_DIR=$XDG_RUNTIME_DIR \
      docker compose -f docker/docker-compose.yml --profile gui run --rm gui

Both backends must be exercised per AI.md PART 0 → "X11 AND Wayland Are
Both Required".

## Images

| File | Tag | Purpose |
|------|-----|---------|
| `Dockerfile` | `:latest`, `:VERSION` | production release binary |
| `Dockerfile.dev` | `:devel` | debug build, `MODE=devel` |

Toolchain image for all builds: `casjaysdev/go:latest` (AI.md PART 5 →
"Toolchain Image").
