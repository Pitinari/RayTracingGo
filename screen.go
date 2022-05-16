package main

import(
	"image"
	"image/color"
	"image/png"
	"os"
)

type Color struct {
	r float64
	g float64
	b float64
}

func screen_init (width int, height int) [][]Color {
	a := make([][]Color, width)
	for i := range a {
	    a[i] = make([]Color, height)
	}
	return a
}

func generate_image(width int, height int, screen [][]Color, filename string){
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for i := 0; i < width; i++{
		for j := 0; j < height; j++{
			pixelR := uint8(255.999 * screen[i][j].r) 
			pixelG := uint8(255.999 * screen[i][j].g) 
			pixelB := uint8(255.999 * screen[i][j].b) 
			pixel := color.RGBA{pixelR, pixelG, pixelB, 255}
			img.Set(i, j, pixel)
		}
	}
	f, _ := os.Create(filename)
	png.Encode(f, img)
}