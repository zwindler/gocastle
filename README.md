# GoCastle

## Introduction

This project is my journey to creating a graphical game while learning Golang

Every session, I'll add an entry in this file telling what I did and what I learned (see [developpement_diary/](developpement_diary/))

## Prerequisites 

```bash
sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev
```

Also, for building releases, you need goreleaser and fyne-cross

```bash
go install github.com/fyne-io/fyne-cross@latest
```

## Build, test, release

Build only

```bash
make build
```

Build and run (useful in dev)

```bash
make buildrun
```

Release (using goreleaser)

```bash
git tag -a 0.0.1 -m "0.0.1 release"
goreleaser --clean
```

Test release (against code not tagged)

```bash
goreleaser --snapshot --clean
```
