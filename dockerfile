FROM golang:1.12.0-alpine3.9

WORKDIR /go/src/app

# Go Modules を使用するために必要な環境変数 GO111MODULE を on 
ENV GO111MODULE=on

# gitやホットリロードのインストール
RUN apk add --no-cache \
    alpine-sdk \
    git \
    && go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]