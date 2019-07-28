package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 50

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

type ABC struct {
	X int
	Y int
	Z complex128
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)
	chs := make(chan ABC, 10)
	var wg sync.WaitGroup

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for ch := range chs {
				img.Set(ch.X, ch.Y, mandelbrot(ch.Z))
			}
			wg.Done()
		}()
	}

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			chs <- ABC{
				X: px,
				Y: py,
				Z: z,
			}
		}
	}
	close(chs)

	wg.Wait()
	png.Encode(os.Stdout, img)
}
