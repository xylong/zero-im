IM即时通讯

### 目录
```
.
├── Makefile                            执行makefile文件命令的总入口
├── README.md
├── apps                                服务
│   └── user                      用户服务
│       ├── api
│       ├── exec.sh               执行shell命令，比如框架命令生成api或rpc代码
│       └── rpc
├── deploy
│   ├── dockerfile
│   │   └── Dockerfile_user_rpc_dev 构建用户服务镜像
│   └── mk
│       └── user-rpc.mk             将用户服务镜像推到阿里云
├── docker-compose.yaml
├── go.mod
├── go.sum
└── pkg
```


### user 部署
1. 先将程序编译成二进制可执行的文件
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/user-rpc ./apps/user/rpc/user.go

2. 然后根据二进制文件构建成镜像文件
docker build -t user-rpc -f ./dockerfile_rpc .

3. 再修改构建的镜像标签
$ docker tag user-rpc registry.cn-hangzhou.aliyuncs.com/easy-chat/user-rpc-test:latest

4. 然后推送到阿里云上
$ docker push registry.cn-hangzhou.aliyuncs.com/easy-chat/user-rpc-test:latest

5. 在部署的时候拉取下来构建容器运行即可

以上步骤放在user-rpc.mk中，方便复用，user-rpc.mk通过Dockerfile_user_rpc_dev来构建镜像

### 命令
```makefile
# 查看本机ip，方便配置etcd的宿主机ip
ifconfig | grep "inet " | awk '{print $2}'

# 执行Makefile中的release-test命令，编译二进制文件，构建镜像推送到阿里云镜像仓库
make release-test

# 构建运行环境：mysql+redis+redis
docker-compose up -d

# 删除所有运行容器后重新拉取镜像运行
make install-server
```

### jaeger
```shell
go get github.com/opentracing/opentracing-go
go get github.com/uber/jaeger-client-go

```
