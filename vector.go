package main

import (
	"math"
)

type vect3 struct {
	x float64
	y float64
	z float64
}

func vector_opsite(vec *vect3) {
	vec.x = -vec.x
	vec.y = -vec.y
	vec.z = -vec.z
}

func vector_add(vec1 vect3, vec2 vect3) vect3 {
	var vecResult vect3
	vecResult.x = vec1.x + vec2.x
	vecResult.y = vec1.y + vec2.y
	vecResult.z = vec1.z + vec2.z
	return vecResult
}

func vector_mul(vec1 vect3, vec2 vect3) vect3 {
	var vecResult vect3
	vecResult.x = vec1.x * vec2.x
	vecResult.y = vec1.y * vec2.y
	vecResult.z = vec1.z * vec2.z
	return vecResult
}

func vector_scalar_mul(vec vect3, scalar float64) vect3 {
	return vector_init(vec.x*scalar, vec.y*scalar, vec.z*scalar)
}

func vector_scalar_div(vec vect3, scalar float64) vect3 {
	return vector_scalar_mul(vec, 1/scalar)
}

func vector_length(vec vect3) float64 {
	return math.Sqrt(vector_squared_length(vec))
}

func vector_squared_length(vec vect3) float64 {
	return (vec.x * vec.x) + (vec.y * vec.y) + (vec.z * vec.z)
}

func vector_dot(vec1 vect3, vec2 vect3) float64 {
	vec := vector_mul(vec1, vec2)
	return vec.x + vec.y + vec.z
}

func vector_cross(vec1 vect3, vec2 vect3) vect3 {
	return vector_init(
		vec1.y*vec2.z-vec1.z*vec2.y,
		vec1.z*vec2.x-vec1.x*vec2.z,
		vec1.x*vec2.y-vec1.y*vec2.x)
}

func vector_unit(vec vect3) vect3 {
	return vector_scalar_div(vec, vector_length(vec))
}

func vector_init(x float64, y float64, z float64) vect3 {
	var vec vect3
	vec.x = x
	vec.y = y
	vec.z = z
	return vec
}
