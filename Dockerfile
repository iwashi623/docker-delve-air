FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod ./

RUN --mount=type=cache,target=/go/pkg/mod/,sharing=locked \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    go build -o /bin/myapp

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /bin/myapp /app/myapp

CMD ["/app/myapp"]
