# 第一阶段：构建阶段
FROM golang:1.21 AS build

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件，以便在下载依赖项时缓存
COPY go.mod go.sum ./

# 下载依赖项
RUN go mod download

# 将应用程序的代码复制到容器中
COPY . .

# 构建应用程序
RUN go build -o main .


FROM alpine:latest

WORKDIR /app

RUN set -ex \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk add ca-certificates \
    && apk add --no-cache tzdata\
    && rm -rf /var/lib/apk/lists/*

COPY --from=build /app/main /app


VOLUME /data

EXPOSE 8080

WORKDIR /

CMD ["./AlarmPawServer"]


# sudo docker build -t alarm-paw-server:latest .
# docker run -v ./data:/data -p 8080:8080  alarm-paw-server
