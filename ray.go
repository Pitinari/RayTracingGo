package main

type Ray struct {
	origin           Point3
	direction        Vect3
	remainingBounces int
	incomingLight    Color
	color            Color
}

type HitRecord struct {
	p         Point3
	normal    Vect3
	t         float64
	frontFace bool
	material  Material
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
		var scatteredRay Ray
		if hit.material.scatter(&ray, hit, &scatteredRay) {
			return scatteredRay.ray_color(world)
		} else {
			return ray.incomingLight
		}
	}
	return ray.incomingLight
}
