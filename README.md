# AlarmPawServer
程序思路和部分代码均来自 [BARK](https://github.com/Finb/bark-server)
## 一个基于golang的推送服务后端，主要用于推送消息给客户端。
程序均采用[BARK](https://github.com/Finb/bark-server)的接口模式，重新实现了一遍，采用了gin框架，方便后续扩展。


```shell

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  main  main.go || echo "编译linux版本失败"

```
## Docker部署
```shell
docker run -d --name alarm-paw-server -p 8080:8080 -v ./data:/data  --restart=always  thurmantsao/alarm-paw-server:latest
```

## Docker-compose部署

### 配置
* 复制项目中的/deploy文件夹到服务器上，然后执行以下命令即可。
* 必须有/data/config.yaml 的配置文件，否则无法启动，文件中的配置项，可以根据自己的需求进行修改。
```yaml
system:
  name: "AlarmPawServer"
  user: ""         # 用户名 非必填
  password: ""    # 密码  非必填
  host: "0.0.0.0"  # 服务监听地址
  port: "8080"   # 服务监听端口 docker-compose中的端口映射必须与此端口一致
  mode: "release"   # debug,release,test
  dbType: "default" # default,mysql 
  dbPath: "/data" # 数据库文件存放路径 

mysql: # 仅在 dbType: "mysql" 时有效
  host: "localhost"
  port: "3306"
  user: "root"
  password: "root"

apple: # 复制项目中的配置，不需要修改，仅在自己编译app时需要修改
    keyId: 
    teamId: 
    topic: 
    apnsPrivateKey:


```
### 启动
```shell
docker-compose up -d
```


