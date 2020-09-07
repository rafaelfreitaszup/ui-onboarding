#!/bin/bash

SYSTEM=$(uname -s)

cd ~
wget https://sciter.com/sdk/sciter-sdk.zip

unzip sciter-sdk.zip

case "$SYSTEM" in
Linux*)
    installSciterLinux
    ;;
Darwin*)
    config "MacOS"
    installSciterMacOS
    ;;
*)
    printf "Unknown operating system.\\n"
    ;;
esac

installSciterLinux() {
    cd sciter-sdk/bin.lnx/x64

    export LIBRARY_PATH=$PWD

    echo $PWD >>libsciter.conf

    sudo cp libsciter.conf /etc/ld.so.conf.d/

    sudo ldconfig

    ldconfig -p | grep sciter

    sudo apt install build-essential -y

    sudo apt-get install build-essential libgtk-3-dev -y

    dpkg -L libgtk-3-dev | grep '\.pc'

    pkg-config --modversion gtk+-3.0

    go get -x github.com/sciter-sdk/go-sciter
}

installSciterMacOS() {
    cd sciter-sdk/bin.osx/
    export DYLD_LIBRARY_PATH=$PWD
}