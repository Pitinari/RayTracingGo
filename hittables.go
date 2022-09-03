package main

import (
	"math"
)

type Hittable interface {
	hit(Ray, float64, float64, *HitRecord) bool
}

type Sphere struct {
	center Point3
	radius float64
}

func create_sphere(center Point3, radius float64) Sphere {
	return Sphere{center, radius}
}

func (sp Sphere) hit(ray Ray, t_min float64, t_max float64, hit *HitRecord) bool {
	oc := vector_sub(Vect3(ray.origin), Vect3(sp.center))
	a := ray.direction.vector_squared_length()
	half_b := vector_dot(oc, ray.direction)
	c := oc.vector_squared_length() - sp.radius*sp.radius
	discriminant := half_b*half_b - a*c
	if discriminant < 0 || a == 0 {
		return false
	} else {
		sqrtd := math.Sqrt(discriminant)
		root := (-half_b - sqrtd) / a
		if root < t_min || t_max < root {
			root = (-half_b + sqrtd) / a
			if root < t_min || t_max < root {
				return false
			}
		}
		hit.t = root
		hit.p = ray.ray_at(root)
		hit.normal = vector_sub(Vect3(hit.p), Vect3(sp.center)).vector_opsite().vector_scalar_div(sp.radius)
		hit.set_face_normal(ray)
		return true
	}
}
