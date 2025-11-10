package main

import (
	//"image/color"
	//"math/rand"
	"fmt"
	"image/color"

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
		// make a blue square 🎨
		fmt.Println("Pressed:", e.Name)
		square := canvas.NewRectangle(color.NRGBA{R: 100, G: 150, B: 255, A: 255})
		square.Resize(fyne.NewSize(50, 50)) // width x height
		square.Move(fyne.NewPos(100, 100))  // position (x, y)

		content.Add(square)     // add to container
		canvas.Refresh(content) // redraw so it appears
	})

	w.ShowAndRun()
}
