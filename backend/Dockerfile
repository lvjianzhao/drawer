FROM golang:alpine as builder

WORKDIR /server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o drawer .

FROM alpine:latest

LABEL MAINTAINER="jianzhao.lv@baishan.com"

# 安装时区数据
RUN echo -e "http://mirrors.aliyun.com/alpine/v3.11/main\nhttp://mirrors.aliyun.com/alpine/v3.11/community" > /etc/apk/repositories \
    && apk add -U tzdata bash vim\
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /opt/drawer

COPY --from=builder /server/drawer ./
COPY --from=builder /server/docker/config_tpl.yaml /tmp/
COPY --from=builder /server/docker/docker-entrypoint.sh /entrypoint.sh

EXPOSE 5558

ENTRYPOINT ["bash", "/entrypoint.sh"]
