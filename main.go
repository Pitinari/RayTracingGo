package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numberOfObjects = 200

func generate_world() ArrayOfHittables {
	var world [numberOfObjects]Hittable
	for i := 2; i < 20; i++ {
		radius := rand.Float64()*3 - 2
		center := Point3{rand.Float64()*36 - 18, radius, rand.Float64()*36 - 18}

		var mat Material
		mat = MatteMaterial{1, Color{rand.Float64(), rand.Float64(), rand.Float64()}}
		switch rand.Int() % 3 {
		case 0:
			mat = MirroredlightMaterial{Color{rand.Float64(), rand.Float64(), rand.Float64()}}
		case 1:
			mat = FuzzyMaterial{Color{rand.Float64(), rand.Float64(), rand.Float64()}, 0.5}
		case 2:
			mat = DielectricMaterial{Color{rand.Float64(), rand.Float64(), rand.Float64()}, 1.5}
		}

		t := create_sphere(center, radius, mat)
		world[i] = t
	}
	for i := 20; i < 200; i++ {
		radius := rand.Float64()*2 - 0.5
		center := Point3{rand.Float64()*36 - 18, radius, rand.Float64()*36 - 18}

		var mat Material
		mat = MatteMaterial{1, Color{rand.Float64(), rand.Float64(), rand.Float64()}}
		switch rand.Int() % 4 {
		case 0:
			mat = MirroredlightMaterial{Color{rand.Float64(), rand.Float64(), rand.Float64()}}
		case 1:
			mat = MatteMaterial{1, Color{rand.Float64(), rand.Float64(), rand.Float64()}}
		case 2:
			mat = FuzzyMaterial{Color{rand.Float64(), rand.Float64(), rand.Float64()}, 0.5}
		case 3:
			mat = DielectricMaterial{Color{rand.Float64(), rand.Float64(), rand.Float64()}, 1.5}
		}

		t := create_sphere(center, radius, mat)
		world[i] = t
	}
	// for i := 0; i < numberOfObjects; i++ {
	// 	p1 := Point3{rand.Float64()*6 - 3, rand.Float64()*6 - 3, rand.Float64()*6 - 3}
	// 	p2 := Point3{rand.Float64()*6 - 3, rand.Float64()*6 - 3, rand.Float64()*6 - 3}
	// 	p3 := Point3{rand.Float64()*6 - 3, rand.Float64()*6 - 3, rand.Float64()*6 - 3}

	// 	var mat Material
	// 	mat = MatteMaterial{1, Color{rand.Float64(), rand.Float64(), rand.Float64()}}
	// 	switch rand.Int() % 3 {
	// 	case 0:
	// 		mat = MirroredlightMaterial{Color{rand.Float64(), rand.Float64(), rand.Float64()}}
	// 	case 1:
	// 		mat = MatteMaterial{1, Color{rand.Float64(), rand.Float64(), rand.Float64()}}
	// 	case 2:
	// 		mat = FuzzyMaterial{Color{rand.Float64(), rand.Float64(), rand.Float64()}, 0.5}
	// 	}
	// 	t := create_triangle(p1, p2, p3, mat)
	// 	world[i] = t
	// }
	return world[:]
}

func main() {
	rand.Seed(time.Now().Unix())
	aspectRatio := 16.0 / 9.0
	verticalFov := 20.0
	width := 1920
	// width = 400
	height := int(float64(width) / aspectRatio)
	lookFrom := point_init(0, 10, 50)
	lookAt := point_init(0, 0, 0)
	vup := vector_init(0, -1, 0)
	distFocus := 50.0
	aperture := 0.1
	cam := camera_init(verticalFov, aspectRatio, lookFrom, lookAt, vup, aperture, distFocus)

	// Render
	samplesPerPixel := 200
	maxBounces := 50
	screen := screen_init(width, height)
	world := generate_world()
	world[0] = create_triangle(point_init(0, 0, 50), point_init(50, 0, -50), point_init(-50, 0, -50), MatteMaterial{1, Color{rand.Float64(), rand.Float64(), rand.Float64()}})
	world[1] = create_sphere(point_init(0, 4, 0), 4, MirroredlightMaterial{Color{0.9, 0.9, 0.9}})

	cores := 12

	channel := make(chan Pixel, width*height)

	wg := &sync.WaitGroup{}
	wg.Add(cores)

	for i := 0; i < cores; i++ {
		go func() {
			fmt.Println("Thread init")
			for pixel := range channel {
				advance := len(channel)
				if advance%1000 == 0 {
					fmt.Println(advance, "items left")
				}
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
