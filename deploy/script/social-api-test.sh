#!/bin/bash
reso_addr='registry.cn-hangzhou.aliyuncs.com/0rz/social-api-dev'
tag='latest'

container_name="im-social-api-test"

# 停止容器
docker stop ${container_name}

# 删除容器
docker rm ${container_name}

# 删除镜像
docker rmi ${reso_addr}:${tag}

# 拉取镜像
docker pull ${reso_addr}:${tag}


# 如果需要指定配置文件的
# docker run -p 10001:8080 --network imooc_easy-chat -v /easy-chat/config/user-rpc:/user/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 8881:8881  --name=${container_name} -d ${reso_addr}:${tag}
