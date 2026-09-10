# Infer from git remote — NEVER hardcode
PROJECTNAME := $(shell git remote get-url origin 2>/dev/null | sed -E 's|.*/([^/]+)(\.git)?$$|\1|' || basename "$$(pwd)")
PROJECTORG  := $(shell git remote get-url origin 2>/dev/null | sed -E 's|.*/([^/]+)/[^/]+(\.git)?$$|\1|' || basename "$$(dirname "$$(pwd)")")
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

.PHONY: build release docker test dev clean fmt lint vet

build:
	@mkdir -p $(BINDIR) $(GO_CACHE) $(GO_BUILD)
	$(GO_DOCKER) go build -buildvcs=false -trimpath -ldflags "$(LDFLAGS)" -o $(BINDIR)/$(PROJECTNAME) ./src

fmt:
	@mkdir -p $(GO_CACHE) $(GO_BUILD)
	@$(GO_DOCKER) sh -c "gofmt -l . && gofmt -w ."

vet:
	@mkdir -p $(GO_CACHE) $(GO_BUILD)
	@$(GO_DOCKER) go vet ./...

lint:
	@mkdir -p $(GO_CACHE) $(GO_BUILD)
	@$(GO_DOCKER) golangci-lint run ./...

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
		echo \"Tests complete - Coverage: \$$COVERAGE% (>= 60% required) \""

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
