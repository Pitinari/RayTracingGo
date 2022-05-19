package main

type Ray struct {
	origin    Point3
	direction Vect3
}

func ray_at(ray Ray, t float64) Point3 {
	return Point3(vector_add(Vect3(ray.origin), vector_scalar_mul(ray.direction, t)))
}

func ray_color(ray Ray) Color {
	unitDirection := vector_unit(ray.direction)
	t := 0.5*unitDirection[Y] + 1.0
	return color_add(color_scalar_mul(color_init(1.0, 1.0, 1.0), (1.0-t)), color_scalar_mul(color_init(0.5, 0.7, 1.0), t))
}
