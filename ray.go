package main

type Ray struct {
	origin    Point3
	direction Vect3
}

type HitRecord struct {
	p         Point3
	normal    Vect3
	t         float64
	frontFace bool
}

func (hit HitRecord) set_face_normal(r Ray) {
	hit.frontFace = (vector_dot(r.direction, hit.normal) < 0)
	if !hit.frontFace {
		hit.normal = hit.normal.vector_opsite()
	}
}

func (ray Ray) ray_at(t float64) Point3 {
	return Point3(vector_add(Vect3(ray.origin), ray.direction.vector_scalar_mul(t)))
}

func (ray Ray) ray_color(world Hittables) Color {
	var hit HitRecord
	if world.hit_list(ray, 0, 5, &hit) {
		return color_init(hit.normal.x()+1.0, hit.normal.y()+1.0, hit.normal.z()+1.0).color_scalar_mul(0.5)
	}
	unitDirection := vector_unit(ray.direction)
	s := 0.5 * (unitDirection.y() + 1.0)
	return color_add(color_init(1.0, 1.0, 1.0).color_scalar_mul(1.0-s), color_init(0.5, 0.7, 1.0).color_scalar_mul(s))
}
