package main

type Ray struct {
	origin    Point3
	direction Vect3
}

func (ray Ray) ray_at(t float64) Point3 {
	return Point3(vector_add(Vect3(ray.origin), ray.direction.vector_scalar_mul(t)))
}

func (ray Ray) ray_color() Color {
	var sp Sphere
	sp.center = point_init(0, 0, -1)
	sp.radius = 0.5
	if sphere_will_be_hitted(sp, ray) {
		return color_init(1, 0, 0)
	}
	unitDirection := vector_unit(ray.direction)
	t := 0.5 * (unitDirection[Y] + 1.0)
	return color_add(color_init(1.0, 1.0, 1.0).color_scalar_mul(1.0-t), color_init(0.5, 0.7, 1.0).color_scalar_mul(t))
}
