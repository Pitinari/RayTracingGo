package main

import (
	"math"
)

const (
	X = 0
	Y = 1
	Z = 2
)

type Vect3 [3]float64

type Point3 Vect3

func vector_opsite(vec Vect3) Vect3 {
	var vecOp Vect3
	vecOp[X] = -vec[X]
	vecOp[Y] = -vec[Y]
	vecOp[Z] = -vec[Z]
	return vecOp
}

func vector_add(vec1 Vect3, vec2 Vect3) Vect3 {
	var vecResult Vect3
	vecResult[X] = vec1[X] + vec2[X]
	vecResult[Y] = vec1[Y] + vec2[Y]
	vecResult[Z] = vec1[Z] + vec2[Z]
	return vecResult
}

func vector_sub(vec1 Vect3, vec2 Vect3) Vect3 {
	return vector_add(vec1, vector_opsite(vec2))
}

func vector_mul(vec1 Vect3, vec2 Vect3) Vect3 {
	var vecResult Vect3
	vecResult[X] = vec1[X] * vec2[X]
	vecResult[Y] = vec1[Y] * vec2[Y]
	vecResult[Z] = vec1[Z] * vec2[Z]
	return vecResult
}

func vector_scalar_mul(vec Vect3, scalar float64) Vect3 {
	return vector_init(vec[X]*scalar, vec[Y]*scalar, vec[Z]*scalar)
}

func vector_scalar_div(vec Vect3, scalar float64) Vect3 {
	return vector_scalar_mul(vec, 1/scalar)
}

func vector_length(vec Vect3) float64 {
	return math.Sqrt(vector_squared_length(vec))
}

func vector_squared_length(vec Vect3) float64 {
	return (vec[X] * vec[X]) + (vec[Y] * vec[Y]) + (vec[Z] * vec[Z])
}

func vector_dot(vec1 Vect3, vec2 Vect3) float64 {
	vec := vector_mul(vec1, vec2)
	return vec[X] + vec[Y] + vec[Z]
}

func vector_cross(vec1 Vect3, vec2 Vect3) Vect3 {
	return vector_init(
		vec1[Y]*vec2[Z]-vec1[Z]*vec2[Y],
		vec1[Z]*vec2[X]-vec1[X]*vec2[Z],
		vec1[X]*vec2[Y]-vec1[Y]*vec2[X])
}

func vector_unit(vec Vect3) Vect3 {
	return vector_scalar_div(vec, vector_length(vec))
}

func vector_init(x float64, y float64, z float64) Vect3 {
	var vec Vect3
	vec[X] = x
	vec[Y] = y
	vec[Z] = z
	return vec
}

func piont_init(x float64, y float64, z float64) Point3 {
	return Point3(vector_init(x, y, z))
}
