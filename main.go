package main

import (
	"time"
	"log"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w:= a.NewWindow("Clock")

	w.Resize(fyne.NewSize(200, 200))
	
	clock := widget.NewLabel("")
	startButton := widget.NewButton("Start", func() {
			log.Println("start tapped")
			
	})
	stopButton := widget.NewButton("Stop", func() {
		log.Println("stop tapped")
	})
		
	clockCenter := container.New(layout.NewCenterLayout(), clock)	
	clockBox := container.New(layout.NewVBoxLayout(), clockCenter)
	buttonBox := container.New(layout.NewHBoxLayout(), startButton, stopButton)
	buttonBoxCenter := container.New(layout.NewCenterLayout(), buttonBox)

	content := container.New(layout.NewVBoxLayout(), clockBox, buttonBoxCenter)
	contentCenter := container.New(layout.NewCenterLayout(), content)
	updateTime(clock)

	w.SetContent(contentCenter)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
		}()
		
	w.ShowAndRun()
}
