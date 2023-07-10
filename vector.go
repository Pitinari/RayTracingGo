package main

import (
	"math"
	"math/rand"
)

const (
	X = 0
	Y = 1
	Z = 2
)

type Vect3 [3]float64

type Point3 Vect3

func (vec Vect3) x() float64 {
	return vec[X]
}

func (vec Vect3) y() float64 {
	return vec[Y]
}

func (vec Vect3) z() float64 {
	return vec[Z]
}

func (vec Vect3) vector_opsite() Vect3 {
	var vecOp Vect3
	vecOp[X] = -vec.x()
	vecOp[Y] = -vec.y()
	vecOp[Z] = -vec.z()
	return vecOp
}

func vector_add(vec1 Vect3, vec2 Vect3) Vect3 {
	var vecResult Vect3
	vecResult[X] = vec1.x() + vec2.x()
	vecResult[Y] = vec1.y() + vec2.y()
	vecResult[Z] = vec1.z() + vec2.z()
	return vecResult
}

func vector_sub(vec1 Vect3, vec2 Vect3) Vect3 {
	return vector_add(vec1, vec2.vector_opsite())
}

func vector_mul(vec1 Vect3, vec2 Vect3) Vect3 {
	var vecResult Vect3
	vecResult[X] = vec1.x() * vec2.x()
	vecResult[Y] = vec1.y() * vec2.y()
	vecResult[Z] = vec1.z() * vec2.z()
	return vecResult
}

func (vec Vect3) vector_scalar_mul(scalar float64) Vect3 {
	return vector_init(vec.x()*scalar, vec.y()*scalar, vec.z()*scalar)
}

func (vec Vect3) vector_scalar_div(scalar float64) Vect3 {
	return vec.vector_scalar_mul(1 / scalar)
}

func (vec Vect3) vector_length() float64 {
	return math.Sqrt(vec.vector_squared_length())
}

func (vec Vect3) vector_squared_length() float64 {
	return (vec.x() * vec.x()) + (vec.y() * vec.y()) + (vec.z() * vec.z())
}

func vector_random(min float64, max float64) Vect3 {
	delta := (max - min)
	return vector_init(rand.Float64()*delta+min, rand.Float64()*delta+min, rand.Float64()*delta+min)
}

func vector_dot(vec1 Vect3, vec2 Vect3) float64 {
	vec := vector_mul(vec1, vec2)
	return vec.x() + vec.y() + vec.z()
}

func vector_cross(vec1 Vect3, vec2 Vect3) Vect3 {
	return vector_init(
		vec1.y()*vec2.z()-vec1.z()*vec2.y(),
		vec1.z()*vec2.x()-vec1.x()*vec2.z(),
		vec1.x()*vec2.y()-vec1.y()*vec2.x())
}

func vector_unit(vec Vect3) Vect3 {
	return vec.vector_scalar_div(vec.vector_length())
}

func vector_init(x float64, y float64, z float64) Vect3 {
	return Vect3{x, y, z}
}

func point_init(x float64, y float64, z float64) Point3 {
	return Point3(vector_init(x, y, z))
}
