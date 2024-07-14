#!/bin/bash

# chmod +x ./auto.sh

# 先决条件说明（可选）
# 0. 需要在本机执行以下语句，将密钥同步到云端服务器，这样可以无需密码验证(这部分不执行也行，每次需要输入服务器密码)，直接进行同步
# 1. 本机生成密码（公钥）
#    ssh-keygen -t rsa
#    ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
# 2. 将密钥同步至云端服务器
#    ssh-copy-id root@122.9.41.45
# 3. 此时会需要输入云端服务器密码
# 4. 在云端服务器上及本机上安装rsync
#    在 Ubuntu 或 Debian 系统上安装 rsync
#    在 Ubuntu 或 Debian 系统上，可以使用 apt-get 包管理器安装 rsync。运行以下命令：
#    sudo apt-get update
#    sudo apt-get install rsync
#    在 CentOS 或 Red Hat 系统上安装 rsync
#    在 CentOS 或 Red Hat 系统上，可以使用 yum 包管理器安装 rsync。运行以下命令：
#    sudo yum install rsync
#    在 macOS 上安装 rsync
#    在 macOS 上，可以使用 Homebrew 或 MacPorts 安装 rsync。如果你使用 Homebrew，可以运行以下命令：
#    brew install rsync

# 设置变量
SERVER_ADDR="122.9.41.45"
SOURCE_DIR="/Users/kealuya/mywork/my_git/info_project/back/weixin_service_account_project/"
WINDOWS_SOURCE_DIR="C:/path/to/your/source/"
DEST_DIR="root@$SERVER_ADDR:/root/weixin_service_account_project"
PORT=8086

# 检查操作系统类型并同步代码
if [[ "$OSTYPE" == "linux-gnu"* || "$OSTYPE" == "darwin"* ]]; then
    echo "当前操作系统是类 Unix 系统。"
    echo "开始同步代码到云服务器..."
    rsync -avz $SOURCE_DIR $DEST_DIR
    echo "代码同步完成。"
elif [[ "$OSTYPE" == "cygwin" || "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
    echo "当前操作系统是 Windows 系统。"
    echo "开始同步代码到云服务器..."
    scp -r $WINDOWS_SOURCE_DIR $DEST_DIR
    echo "代码同步完成。"
else
    echo "无法识别的操作系统类型: $OSTYPE"
    exit 1
fi

echo "开始在云服务器上构建和运行 Docker 镜像..."
ssh root@$SERVER_ADDR << EOF
echo "进入目录 /root/weixin_service_account_project"
cd /root/weixin_service_account_project

echo "停止容器 weixin_service（如果存在）..."
docker stop weixin_service || true

echo "删除容器 weixin_service（如果存在）..."
docker rm weixin_service || true

echo "删除镜像 weixin_service_image（如果存在）..."
docker image rm weixin_service_image || true

echo "构建 Docker 镜像 weixin_service_image..."
docker build -t weixin_service_image -f ./Dockerfile . || true

echo "运行 Docker 容器 weixin_service..."
docker run -itd -p $PORT:$PORT --name weixin_service weixin_service_image || true

echo "Docker 容器 weixin_service 运行成功。"
EOF

echo "脚本执行完成。"
