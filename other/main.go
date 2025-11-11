package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Visualizer")
	content := container.NewWithoutLayout()
	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))

	// Input
	w.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {

		fmt.Println("Pressed:", e.Name)
		square := canvas.NewRectangle(color.NRGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255})
		square.Resize(fyne.NewSize(50, 50))                                              // width x height
		square.Move(fyne.NewPos(float32(rand.Intn(800-40)), float32(rand.Intn(600-40)))) // position (x, y)

		content.Add(square)     // add to container
		canvas.Refresh(content) // redraw so it appears
	})

	w.ShowAndRun()
}
