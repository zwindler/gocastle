#!/bin/bash

GOBIN=~/go/bin/
PATH=$PATH:$GOBIN

build_linux_amd64() {
    touch buildlock
    fyne-cross linux -arch=amd64,arm64 -app-id fr.zwindler.gocastle
    fyne-cross android -arch=arm64 -app-id fr.zwindler.gocastle
    fyne-cross windows -app-id fr.zwindler.gocastle
    fyne-cross darwin -arch=amd64,arm64 --macosx-sdk-path /home/zwindler/sources/gocastle/bin/SDKs.12.4/MacOSX.sdk -app-id fr.zwindler.gocastle
    mkdir -p dist/gocastle_linux_amd64_v1/
    cp fyne-cross/bin/linux-amd64/gocastle dist/gocastle_linux_amd64_v1/gocastle
    rm buildlock
}

build_linux_arm64() {
    sleep 1

    # build is done by build_linux_amd64, wait for it to finish
    timeout=240
    while [[ -e buildlock ]] && [[ $timeout -gt 0 ]]; do
        sleep 1
        ((timeout--))
    done

    if [[ -e buildlock ]]; then
        echo "Timeout: Lock file not removed after 240 seconds."
        exit 1
    fi

    mkdir -p dist/gocastle_linux_arm64/
    cp fyne-cross/bin/linux-arm64/gocastle dist/gocastle_linux_arm64/gocastle
}

build_android() {
    sleep 1

    # build is done by build_linux_amd64, wait for it to finish
    timeout=240
    while [[ -e buildlock ]] && [[ $timeout -gt 0 ]]; do
        sleep 1
        ((timeout--))
    done

    if [[ -e buildlock ]]; then
        echo "Timeout: Lock file not removed after 240 seconds."
        exit 1
    fi

    mkdir -p dist/gocastle_android_arm64/
    cp fyne-cross/dist/android-arm64/gocastle.apk dist/gocastle_android_arm64/gocastle
}

build_windows() {
    sleep 1

    # build is done by build_linux_amd64, wait for it to finish
    timeout=240
    while [[ -e buildlock ]] && [[ $timeout -gt 0 ]]; do
        sleep 1
        ((timeout--))
    done

    if [[ -e buildlock ]]; then
        echo "Timeout: Lock file not removed after 240 seconds."
        exit 1
    fi

    mkdir -p dist/gocastle_windows_amd64_v1/
    cp fyne-cross/bin/windows-amd64/gocastle.exe dist/gocastle_windows_amd64_v1/
}

build_darwin_amd64() {
    sleep 1

    # build is done by build_linux_amd64, wait for it to finish
    timeout=240
    while [[ -e buildlock ]] && [[ $timeout -gt 0 ]]; do
        sleep 1
        ((timeout--))
    done

    if [[ -e buildlock ]]; then
        echo "Timeout: Lock file not removed after 240 seconds."
        exit 1
    fi

    mkdir -p dist/gocastle_darwin_amd64_v1/
    cp fyne-cross/bin/darwin-amd64/gocastle dist/gocastle_darwin_amd64_v1/
}

build_darwin_arm64() {
    sleep 1

    # build is done by build_linux_amd64, wait for it to finish
    timeout=240
    while [[ -e buildlock ]] && [[ $timeout -gt 0 ]]; do
        sleep 1
        ((timeout--))
    done

    if [[ -e buildlock ]]; then
        echo "Timeout: Lock file not removed after 240 seconds."
        exit 1
    fi

    mkdir -p dist/gocastle_darwin_arm64/
    cp fyne-cross/bin/darwin-arm64/gocastle dist/gocastle_darwin_arm64/
}

if [[ "$4" == *"/gocastle_linux_amd64_v1/gocastle" ]]; then
    build_linux_amd64
elif [[ "$4" == *"/gocastle_linux_arm64/gocastle" ]]; then
    build_linux_arm64
elif [[ "$4" == *"/gocastle_android_arm64/gocastle" ]]; then
    build_android
elif [[ "$4" == *"/gocastle_windows_amd64_v1/gocastle.exe" ]]; then
    build_windows
elif [[ "$4" == *"/gocastle_darwin_amd64_v1/gocastle" ]]; then
    build_darwin_amd64
elif [[ "$4" == *"/gocastle_darwin_arm64/gocastle" ]]; then
    build_darwin_arm64
else
    echo "Invalid or unsupported path argument '$4'"
    exit 1
fi
