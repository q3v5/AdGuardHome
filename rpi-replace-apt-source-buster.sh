#!/bin/bash

set -e

echo "开始替换树莓派软件源到 tuna 镜像"

echo "replace /etc/apt/sources.list"
cat << EOF > /etc/apt/sources.list
deb [arch=armhf] http://mirrors.tuna.tsinghua.edu.cn/raspbian/raspbian/ bullseye main non-free contrib rpi
deb-src http://mirrors.tuna.tsinghua.edu.cn/raspbian/raspbian/ bullseye main non-free contrib rpi
deb [arch=arm64] http://mirrors.tuna.tsinghua.edu.cn/raspbian/multiarch/ bullseye main
EOF

echo "replace /etc/apt/sources.list.d/raspi.list"
cat << EOF > /etc/apt/sources.list.d/raspi.list
deb http://mirrors.tuna.tsinghua.edu.cn/raspberrypi/ bullseye main ui
EOF

echo "[电气罐头] 替换完成"
echo "使用文档：https://tech.biko.pub/#/tool/rpi-apt-sources"
