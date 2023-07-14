FROM node:17 AS web

WORKDIR /web
COPY . .

RUN npm install pnpm --location=global && pnpm config set registry http://registry.npm.taobao.org

RUN pnpm i && pnpm build:prod

FROM nginx:alpine
LABEL MAINTAINER="jianzhao.lv@baishan.com"

# 安装时区数据
RUN echo -e "http://mirrors.aliyun.com/alpine/v3.11/main\nhttp://mirrors.aliyun.com/alpine/v3.11/community" > /etc/apk/repositories \
    && apk add -U tzdata bash vim \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && mkdir /usr/share/nginx/html/dist \
    && rm -f /etc/nginx/conf.d/default.conf

COPY ./docker/nginx-default.conf /
COPY --from=web /web/dist /usr/share/nginx/html/dist
COPY --from=web /web/docker/docker-entrypoint.sh /entrypoint.sh
EXPOSE 8500
ENTRYPOINT ["bash", "/entrypoint.sh"]
