FROM alpine:latest

ENV AppName=push
RUN mkdir -p /app/logs
COPY ${AppName} /app/
COPY conf /app/configs

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
        && apk --no-cache add ca-certificates tzdata\
        && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
        && echo "Asia/Shanghai" > /etc/timezone \
        && apk del tzdata

WORKDIR /app
VOLUME /app/logs

EXPOSE 1241
EXPOSE 1242
EXPOSE 1243
EXPOSE 1244

LABEL maintainer = "2912882908@qq.com"

CMD ["sh","-c","./$AppName"]