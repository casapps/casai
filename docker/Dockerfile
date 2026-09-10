# =============================================================================
# Build Stage - Compile Go binary
# =============================================================================
FROM casjaysdev/go:latest AS builder

ARG TARGETARCH
ARG VERSION=devel
ARG BUILD_DATE
ARG BUILD_EPOCH
ARG COMMIT_ID

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build \
    -buildvcs=false -trimpath \
    -ldflags "-s -w \
      -X main.Version=${VERSION} \
      -X main.CommitID=${COMMIT_ID} \
      -X 'main.BuildEpoch=${BUILD_EPOCH}' \
      -X main.OfficialSite=" \
    -o /out/casai ./src

# =============================================================================
# Runtime Stage
# =============================================================================
FROM alpine:latest

ARG VERSION=devel
ARG BUILD_DATE
ARG COMMIT_ID
ARG LICENSE=MIT

RUN apk add --no-cache bash tini ca-certificates

COPY --from=builder /out/casai /usr/local/bin/casai
COPY docker/rootfs/ /
COPY docker/Dockerfile /usr/local/share/casai/Dockerfile

RUN chmod 755 /usr/local/bin/casai /usr/local/bin/entrypoint.sh

STOPSIGNAL SIGTERM

HEALTHCHECK --start-period=10s --interval=5m --timeout=15s --retries=3 \
    CMD /usr/local/bin/casai --version || exit 1

ENTRYPOINT [ "tini", "-p", "SIGTERM", "--", "/usr/local/bin/entrypoint.sh" ]
CMD [ "casai", "--help" ]
