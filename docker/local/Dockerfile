# ローカル開発環境
FROM golang:1.22.2-bullseye as debugger
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2

ENV DOCKERIZE_VERSION v0.7.0

RUN apt-get update \
    && apt-get install -y wget \
    && wget -O - https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz | tar xzf - -C /usr/local/bin \
    && apt-get autoremove -yqq --purge wget && rm -rf /var/lib/apt/lists/*

RUN go install github.com/rakyll/gotest@latest \
  && go install github.com/cosmtrek/air@latest \
  && go install github.com/go-delve/delve/cmd/dlv@latest
WORKDIR /go/src/github.com/take0fit/ddd-clean-temp
ENV CGO_ENABLED=0 GOOS=linux
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080

CMD ["dockerize", "-wait", "tcp://db:3306", "-timeout", "60s", "air", "-c", ".air.toml"]



#
#RUN go build -o main ./cmd/api-server/main.go
#RUN go build -o migration ./cmd/migration/migration.go
#
#RUN chmod +x main migration
#
## 実行ステージ
#FROM alpine:3.15
#
#WORKDIR /app
#
## ビルドステージから実行ファイルをコピー
#COPY --from=builder /app/main .
#COPY --from=builder /app/migration .