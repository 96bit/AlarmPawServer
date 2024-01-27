# AlarmPawServer

## A push service backend, done with golang.

```golang


```

```shell
# 清除无用的镜像
docker builder prune --force
sudo docker build -t thurmantsao/alarm-paw-server:latest .
docker run -v ./deploy:/deploy -p 8080:8080  alarm-paw-server
docker rmi $(docker images  -q)
docker rm $(docker ps -a -q)
```