#!/bin/bash

GOBIN=~/go/bin/
PATH=$PATH:$GOBIN

mkdir -p dist/gocastle_linux_amd64_v1/ dist/gocastle_android_amd64 dist/gocastle_android_arm64/ dist/gocastle_windows_amd64_v1/

fyne-cross android -arch=arm64 -app-id fr.zwindler.gocastle
cp fyne-cross/dist/android-arm64/gocastle.apk dist/gocastle_android_amd64/

fyne-cross linux -arch=amd64,arm64 -app-id fr.zwindler.gocastle
cp fyne-cross/dist/linux-amd64/gocastle dist/gocastle_linux_amd64_v1/gocastle
cp fyne-cross/dist/linux-arm64/gocastle dist/gocastle_linux_amd64_v1/gocastle

fyne-cross windows -app-id fr.zwindler.gocastle
cp fyne-cross/dist/windows-amd64/gocastle.exe.zip dist/gocastle_windows_amd64_v1/gocastle.exe.zip
cd dist/gocastle_windows_amd64_v1/
unzip gocastle.exe.zip
cd ../..

rm -rf fyne-cross/