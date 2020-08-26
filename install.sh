#!/bin/bash

DOWNLAOD_URL="https://github.com/sakz/fanyi/releases/download"

VERSION_CHECK="https://api.github.com/repos/sakz/fanyi/releases/latest"

UPDATE=0

[[ -e /usr/local/bin/fanyi ]] && UPDATE=1

RED="31m"
GREEN="32m"
YELLOW="33m"
BLUE="36m"
FUCHSIA="35m"

colorEcho(){
    COLOR=$1
    echo -e "\033[${COLOR}${@:2}\033[0m"
}

checkSys() {
    #检查是否为Root
#    [ $(id -u) != "0" ] && { colorEcho ${RED} "Error: You must be root to run this script"; exit 1; }
    if [[ $(uname -m 2> /dev/null) != x86_64 ]]; then
        colorEcho $YELLOW "Please run this script on x86_64 machine."
        exit 1
    fi

    if [[ `command -v apt-get` ]];then
        PACKAGE_MANAGER='apt-get'
    elif [[ `command -v dnf` ]];then
        PACKAGE_MANAGER='dnf'
    elif [[ `command -v yum` ]];then
        PACKAGE_MANAGER='yum'
    else
        colorEcho $RED "Not support OS!"
        exit 1
    fi

    # 缺失/usr/local/bin路径时自动添加
    [[ -z `echo $PATH|grep /usr/local/bin` ]] && { echo 'export PATH=$PATH:/usr/local/bin' >> /etc/profile; source /etc/profile; }
}

install(){
    if [[ $UPDATE == 1 ]];then
        rm -rf /usr/local/bin/fanyi
    fi
    LASTEST_VERSION=$(curl -H 'Cache-Control: no-cache' -s "$VERSION_CHECK" | grep 'tag_name' | cut -d\" -f4)
    echo "正在下载程序`colorEcho $BLUE $LASTEST_VERSION`版本..."
#   https://github.com/sakz/fanyi/releases/download/v0.0.4/fanyi_0.0.4_Linux_x86_64.tar.gz
    cd /usr/local/bin/
    echo "$DOWNLAOD_URL/$LASTEST_VERSION/fanyi_${LASTEST_VERSION#*v}_Linux_x86_64.tar.gz"
    curl -LO "$DOWNLAOD_URL/$LASTEST_VERSION/fanyi_${LASTEST_VERSION#*v}_Linux_x86_64.tar.gz"
    tar zxf fanyi_${LASTEST_VERSION#*v}_Linux_x86_64.tar.gz
    rm -rf fanyi_${LASTEST_VERSION#*v}_Linux_x86_64.tar.gz README.md
    cd
    echo "安装完成，在终端输入 `colorEcho $BLUE fanyi` 体验"
}

main(){
  checkSys
  install
}
main
