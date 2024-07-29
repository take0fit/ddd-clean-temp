# ローカル開発環境
FROM golang:1.22.2-bullseye as debugger
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2

ENV DOCKERIZE_VERSION v0.7.0

RUN apt-get update \
    && apt-get install -y wget \
    && wget -O - https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz | tar xzf - -C /usr/local/bin \
    && apt-get autoremove -yqq --purge wget && rm -rf /var/lib/apt/lists/*

RUN go install github.com/rakyll/gotest@latest \
  && go install github.com/air-verse/air@latest \
  && go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /go/src/github.com/take0fit/ddd-clean-temp
ENV CGO_ENABLED=0 GOOS=linux

COPY go.mod go.sum ./
RUN go mod download
COPY . .

EXPOSE 8080

CMD ["dockerize", "-wait", "tcp://auth_db:5432", "-wait", "tcp://user_db:5432", "-timeout", "60s", "air", "-c", ".air.toml"]
