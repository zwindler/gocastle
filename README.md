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

For the development, you need to install [Golangci-lint](https://golangci-lint.run/usage/install/#local-installation)

## Contributions

Contributions are welcome **but** I need you to follow a few rules:
* the primary goal of this project is for me to learn Go while having fun (game dev is so much fun). Keep that in mind when asking / proposing modifications.
* since I'm a beginner, my code quality can be improved and I've a lot to learn. Fixing my code where it's badly written or helping me learn better patterns is greatly appreciated, but please add a little context/explaination so I can grow from your contributions. See [developpement_diary/](https://github.com/zwindler/gocastle/tree/main/developpement_diary) folder.
* most of the future work for the game is written in the [Issues](https://github.com/zwindler/gocastle/issues) of this project. You look at them and ask me if you can take those!
  * issues labelled "game design" are by default reserved to me (@zwindler) since this needs not only code but also thinking about how the game will work.
  * issues labelled "good first issues" are self explainatory ;-\).
  * issues labelled "on hold" as well.
  * issues labelled "refactoring" need to be discussed first. I openned a Discord server for this, please ask me the details.

Most importantly, have fun!

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

Note: for darwin compilation, we need fyne-cross which requires having locally the XCode SDK, especially `Command_Line_Tools_for_Xcode_12.4.dmg`. [Download it here](https://developer.apple.com/download/all/?q=Command%20Line%20Tools). See [fyne-cross documentation for more information](https://github.com/fyne-io/fyne-cross/blob/master/README.md).
