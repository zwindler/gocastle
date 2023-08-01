#!/bin/bash

GOBIN=~/go/bin/
PATH=$PATH:$GOBIN

mkdir -p build/

# android
fyne-cross android -arch=arm64 -app-id fr.zwindler.gocastle
cp fyne-cross/dist/android-arm64/gocastle.apk build/

# linux
fyne-cross linux -arch=amd64,arm64 -app-id fr.zwindler.gocastle
cp fyne-cross/dist/linux-amd64/gocastle.tar.xz build/gocastle.amd64.tar.xz
cp fyne-cross/dist/linux-arm64/gocastle.tar.xz build/gocastle.arm64.tar.xz

# windows
fyne-cross windows -app-id fr.zwindler.gocastle
cp fyne-cross/dist/windows-amd64/gocastle.exe.zip build/gocastle.exe.zip

# macos TODO
# fyne-cross darwin -arch=amd64,arm64 --macosx-sdk-path /home/zwindler/sources/gocastle/bin/SDKs -app-id fr.zwindler.gocastle

# clean fyn-cross dir
rm -rf fyne-cross/

# unzip archives
cd build/
unzip gocastle.exe.zip && rm gocastle.exe.zip
mv build/usr/local/bin/gocastle.exe build/
tar xJf gocastle.amd64.tar.xz
mv usr/local/bin/gocastle ./gocastle.amd64
rm gocastle.amd64.tar.xz
tar xJf gocastle.arm64.tar.xz
mv usr/local/bin/gocastle ./gocastle.arm64
rm gocastle.arm64.tar.xz
