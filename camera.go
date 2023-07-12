package main

type Camera struct {
	aspectRatio, viewportHeight, viewportWidth, focalLength float64
	origin                                                  Point3
	horizontal, vertical, lowerLeftCorner                   Vect3
}

func camera_init() Camera {
	var cam Camera
	cam.aspectRatio = 16.0 / 9.0
	cam.viewportHeight = 2.0
	cam.viewportWidth = cam.aspectRatio * cam.viewportHeight
	cam.focalLength = 1.0
	cam.origin = point_init(0, 0, 0)
	cam.horizontal = vector_init(cam.viewportWidth, 0, 0)
	cam.vertical = vector_init(0, cam.viewportHeight, 0)
	cam.lowerLeftCorner = vector_sub(Vect3(cam.origin), vector_add(cam.horizontal.vector_scalar_div(2), vector_add(cam.vertical.vector_scalar_div(2), vector_init(0, 0, cam.focalLength))))
	return cam
}

func (cam Camera) get_ray(u float64, v float64, maxBounces int) Ray {
	return Ray{cam.origin, vector_add(cam.lowerLeftCorner, vector_add(cam.horizontal.vector_scalar_mul(u), vector_sub(cam.vertical.vector_scalar_mul(v), Vect3(cam.origin)))), maxBounces, color_init(0, 0, 0), color_init(1, 1, 1)}
}
