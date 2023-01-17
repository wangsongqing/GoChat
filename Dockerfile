FROM golang:alpine AS builder

LABEL stage=gobuilder

# ENV 用来设置环境变量
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# RUN容器构建时需要运行的命令
# RUN是在 docker build时运行。
RUN apk update --no-cache && apk add --no-cache tzdata

# WORKDIR 登录终端的落脚点，就是进入容器默认的路径
WORKDIR /build

# ADD 将宿主机目录下的文拷贝进镜像且回自动处理URL和解压tar压缩包
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/main main.go


FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY  .env /app
COPY --from=builder /app/main /app/main

# CMD 指定容器启动后要干的事情
# Dockerfile 中可以有多个 CMD 指令，但只有最后一个生效，CMD 会被 docker run 之后的参数替换
# CMD是在docker run 时运行。
CMD ["./main"]
