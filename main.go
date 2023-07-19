package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numberOfObjects = 200

func generate_world() ArrayOfHittables {
	var world [numberOfObjects]Hittable
	for i := 2; i < 20; i++ {
		radius := rand.Float64()*3 - 2
		center := Point3{rand.Float64()*100 - 50, radius, rand.Float64()*100 - 50}

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
			mat = FuzzyMaterial{Color{rand.Float64(), rand.Float64(), rand.Float64()}, rand.Float64()}
		case 3:
			mat = DielectricMaterial{Color{rand.Float64(), rand.Float64(), rand.Float64()}, rand.Float64() + 0.5}
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
	width = 3840
	height := int(float64(width) / aspectRatio)
	lookFrom := point_init(0, 10, 50)
	lookAt := point_init(0, 0, 0)
	vup := vector_init(0, -1, 0)
	distFocus := 50.0
	aperture := 0.1
	cam := camera_init(verticalFov, aspectRatio, lookFrom, lookAt, vup, aperture, distFocus)

	// Render
	samplesPerPixel := 500
	maxBounces := 50
	world := generate_world()
	world[0] = create_triangle(point_init(0, 0, 4000), point_init(4000, 0, -4000), point_init(-4000, 0, -4000), MatteMaterial{1, Color{0.7, 0.7, 0.8}})
	world[1] = create_sphere(point_init(0, 4, 0), 4, MirroredlightMaterial{Color{1, 1, 1}})

	cores := 12

	channelSamples := make(chan int, samplesPerPixel)
	channelResults := make(chan [][]Color, samplesPerPixel)

	for i := 0; i < cores; i++ {
		go func(id int) {
			fmt.Println("Thread init", id)
			for sample := range channelSamples {
				screen := screen_init(width, height)
				fmt.Println("Processing sample: ", sample)
				for y := 0; y < height; y++ {
					for x := 0; x < width; x++ {
						u := (float64(x) + rand.Float64()) / (float64(width) - 1.0)
						v := (float64(y) + rand.Float64()) / (float64(height) - 1.0)
						ray := cam.get_ray(u, v, maxBounces)
						screen[x][y] = ray.ray_color(world)
					}
				}
				channelResults <- screen
				fmt.Println("Processed sample: ", sample)
			}
		}(i)
	}

	for i := 1; i <= samplesPerPixel; i++ {
		channelSamples <- i
	}
	close(channelSamples)
	collectedScreen := screen_init(width, height)
	samplesProcessed := 0

	for processedScreen := range channelResults {
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				currentR := collectedScreen[x][y][R]
				currentG := collectedScreen[x][y][G]
				currentB := collectedScreen[x][y][B]

				incomingR := processedScreen[x][y][R]
				incomingG := processedScreen[x][y][G]
				incomingB := processedScreen[x][y][B]

				collectedScreen[x][y][R] = currentR - (currentR / float64(samplesProcessed+1)) + (incomingR / float64(samplesProcessed+1))
				collectedScreen[x][y][G] = currentG - (currentG / float64(samplesProcessed+1)) + (incomingG / float64(samplesProcessed+1))
				collectedScreen[x][y][B] = currentB - (currentB / float64(samplesProcessed+1)) + (incomingB / float64(samplesProcessed+1))
			}
		}
		generate_image(width, height, collectedScreen, "image.png")
		samplesProcessed++
		if samplesProcessed == samplesPerPixel {
			close(channelResults)
		}
	}
}
