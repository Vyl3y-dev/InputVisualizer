package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type particle struct {
	rect     *canvas.Rectangle
	baseRGBA color.NRGBA
	alpha    uint8
	x, y     float32
	size     float32
	angle    float64
	speedX   float32
	speedY   float32
	fadeRate uint8
}

func main() {
	a := app.New()
	w := a.NewWindow("Floaty Squares ✨")

	content := container.NewWithoutLayout()
	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))

	rand.Seed(time.Now().UnixNano())

	w.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		// Random base color
		base := color.NRGBA{
			R: uint8(rand.Intn(255)),
			G: uint8(rand.Intn(255)),
			B: uint8(rand.Intn(255)),
			A: 255,
		}

		// Random size, position, drift speed, and fade rate
		size := float32(rand.Intn(40) + 20) // 20–60 px
		x := float32(rand.Intn(800 - int(size)))
		y := float32(rand.Intn(600 - int(size)))
		speedX := (rand.Float32() - 0.5) * 2 // small left/right drift
		speedY := rand.Float32()*1.5 + 0.5   // gentle upward float
		fadeRate := uint8(rand.Intn(4) + 3)  // different fade speeds

		sq := canvas.NewRectangle(base)
		sq.Resize(fyne.NewSize(size, size))
		sq.Move(fyne.NewPos(x, y))

		content.Add(sq)
		canvas.Refresh(content)

		p := &particle{
			rect:     sq,
			baseRGBA: base,
			alpha:    255,
			x:        x,
			y:        y,
			size:     size,
			angle:    rand.Float64() * math.Pi * 2,
			speedX:   speedX,
			speedY:   speedY,
			fadeRate: fadeRate,
		}

		// animation loop (float + fade + drift)
		go func() {
			ticker := time.NewTicker(16 * time.Millisecond)
			defer ticker.Stop()

			for range ticker.C {
				p.y -= p.speedY // float upward
				p.x += float32(math.Sin(p.angle)) * p.speedX
				p.angle += 0.05 // create a gentle wave motion
				if p.alpha > p.fadeRate {
					p.alpha -= p.fadeRate
				} else {
					p.alpha = 0
				}

				// rebuild the color with new alpha
				c := color.NRGBA{
					R: p.baseRGBA.R,
					G: p.baseRGBA.G,
					B: p.baseRGBA.B,
					A: p.alpha,
				}

				fyne.Do(func() {
					p.rect.Move(fyne.NewPos(p.x, p.y))
					p.rect.FillColor = c
					p.rect.Refresh()
				})

				if p.alpha == 0 || p.y < -p.size {
					fyne.Do(func() {
						content.Remove(p.rect)
						canvas.Refresh(content)
					})
					return
				}
			}
		}()
	})

	w.ShowAndRun()
}
