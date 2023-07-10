package main

import "math"

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
			hitAnything = true
			if tempRec.t < closestSoFar {
				closestSoFar = tempRec.t
				*hit = tempRec
			}
		}
	}

	return hitAnything
}

type Material struct {
	color Color
}

func create_material(color Color) Material {
	return Material{color}
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
		hit.normal = vector_sub(Vect3(hit.p), Vect3(sp.center)).vector_opsite().vector_scalar_div(sp.radius)
		// hit.set_face_normal(ray)
		hit.material = sp.material
		return true
	}
}
