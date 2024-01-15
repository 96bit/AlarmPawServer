#!/bin/zsh

# 编译linux版本
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./data/AlarmPawServer  main.go || echo "编译linux版本失败"

##!/bin/zsh
#docker pull --platform linux/amd64  alpine:latest
#docker buildx build --platform linux/amd64 -t AlarmPawServer .
#docker tag alarm-paw-server thurmantsao/alarm-paw-server:v1.0.3
#docker push thurmantsao/alarm-paw-server:v1.0.3