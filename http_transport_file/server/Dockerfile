# 基础镜像，基于golang的alpine镜像构建--编译阶段
FROM --platform=linux/amd64 golang:1.17.3-alpine AS builder
# 作者
MAINTAINER zhaomin
# 全局工作目录
WORKDIR /go/srv
# 把运行Dockerfile文件的当前目录所有文件复制到目标目录
COPY . /go/srv
# 环境变量
#  用于代理下载go项目依赖的包
ENV GOPROXY https://goproxy.cn,direct
# 编译，关闭CGO，防止编译后的文件有动态链接，而alpine镜像里有些c库没有，直接没有文件的错误
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" main.go


# 使用alpine这个轻量级镜像为基础镜像--运行阶段
FROM alpine AS runner
# 全局工作目录
WORKDIR /go/srv
# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /go/srv/main .
# 复制编译阶段里的config文件夹到目标目录
# COPY --from=builder /go/srv/config ./config
# 将时区设置为东八区
RUN echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && echo Asia/Shanghai > /etc/timezone \
    && apk del tzdata \
    && echo /bin/sh > /etc/shells \
    && mkdir /lib64 \
    && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
# 需暴露的端口
EXPOSE 9999
# 可外挂的目录
VOLUME ["/go/srv/config","/go/srv/log", "/go/srv/data"]
# docker run命令触发的真实命令(相当于直接运行编译后的可运行文件)
ENTRYPOINT ["./main"]