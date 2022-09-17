FROM golang:1.18.1-alpine
MAINTAINER https://github.com/magicnana999

WORKDIR /go-training

ARG SV
COPY . .


ENV GOPROXY=https://goproxy.cn,direct \

    GO111MODULE=on \

    CGO_ENABLED=0 \

    GOOS=linux \

    GOARCH=amd64 \

    SERVER_VERSION=$SV

RUN go mod download
RUN go build -o ws main.go
RUN chmod +x ws
ENTRYPOINT env & ./ws
