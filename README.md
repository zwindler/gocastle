# GoCastle

## Introduction

This project is my journey to creating a graphical game while learning Golang

Every session, I'll add an entry in this file telling what I did and what I learned

## 2023-07-07

**Bootstraping the project**

I'm going to use [fyne](https://developer.fyne.io/started/) toolkit which is a simple tool to create graphical user interfaces in Golang. It seems to be simple enough, I hope it's not too small.

**Installing prerequisites**

```bash
sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev
```

Creating a directory and bootstraping golang project

```bash
mkdir gocastle
cd gocastle
```

Creating a new go module

```bash
go mod init gocastle
go: creating new go.mod: module gocastle
go: to add module requirements and sums:
        go mod tidy
zwindler@normandy:/windows/Users/zwindler/sources/gocastle$ 
```

Creating a main.go file

```golang
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	mainWindow := myApp.NewWindow("GoCastle")

	newGameButton := widget.NewButton("New Game", func() {
	})
	loadGameButton := widget.NewButton("Load Game", func() {
	})
	quitButton := widget.NewButton("Quit", func() {
		myApp.Quit()
	})

	menu := container.NewVBox(
		newGameButton,
		loadGameButton,
		quitButton,
	)

	mainWindow.SetContent(menu)
	mainWindow.ShowAndRun()
}
```

Init the project, download fyne, build and run it

```bash
go mod init gocastle
go mod tidy
go build
./gocastle
```

It will crash with this error

```
GLX: Failed to create context: BadValue (integer parameter out of range for operation)
```