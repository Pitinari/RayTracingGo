package main

type Sphere struct {
	center Point3
	radius float64
}

func sphere_will_be_hitted(sp Sphere, ray Ray) bool {
	oc := vector_add(Vect3(ray.origin), Vect3(sp.center))
	a := vector_dot(ray.direction, ray.direction)
	b := 2.0 * vector_dot(oc, ray.direction)
	c := vector_dot(oc, oc) - sp.radius*sp.radius
	discriminant := b*b - 4*a*c
	return (discriminant > 0)
}
