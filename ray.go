package main

type Ray struct {
	origin    point3
	direction Vect3
}

func ray_origin(ray Ray) point3 {
	return ray.origin
}

func ray_direction(ray Ray) Vect3 {
	return ray.direction
}

func ray_at(ray Ray, t float64) point3 {
	return point3(vector_add(Vect3(ray.origin), vector_scalar_mul(ray.direction, t)))
}
