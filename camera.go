package main

import "math"

type Camera struct {
	aspectRatio, viewportHeight, viewportWidth, focalLength float64
	origin                                                  Point3
	horizontal, vertical, lowerLeftCorner                   Vect3
}

func camera_init(verticalFov float64, aspectRatio float64, lookFrom Point3, lookAt Point3, vup Vect3) Camera {
	theta := degrees_to_radians(verticalFov)
	h := math.Tan(theta / 2)

	var cam Camera
	cam.aspectRatio = aspectRatio
	cam.viewportHeight = 2.0 * h
	cam.viewportWidth = cam.aspectRatio * cam.viewportHeight
	cam.focalLength = 1.0
	cam.origin = lookFrom

	w := vector_unit(vector_sub(Vect3(lookFrom), Vect3(lookAt)))
	u := vector_unit(vector_cross(vup, w))
	v := vector_cross(w, u)

	cam.horizontal = u.vector_scalar_mul(cam.viewportWidth)
	cam.vertical = v.vector_scalar_mul(cam.viewportHeight)
	cam.lowerLeftCorner = vector_sub(
		Vect3(cam.origin),
		vector_add(
			cam.horizontal.vector_scalar_div(2),
			vector_add(
				cam.vertical.vector_scalar_div(2),
				w,
			),
		),
	)
	return cam
}

func (cam Camera) get_ray(u float64, v float64, maxBounces int) Ray {
	return Ray{
		cam.origin,
		vector_add(
			cam.lowerLeftCorner,
			vector_add(
				cam.horizontal.vector_scalar_mul(u),
				vector_sub(
					cam.vertical.vector_scalar_mul(v),
					Vect3(cam.origin),
				),
			),
		),
		maxBounces,
		color_init(0, 0, 0),
		color_init(1, 1, 1),
	}
}
