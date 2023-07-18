package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	aspectRatio := 16.0 / 9.0
	verticalFov := 90.0
	width := 1920
	// width = 400
	height := int(float64(width) / aspectRatio)
	cam := camera_init(verticalFov, aspectRatio, point_init(-1, 1, 0), point_init(0, 0, 1), vector_init(0, -1, 0))

	// Render
	samplesPerPixel := 10
	maxBounces := 10
	screen := screen_init(width, height)
	world := ArrayOfHittables{
		create_sphere(point_init(0, 0, 2), 0.5, MatteMaterial{0.5, color_init(1, 0, 0)}),
		create_sphere(point_init(-1, 0, 2), 0.5, FuzzyMaterial{color_init(0, 0.4, 0.8), 0.9}),
		create_sphere(point_init(1, 0, 2), 0.5, MirroredlightMaterial{color_init(0.3, 0.8, 0.5)}),
		create_sphere(point_init(1, 1, 3), 0.5, MirroredlightMaterial{color_init(0.8, 0.4, 0.2)}),
		create_sphere(point_init(0, -100, -2), 99.5, MatteMaterial{0.7, color_init(0.9, 0.3, 0.2)}),
		create_sphere(point_init(0, -0.3, 1.2), 0.2, DielectricMaterial{color_init(1, 1, 1), 0.7}),
		// create_sphere(point_init(2, -16, -2), 15, DiffuselightMaterial{color_init(1, 1, 1)}),
		// create_sphere(point_init(-5, -15, 2), 12, DiffuselightMaterial{color_init(1, 1, 1)}),
	}

	cores := 12

	channel := make(chan Pixel, width*height)

	wg := &sync.WaitGroup{}
	wg.Add(cores)

	for i := 0; i < cores; i++ {
		go func() {
			fmt.Println("Thread init")
			for pixel := range channel {
				col := color_init(0, 0, 0)
				for s := 0; s < samplesPerPixel; s++ {
					u := (float64(pixel.x) + rand.Float64()) / (float64(width) - 1.0)
					v := (float64(pixel.y) + rand.Float64()) / (float64(height) - 1.0)
					ray := cam.get_ray(u, v, maxBounces)
					col = col.add(ray.ray_color(world))
				}
				screen[pixel.x][pixel.y] = write_color(col, samplesPerPixel)
			}
			wg.Done()
		}()
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			channel <- Pixel{i, j}
		}
	}
	close(channel)
	fmt.Println(len(channel))
	wg.Wait()

	generate_image(width, height, screen, "image.png")
}
