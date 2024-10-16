FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod ./

RUN go mod download
COPY . .

RUN go build -o myapp

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/myapp .

CMD ["./myapp"]
