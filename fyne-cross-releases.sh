#!/bin/bash

GOBIN=~/go/bin/
PATH=$PATH:$GOBIN

mkdir -p tmp/

# android
#fyne-cross android -arch=arm64 -app-id fr.zwindler.gocastle
cp fyne-cross/dist/android-arm64/gocastle.apk tmp/

# linux
#fyne-cross linux -arch=amd64,arm64 -app-id fr.zwindler.gocastle
cp fyne-cross/dist/linux-amd64/gocastle.tar.xz tmp/gocastle.amd64.tar.xz
cp fyne-cross/dist/linux-arm64/gocastle.tar.xz tmp/gocastle.arm64.tar.xz

# windows
#fyne-cross windows -app-id fr.zwindler.gocastle
cp fyne-cross/dist/windows-amd64/gocastle.exe.zip tmp/gocastle.exe.zip

# macos TODO
# fyne-cross darwin -arch=amd64,arm64 --macosx-sdk-path /home/zwindler/sources/gocastle/bin/SDKs -app-id fr.zwindler.gocastle

# clean fyn-cross dir
#rm -rf fyne-cross/

cd tmp/

# windows
unzip gocastle.exe.zip
mkdir -p ../dist/gocastle_windows_amd64_v1/
cp gocastle.exe ../dist/gocastle_windows_amd64_v1/gocastle.exe

# linux amd64
tar --strip-components=3 --wildcards -xJ '*/gocastle' -f gocastle.amd64.tar.xz
mkdir -p ../dist/gocastle_linux_amd64_v1/
mv gocastle ../dist/gocastle_linux_amd64_v1/gocastle
ls ../dist/gocastle_linux_amd64_v1/gocastle > toto.log

#linux arm64
tar --strip-components=3 --wildcards -xJ '*/gocastle' -f gocastle.arm64.tar.xz
mkdir -p ../dist/gocastle_linux_arm64
mv gocastle ../dist/gocastle_linux_arm64/gocastle
ls ../dist/gocastle_linux_arm64/gocastle >> toto.log

# android
mkdir -p ../dist/gocastle_android_arm64/gocastle
cp gocastle.apk ../dist/gocastle_android_arm64/gocastle

