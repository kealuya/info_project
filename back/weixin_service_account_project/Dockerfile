# 使用官方的 Go 镜像作为基础镜像
FROM golang:1.20 AS builder

ENV GOPROXY https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 添加 Go 模块文件并下载依赖
ADD go.mod .
RUN go mod download
RUN go mod tidy

# 复制源码并进行构建
COPY . .
COPY conf /app/conf
RUN go build -ldflags="-s -w" -o /app/main ./main.go

# 使用 Debian slim 镜像作为基础镜像
FROM debian:bookworm-slim

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
ENV TZ Asia/Shanghai

# 设置工作目录并复制构建的可执行文件
WORKDIR /app
COPY --from=builder /app/conf /app/conf
COPY --from=builder /app/main /app/main

# 暴露端口并设置启动命令
# EXPOSE 8000
CMD ["./main"]

# 如果有其他服务需要运行，可以使用下面的命令替换上面的 CMD
# CMD ["./other_service_api", "-f", "etc/other_service_api.yaml"]