package main

func main() {

	// constants
	aspectRatio := 16.0 / 9.0
	width := 400
	height := int(float64(width) / aspectRatio)

	// Camera constants
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0
	origin := point_init(0, 0, 0)
	horizontal := vector_init(viewportWidth, 0, 0)
	vertical := vector_init(0, viewportHeight, 0)
	lowerLeftCorner := vector_sub(Vect3(origin), vector_add(horizontal.vector_scalar_div(2), vector_add(vertical.vector_scalar_div(2), vector_init(0, 0, focalLength))))

	// Render

	screen := screen_init(width, height)

	for i := width - 1; i >= 0; i-- {
		for j := height - 1; j >= 0; j-- {
			u := float64(i) / (float64(width) - 1.0)
			v := float64(j) / (float64(height) - 1.0)
			var ray Ray
			ray.origin = origin
			ray.direction = vector_add(lowerLeftCorner, vector_add(horizontal.vector_scalar_mul(u), vector_sub(vertical.vector_scalar_mul(v), Vect3(origin))))
			screen[i][j] = ray.ray_color()
		}
	}

	generate_image(width, height, screen, "image.png")
}
