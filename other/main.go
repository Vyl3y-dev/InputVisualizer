package main

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Vee's Visualizer")

	content := container.NewWithoutLayout()
	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))

	rand.Seed(time.Now().UnixNano())

	w.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		// random pretty circle
		circle := canvas.NewCircle(color.NRGBA{
			R: uint8(rand.Intn(255)),
			G: uint8(rand.Intn(255)),
			B: uint8(rand.Intn(255)),
			A: 255,
		})
		circle.Resize(fyne.NewSize(40, 40))
		circle.Move(fyne.NewPos(
			float32(rand.Intn(800-40)),
			float32(rand.Intn(600-40)),
		))

		content.Add(circle)
		canvas.Refresh(content)
	})

	w.ShowAndRun()
}
