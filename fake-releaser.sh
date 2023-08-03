#!/bin/bash

GOBIN=~/go/bin/
PATH=$PATH:$GOBIN

build_linux_amd64() {
    touch buildlock
    fyne-cross linux -arch=amd64,arm64 -app-id fr.zwindler.gocastle
    mkdir -p dist/gocastle_linux_amd64_v1/
    cp fyne-cross/bin/linux-amd64/gocastle dist/gocastle_linux_amd64_v1/gocastle
    rm buildlock
}

build_linux_arm64() {
    sleep 1

    # build is done by build_linux_amd64, wait for it to finish
    timeout=120
    while [[ -e buildlock ]] && [[ $timeout -gt 0 ]]; do
        sleep 1
        ((timeout--))
    done

    if [[ -e buildlock ]]; then
        echo "Timeout: Lock file not removed after 20 seconds."
        exit 1
    fi

    mkdir -p dist/gocastle_linux_arm64/
    cp fyne-cross/bin/linux-arm64/gocastle dist/gocastle_linux_arm64/gocastle
}

build_android() {
    fyne-cross android -arch=arm64 -app-id fr.zwindler.gocastle
    mkdir -p dist/gocastle_android_arm64/
    cp fyne-cross/dist/android-arm64/gocastle.apk dist/gocastle_android_arm64/
}

build_windows() {
    fyne-cross windows -app-id fr.zwindler.gocastle
    mkdir -p dist/gocastle_windows_amd64_v1/
    cp fyne-cross/bin/windows-amd64/gocastle.exe dist/gocastle_windows_amd64_v1/
}

if [[ "$4" == *"/gocastle_linux_amd64_v1/gocastle" ]]; then
    build_linux_amd64
elif [[ "$4" == *"/gocastle_linux_arm64/gocastle" ]]; then
    build_linux_arm64
elif [[ "$4" == *"/gocastle_android_arm64/gocastle" ]]; then
    build_android
elif [[ "$4" == *"/gocastle_windows_amd64_v1/gocastle.exe" ]]; then
    build_windows
else
    echo "Invalid or unsupported path argument."
    exit 1
fi
