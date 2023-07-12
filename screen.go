package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const (
	R = X
	G = Y
	B = Z
)

type Pixel struct {
	x, y int
}

func screen_init(width int, height int) [][]Color {
	a := make([][]Color, width)
	for i := range a {
		a[i] = make([]Color, height)
	}
	return a
}

func write_color(pixel Color, samplesPerPixel int) Color {
	r := pixel[R]
	g := pixel[G]
	b := pixel[B]

	// Divide the color by the number of samples.
	scale := 1.0 / float64(samplesPerPixel)
	r = math.Sqrt(r * scale)
	g = math.Sqrt(g * scale)
	b = math.Sqrt(b * scale)
	return color_init(clamp(r, 0.0, 0.999), clamp(g, 0.0, 0.999), clamp(b, 0.0, 0.999))
}

func generate_image(width int, height int, screen [][]Color, filename string) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			pixelR := uint8(255 * screen[i][j][R])
			pixelG := uint8(255 * screen[i][j][G])
			pixelB := uint8(255 * screen[i][j][B])
			pixel := color.RGBA{pixelR, pixelG, pixelB, 255}
			img.Set(i, j, pixel)
		}
	}
	f, _ := os.Create(filename)
	png.Encode(f, img)
}
