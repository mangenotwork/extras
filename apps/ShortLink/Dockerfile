FROM alpine:latest

ENV AppName=shortlink
RUN mkdir -p /app/logs
COPY ${AppName} /app/
COPY configs /app/configs

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
        && apk --no-cache add ca-certificates tzdata\
        && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
        && echo "Asia/Shanghai" > /etc/timezone \
        && apk del tzdata

WORKDIR /app

VOLUME /app/logs

EXPOSE 12333
EXPOSE 12334

LABEL maintainer = "2912882908@qq.com"

CMD ["sh","-c","./$AppName"]