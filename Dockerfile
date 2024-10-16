FROM golang:1.23 AS builder
WORKDIR /app

ARG APP_ENV=production

RUN if [ "$APP_ENV" = "development" ]; then \
    go install github.com/air-verse/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest; \
    fi

RUN --mount=type=cache,target=/go/pkg/mod/,sharing=locked \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    go build -o /bin/myapp

# --------------------prod--------------------
FROM debian:bookworm-slim AS prod
WORKDIR /app
COPY --from=builder /bin/myapp /app/myapp

CMD ["/app/myapp"]

# --------------------dev--------------------
FROM debian:bookworm-slim AS dev
WORKDIR /app
COPY --from=builder /bin/myapp /app/myapp

COPY --from=builder /go/bin/air /usr/local/bin/air
COPY --from=builder /go/bin/dlv /usr/local/bin/dlv

CMD ["/app/myapp"]
