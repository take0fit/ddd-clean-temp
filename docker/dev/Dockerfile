FROM golang:1.22.2-bullseye as builder
WORKDIR /go/src/github.com/take0fit/ddd-clean-temp
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build main.go

# 本番の実行環境
FROM debian:11.2-slim as executor
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /go/src/github.com/take0fit/ddd-clean-temp
COPY config/logger.json.tpl config/logger.json.tpl
COPY static /static

# SwaggerUI
COPY api/swagger-ui /api/swagger-ui
COPY api/output/openapi.json /api/output/openapi.json

WORKDIR /
COPY --from=builder /go/src/github.com/take0fit/ddd-clean-temp/main /ddd-clean-temp

EXPOSE 80
