package main

type Ray struct {
	origin           Point3
	direction        Vect3
	remainingBounces int
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
	if ray.remainingBounces <= 0 {
		return color_init(0, 0, 0)
	}
	var hit HitRecord
	if world.hit_list(ray, 0.001, 100, &hit) {
		randomVect := vector_unit(vector_random(-1, 1))
		if vector_dot(randomVect, hit.normal) > 0 {
			randomVect = randomVect.vector_opsite()
		}
		bouncedRay := Ray{hit.p, vector_sub(Vect3(hit.p), vector_add(hit.normal, randomVect)), ray.remainingBounces - 1}
		return bouncedRay.ray_color(world).color_scalar_mul(0.5)
	}
	unitDirection := vector_unit(ray.direction)
	s := 0.5 * (unitDirection.y() + 1.0)
	return (color_init(1.0, 1.0, 1.0).color_scalar_mul(1.0 - s)).add(color_init(0.5, 0.7, 1.0).color_scalar_mul(s))
}
