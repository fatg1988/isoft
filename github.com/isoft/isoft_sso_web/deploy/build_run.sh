#!/bin/bash

# 使用Dockerfile创建镜像, '.'代表使用当前目录的Dockerfile创建镜像
docker build -t isoft_sso_web_image:v1 .

# docker run命令可用于通过镜像运行容器
# -it标记可以用交互式模式启动该容器
# --rm标记可以在容器关闭后清理其中的内容
# --name isoft_sso_web_image_v1可以将容器命名为isoft_sso_web_image_v1
# -p 8080:8080标记使得容器可以通过8080端口访问
# -v /root/soft/install/gopath/src:go/src,可以将root/soft/install/gopath/src
# 从计算机映射至容器的go/src目录,这样可以确保在容器内部和外部均可访问这些开发文件
# -w 指定容器的工作目录
# isoft_sso_web_image:v1部分指定了容器内部使用的映像名称
docker run -it --rm --name isoft_sso_web_image_v1 -p 8080:8080 \
   -v /root/soft/install/gopath/src:/go/src -w /go/src/isoft_sso_web isoft_sso_web_image:v1

