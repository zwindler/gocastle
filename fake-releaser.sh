#!/bin/bash

GOBIN=~/go/bin/
PATH=$PATH:$GOBIN

# Function to extract tar archive and move files to the destination directory
extract_and_move() {
    local source_file="$1"
    local target_dir="$2"
    
    tar --strip-components=3 --wildcards -xJ '*/gocastle' -f "$source_file" -C "$target_dir"
}

build_linux_amd64() {
    touch lock
    fyne-cross linux -arch=amd64,arm64 -app-id fr.zwindler.gocastle
    extract_and_move fyne-cross/dist/linux-amd64/gocastle.tar.xz dist/gocastle_linux_amd64_v1/
    echo `ls -al dist/gocastle_linux_amd64_v1/` > toto.log
    sleep 1
    rm lock
}

build_linux_arm64() {
    sleep 1
    fyne-cross linux -arch=amd64,arm64 -app-id fr.zwindler.gocastle

    timeout=20
    while [[ -e lock ]] && [[ $timeout -gt 0 ]]; do
        sleep 1
        ((timeout--))
    done

    if [[ -e lock ]]; then
        echo "Timeout: Lock file not removed after 20 seconds."
        exit 1
    fi

    extract_and_move fyne-cross/dist/linux-arm64/gocastle.tar.xz dist/gocastle_linux_arm64/
}

build_android() {
    fyne-cross android -arch=arm64 -app-id fr.zwindler.gocastle
    cp fyne-cross/dist/android-arm64/gocastle.apk dist/gocastle_android_arm64/
}

build_windows() {
    fyne-cross windows -app-id fr.zwindler.gocastle
    unzip fyne-cross/dist/windows-amd64/gocastle.exe.zip -d dist/gocastle_windows_amd64_v1/
}

if [[ "$4" == *"/gocastle_linux_amd64_v1/gocastle" ]]; then
    build_linux_amd64
elif [[ "$4" == *"/gocastle_linux_arm64/gocastle" ]]; then
    build_linux_arm64
elif [[ "$4" == *"/gocastle_android_arm64/gocastle" ]]; then
    build_android
elif [[ "$4" == *"/gocastle_windows_amd64_v1/gocastle" ]]; then
    build_windows
else
    echo "Invalid or unsupported path argument."
    exit 1
fi
