# Building stage
FROM golang:1.20.3 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOPROXY="https://goproxy.cn,direct"

WORKDIR /build

COPY . .

ENV TIMEZONE "Asia/Shanghai"

RUN go build -v -o balancer ./

FROM charlie1380/nettest:latest

COPY --from=builder /build/balancer /opt/balancer

# 这里为啥不能写 Entrypoint 啊寄
CMD ["/opt/balancer"]
