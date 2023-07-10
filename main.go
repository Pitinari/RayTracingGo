package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	aspectRatio := 16.0 / 9.0
	width := 480
	height := int(float64(width) / aspectRatio)
	cam := camera_init()

	// Render
	samplesPerPixel := 30
	maxBounces := 20
	screen := screen_init(width, height)
	world := ArrayOfHittables{
		create_sphere(point_init(0, 0, -2), 0.5, create_material(color_init(256, 0, 0))),
		create_sphere(point_init(-1, -1, -3), 0.5, create_material(color_init(0, 256, 0))),
		create_sphere(point_init(0, 10, -2), 9.5, create_material(color_init(0, 0, 256))),
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

	for i := width - 1; i >= 0; i-- {
		for j := height - 1; j >= 0; j-- {
			channel <- Pixel{i, j}
		}
	}
	close(channel)
	fmt.Println(len(channel))
	wg.Wait()

	generate_image(width, height, screen, "image.png")
}
