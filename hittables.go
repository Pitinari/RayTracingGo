package main

import (
	"math"
)

type Hittable interface {
	hit(Ray, float64, float64, *HitRecord) bool
}

type Hittables interface {
	hit_list(Ray, float64, float64, *HitRecord) bool
}

type ArrayOfHittables []Hittable

func (world ArrayOfHittables) hit_list(ray Ray, tMin float64, tMax float64, hit *HitRecord) bool {
	var tempRec HitRecord
	hitAnything := false
	closestSoFar := tMax

	for i := 0; i < len(world); i++ {
		if world[i].hit(ray, tMin, closestSoFar, &tempRec) {
			if tempRec.t < closestSoFar {
				hitAnything = true
				closestSoFar = tempRec.t
				*hit = tempRec
			}
		}
	}

	return hitAnything
}

type Sphere struct {
	center   Point3
	radius   float64
	material Material
}

func create_sphere(center Point3, radius float64, material Material) Sphere {
	return Sphere{center, radius, material}
}

func (sp Sphere) hit(ray Ray, tMin float64, tMax float64, hit *HitRecord) bool {
	oc := vector_sub(Vect3(ray.origin), Vect3(sp.center))
	a := ray.direction.vector_squared_length()
	halfB := vector_dot(oc, ray.direction)
	c := oc.vector_squared_length() - sp.radius*sp.radius
	discriminant := halfB*halfB - a*c
	if discriminant < 0 || a == 0 {
		return false
	} else {
		sqrtd := math.Sqrt(discriminant)
		root := (-halfB - sqrtd) / a
		if root < tMin || tMax < root {
			root = (-halfB + sqrtd) / a
			if root < tMin || tMax < root {
				return false
			}
		}
		hit.t = root
		hit.p = ray.ray_at(root)
		hit.normal = vector_sub(Vect3(hit.p), Vect3(sp.center)).vector_scalar_div(sp.radius)
		hit.set_face_normal(ray)
		hit.material = sp.material
		return true
	}
}

type Triangle struct {
	vertices [3]Point3
	normal   Vect3
	material Material
}

func create_triangle(v1 Point3, v2 Point3, v3 Point3, material Material) Triangle {
	return Triangle{[3]Point3{v1, v2, v3}, vector_unit(vector_cross(vector_sub(Vect3(v2), Vect3(v1)), vector_sub(Vect3(v3), Vect3(v1)))), material}
}

func (tr Triangle) hit(ray Ray, tMin float64, tMax float64, hit *HitRecord) bool {
	const EPSILON = 0.0000001
	vertex0 := tr.vertices[0]
	vertex1 := tr.vertices[1]
	vertex2 := tr.vertices[2]
	var edge1, edge2, h, s, q Vect3
	var a, f, u, v float64
	edge1 = vector_sub(Vect3(vertex1), Vect3(vertex0))
	edge2 = vector_sub(Vect3(vertex2), Vect3(vertex0))
	h = vector_cross(ray.direction, edge2)
	a = vector_dot(h, edge1)

	if a > -EPSILON && a < EPSILON {
		return false // This ray is parallel to this triangle.
	}

	f = 1.0 / a
	s = vector_sub(Vect3(ray.origin), Vect3(vertex0))
	u = vector_dot(s, h) * f

	if u < 0.0 || u > 1.0 {
		return false
	}

	q = vector_cross(s, edge1)
	v = vector_dot(ray.direction, q) * f

	if v < 0.0 || u+v > 1.0 {
		return false
	}

	// At this stage we can compute t to find out where the intersection point is on the line.
	t := vector_dot(edge2, q) * f

	if t > EPSILON {
		hit.t = t
		hit.p = Point3(vector_add(Vect3(ray.origin), ray.direction.vector_scalar_mul(t)))
		hit.normal = tr.normal
		hit.set_face_normal(ray)
		hit.material = tr.material
		return true
	} else { // This means that there is a line intersection but not a ray intersection.
		return false
	}
}
