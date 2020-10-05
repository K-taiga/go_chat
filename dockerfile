
FROM golang:latest

ENV GOBIN=$GOPATH/bin
WORKDIR /go
ADD ./server /go

CMD ["go", "run", "main.go"]