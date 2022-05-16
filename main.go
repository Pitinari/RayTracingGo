package main

func main () {
	
	// constants
	width := 256
	height := 256

	// Render

	screen := screen_init(width, height)
	
	for i := width-1; i>=0; i--{
		for j := height-1; j>=0; j--{
			screen[i][j].r = float64(i) / float64(width-1)
			screen[i][j].g = float64(j) / float64(height-1)
			screen[i][j].b = 0.25
		}
	}

	generate_image(width, height, screen, "image.png")
}