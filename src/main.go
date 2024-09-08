package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Grand Casino Capital Launcher")

	w.CenterOnScreen()
	w.Resize(fyne.NewSize(400, 300))

	downloadNewGameButton := widget.NewButton("Download Recent Game", func() { DownloadRecentGame() })
	launchButton := widget.NewButton("Launch Grand Casino Capital", func() { launchGame() })

	c := container.NewVBox(launchButton, downloadNewGameButton)

	w.SetContent(c)
	w.ShowAndRun()
}

func launchGame() {
	var stderr bytes.Buffer
	var out bytes.Buffer

	fullPath, _ := filepath.Abs("../src/game")

	cmd := exec.Command("cmd.exe", "/C", "Start.BAT")
	cmd.Dir = fullPath
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	fmt.Print(out.String())
	fmt.Print(stderr.String())

	if err != nil {
		fmt.Print(stderr.String())
		fmt.Print(err.Error())
	}
}
