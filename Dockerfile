FROM golang:1.16

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release \
    PORT=8080

WORKDIR /app

COPY droneRepo main

ENTRYPOINT ["./main"]