FROM alpine:latest

# app name
ENV Project=gozone

# mkdir directory
RUN mkdir -p /&{Project}/logs

# copy file to image
COPY src /${Project}/
COPY library /${Project}/library

# add timezone
RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.ustc.edu.cn/!g' /etc/apk/repositories \
&& apk --no-cache add ca-certificates tzdata\
&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Asia/Shanghai" > /etc/timezone \
&& apk del tzdata