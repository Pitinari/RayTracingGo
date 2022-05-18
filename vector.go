package main

import (
	"math"
)

type Vect3 struct {
	x float64
	y float64
	z float64
}

type point3 Vect3

func vector_opsite(vec *Vect3) {
	vec.x = -vec.x
	vec.y = -vec.y
	vec.z = -vec.z
}

func vector_add(vec1 Vect3, vec2 Vect3) Vect3 {
	var vecResult Vect3
	vecResult.x = vec1.x + vec2.x
	vecResult.y = vec1.y + vec2.y
	vecResult.z = vec1.z + vec2.z
	return vecResult
}

func vector_mul(vec1 Vect3, vec2 Vect3) Vect3 {
	var vecResult Vect3
	vecResult.x = vec1.x * vec2.x
	vecResult.y = vec1.y * vec2.y
	vecResult.z = vec1.z * vec2.z
	return vecResult
}

func vector_scalar_mul(vec Vect3, scalar float64) Vect3 {
	return vector_init(vec.x*scalar, vec.y*scalar, vec.z*scalar)
}

func vector_scalar_div(vec Vect3, scalar float64) Vect3 {
	return vector_scalar_mul(vec, 1/scalar)
}

func vector_length(vec Vect3) float64 {
	return math.Sqrt(vector_squared_length(vec))
}

func vector_squared_length(vec Vect3) float64 {
	return (vec.x * vec.x) + (vec.y * vec.y) + (vec.z * vec.z)
}

func vector_dot(vec1 Vect3, vec2 Vect3) float64 {
	vec := vector_mul(vec1, vec2)
	return vec.x + vec.y + vec.z
}

func vector_cross(vec1 Vect3, vec2 Vect3) Vect3 {
	return vector_init(
		vec1.y*vec2.z-vec1.z*vec2.y,
		vec1.z*vec2.x-vec1.x*vec2.z,
		vec1.x*vec2.y-vec1.y*vec2.x)
}

func vector_unit(vec Vect3) Vect3 {
	return vector_scalar_div(vec, vector_length(vec))
}

func vector_init(x float64, y float64, z float64) Vect3 {
	var vec Vect3
	vec.x = x
	vec.y = y
	vec.z = z
	return vec
}
