package main

func main() {

	// constants
	aspectRatio := 16.0 / 9.0
	width := 256
	height := int(float64(width) / aspectRatio)

	// Camera constants
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0
	origin := piont_init(0, 0, 0)
	horizontal := vector_init(viewportWidth, 0, 0)
	vertical := vector_init(0, viewportHeight, 0)
	lowerLeftCorner := vector_sub(Vect3(origin), vector_sub(vector_scalar_div(horizontal, 2), vector_sub(vector_scalar_div(vertical, 2), vector_init(0, 0, focalLength))))

	// Render

	screen := screen_init(width, height)

	for i := width - 1; i >= 0; i-- {
		for j := height - 1; j >= 0; j-- {
			u := float64(i) / (float64(width) - 1.0)
			v := float64(j) / (float64(height) - 1.0)
			println(u, v)
			var ray Ray
			ray.origin = origin
			ray.direction = vector_add(lowerLeftCorner, vector_add(vector_scalar_mul(horizontal, u), vector_add(vector_scalar_mul(vertical, v), vector_opsite(Vect3(origin)))))
			screen[i][j] = ray_color(ray)
		}
	}

	generate_image(width, height, screen, "image.png")
}
