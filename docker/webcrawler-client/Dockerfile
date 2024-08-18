FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod download && go build -o webcrawler-client ./cmd/webcrawler-client

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/webcrawler-client .

ENTRYPOINT ["./webcrawler-client"]