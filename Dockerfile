FROM golang:1.17-alpine AS builder
# 配置alpine使用清华源
RUN echo -e "https://mirror.tuna.tsinghua.edu.cn/alpine/latest-stable/main\nhttps://mirror.tuna.tsinghua.edu.cn/alpine/latest-stable/community" > /etc/apk/repositories
# 安装git
RUN apk update && apk add --no-cache git gcc musl-dev
# 设置cibn.job为工作目录（RUN执行的命令会在改目录下操作）
#WORKDIR $GOPATH/src/x/x/x
# /home/duneng/Graduation-project
WORKDIR /home/duneng/Graduation-project
# 将当前目录（也就是执行docker build xxx 的目录）下的文件拷贝到工作目录的上级目录下
COPY go.mod go.mod
COPY go.sum go.sum

# 配置docker镜像的系统环境：启用go modules
RUN go env -w GO111MODULE=on
# 配置go modules代理，package将从该开源代理拉取，避免GFW的问题
RUN go env -w GOPROXY=https://proxy.golang.com.cn,direct
# 获取所有packagei
RUN go mod download -x
RUN go mod tidy
# 构建项目，输出到镜像的指定目录下
COPY  . .
RUN go build -ldflags="-w -s" -o /home/duneng/Graduation-project/apiserver
# 使用一个空的镜像打包发布后的go项目，以达到镜像体积的最小化
FROM alpine:3.14
#    echo "http://mirrors.cloud.aliyuncs.com/alpine/v3.14/main/" > /etc/apk/repositories
# 从上面的镜像中拷贝编译后的程序到当前镜像x的指定位置

WORKDIR /Graduation-project
COPY --from=builder /home/duneng/Graduation-project/apiserver .
COPY --from=builder /home/duneng/Graduation-project/static static/
COPY --from=builder /home/duneng/Graduation-project/conf conf/
COPY --from=builder /home/duneng/Graduation-project/build/docker-entrypoint.sh /usr/local/bin/


RUN echo "http://mirrors.aliyun.com/alpine/v3.14/main/" > /etc/apk/repositories && \
    apk --no-cache add tzdata bash coreutils && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    chmod u+s /bin/ping && \
    chmod +x /usr/local/bin/docker-entrypoint.sh && \
    chmod +x apiserver && \
    rm -rf /var/cache/apk/*
#COPY lcc-apiserver/web/ /home/admin/lcc-apiserver/web/
ENTRYPOINT ["docker-entrypoint.sh"]
# 暴露镜像的80端口
EXPOSE 80
# 运行
