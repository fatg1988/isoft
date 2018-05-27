# 第一行必须是基础镜像
# 使用Go的官方映像作为基础映像.该映像的$GOPATH值已被设置为/go,所有安装在/go/src的程序包都能通过go命令访问
FROM golang:latest

# 维护者信息
MAINTAINER zhourui  389093982@qq.com

# 使用 WORKDIR 指令可以来指定工作目录(或者称为当前目录),以后各层的当前目录就被改为指定的目录,如该目录不存在,WORKDIR 会帮你建立目录
# WORKDIR $GOPATH/src
# Dockerfile中的COPY指令和ADD指令都可以将主机上的资源复制或加入到容器镜像中,都是在构建镜像的过程中完成的
# COPY指令和ADD指令的唯一区别在于是否支持从远程URL获取资源.COPY指令只能从执行docker build所在的主机上读取资源并复制到镜像中.
# 而ADD指令还支持通过URL从远程服务器读取资源并复制到镜像中.满足同等功能的情况下.推荐使用COPY指令.ADD指令更擅长读取本地tar文件并解压缩
# ADD ./arvonblog/ $GOPATH/src/arvonblog

# 安装beego程序包和bee工具
RUN go get github.com/astaxie/beego && go get github.com/astaxie/beego/orm && go get github.com/beego/bee
# -u 强制使用网络去更新包和它的依赖包
RUN go get -u github.com/389093982/isoft

# 通过开发计算机上容器的8080端口暴露该应用程序
EXPOSE 8080

# 使用bee命令开始对我们的应用程序进行实时重载
CMD ["bee", "run"]