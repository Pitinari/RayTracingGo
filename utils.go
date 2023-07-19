package main

import (
	"math"
	"math/rand"
)

func clamp(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func smoothstep(x float64, min float64, max float64) float64 {
	x = clamp((x-min)/(max-min), min, max)
	return x * x * (3.0 - 2.0*x)
}

func degrees_to_radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func random_in_unit_disk() Vect3 {
	for {
		p := Vect3{rand.Float64()*2 - 1, rand.Float64()*2 - 1, 0}
		if p.vector_squared_length() < 1 {
			return p
		}
	}
}
