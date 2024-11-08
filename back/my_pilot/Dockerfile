# 第一阶段：构建阶段
FROM golang:1.22-alpine AS builder

ENV GOPROXY=https://goproxy.cn,direct
# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download
RUN go mod tidy

# 复制源代码
COPY . .

# 编译
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 第二阶段：运行阶段
FROM alpine:latest

# 安装必要的包
#RUN apk add --no-cache ca-certificates tzdata


# 创建工作目录
WORKDIR /app

# 从builder阶段复制编译好的程序
COPY --from=builder /app/main .

# 复制配置文件（如果有的话）
COPY --from=builder /app/config ./config
COPY --from=builder /app/pkg/api_szjl ./pkg/api_szjl
# 确保文件具有执行权限
RUN chmod +x /app/main
# 暴露端口
EXPOSE 8080

# 启动命令
CMD ["./main"]
