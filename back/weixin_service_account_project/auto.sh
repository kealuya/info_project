#!/bin/bash

# chmod +x ./auto.sh

#0，需要先在本机执行一下语句，将密钥同步到云端服务器，这样就可以无需密码验证，直接进行同步
#1，本机生成密码（公钥）
    #ssh-keygen -t rsa
    #ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
#2，将密钥同步至云端服务器
    #ssh-copy-id root@122.9.41.45
#3，此时会需要输入云端服务器密码
#4，在云端服务器上及本机上安装rsync
    #在 Ubuntu 或 Debian 系统上安装 rsync
    #在 Ubuntu 或 Debian 系统上，可以使用 apt-get 包管理器安装 rsync。运行以下命令：
    #
    #sudo apt-get update
    #sudo apt-get install rsync
    #
    #在 CentOS 或 Red Hat 系统上安装 rsync
    #在 CentOS 或 Red Hat 系统上，可以使用 yum 包管理器安装 rsync。运行以下命令：
    #
    #sudo yum install rsync
    #
    #在 macOS 上安装 rsync
    #在 macOS 上，可以使用 Homebrew 或 MacPorts 安装 rsync。如果你使用 Homebrew，可以运行以下命令：
    #
    #brew install rsync

# 同步代码到云服务器（rsync只同步变更的文件）
echo "开始同步代码到云服务器..."
rsync -avz /Users/kealuya/mywork/my_git/testDockerFile_oracle_service/ root@124.70.48.30:/root/testDockerFile_oracle_service
echo "代码同步完成。"

echo "开始在云服务器上构建和运行 Docker 镜像..."
ssh root@124.70.48.30 << 'EOF'
echo "进入目录 /root/testDockerFile_oracle_service"
cd /root/testDockerFile_oracle_service

echo "停止容器 testmygo（如果存在）..."
docker stop testmygo || true

echo "删除容器 testmygo（如果存在）..."
docker rm testmygo || true

echo "删除镜像 ren_test_go_oracle（如果存在）..."
docker image rm ren_test_go_oracle || true

echo "构建 Docker 镜像 ren_test_go_oracle..."
docker build -t ren_test_go_oracle -f ./Dockerfile  . || true

echo "运行 Docker 容器 testmygo..."
docker run -itd -p 8100:8000 --name testmygo ren_test_go_oracle || true

echo "Docker 容器 testmygo 运行成功。"
EOF

echo "脚本执行完成。"