# 使用官方的 Go 镜像作为基础镜像
FROM golang:1.20 AS builder

# 因为自建镜像是linux/amd64环境下编译的，所以在编译go程序时，也要选择linux/amd64环境
# 不然会发生执行时exec ./main: no such file or directory 错误
# mac是linux/arm64 ，云服务器是linux的linux/amd64环境
#FROM --platform=linux/amd64 golang:1.20 AS builder

ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 添加 Go 模块文件并下载依赖
ADD go.mod .
RUN go mod download
RUN go mod tidy

# 复制源码并进行构建
COPY . .
#COPY conf /app/conf
RUN go build -ldflags="-s -w" -o /app/main ./main.go
# 添加调试命令查看文件是否存在
RUN ls -la /app/main


# 使用 Debian slim 镜像作为基础镜像
FROM debian:bookworm-slim

# renhao：使用自建镜像（linux/amd64环境），作为容器镜像，省的下载以下内容，加快构建
#FROM kealuya/szht-go-service-oracle-ffmpeg-cacertificates-base:v1.0

# 安装 Oracle Instant Client 的依赖
#RUN apt-get update && apt-get install -y libaio1 wget unzip

# 下载并解压 Oracle Instant Client
# 注意：Oracle 官网的下载链接可能会变化，请根据需要更新下面的链接和版本号
#RUN wget https://download.oracle.com/otn_software/linux/instantclient/191000/instantclient-basic-linux.x64-19.10.0.0.0dbru.zip \
#    && unzip instantclient-basic-linux.x64-19.10.0.0.0dbru.zip -d /opt/oracle \
#    && rm instantclient-basic-linux.x64-19.10.0.0.0dbru.zip

# 设置 Oracle Instant Client 相关的环境变量
#ENV LD_LIBRARY_PATH=/opt/oracle/instantclient_19_10:$LD_LIBRARY_PATH
#ENV TNS_ADMIN=/opt/oracle/instantclient_19_10
#ENV ORACLE_BASE=/opt/oracle/instantclient_19_10
#ENV ORACLE_HOME=/opt/oracle/instantclient_19_10

# 安装 ffmpeg 如果需要语音处理
#RUN apt-get update && apt-get install -y ffmpeg
# 安装 ca-certificates 如果需要容器内请求https网站（证书认证），需要安装
RUN apt-get update && apt-get install -y ca-certificates


#ENV TZ Asia/Shanghai

# 设置工作目录并复制构建的可执行文件
WORKDIR /app
COPY --from=builder /app/conf /app/conf
COPY --from=builder /app/main /app/main

# 添加调试命令查看文件是否存在
RUN ls -la /app/main
# 确保文件具有执行权限
RUN chmod +x /app/main
# 暴露端口并设置启动命令
EXPOSE 8080
CMD ["./main"]

# 如果有其他服务需要运行，可以使用下面的命令替换上面的 CMD
# CMD ["./other_service_api", "-f", "etc/other_service_api.yaml"]