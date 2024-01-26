FROM alpine:latest

RUN set -ex \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk add ca-certificates \
    && apk add --no-cache tzdata\
    && rm -rf /var/lib/apk/lists/*

COPY ./data/AlarmPawServer /AlarmPawServer


VOLUME /data

EXPOSE 8080

WORKDIR /

CMD ["./AlarmPawServer", "-c", "/data/config.yaml"]



