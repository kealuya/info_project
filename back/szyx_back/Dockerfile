FROM golang:1.17-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY conf /app/conf
RUN go build -ldflags="-s -w" -o /app/szyx_back main.go


FROM alpine

#换成阿里云资源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/conf /app/conf
COPY --from=builder /app/szyx_back /app/szyx_back

EXPOSE 7001

CMD ["./szyx_back"]
