package main

import (
	"math"
)

type Camera struct {
	aspectRatio, viewportHeight, viewportWidth, lensRadius float64
	origin                                                 Point3
	horizontal, vertical, lowerLeftCorner                  Vect3
	u, v, w                                                Vect3
}

func camera_init(verticalFov float64, aspectRatio float64, lookFrom Point3, lookAt Point3, vup Vect3, aperture float64, focusDist float64) Camera {
	theta := degrees_to_radians(verticalFov)
	h := math.Tan(theta / 2.0)

	var cam Camera
	cam.aspectRatio = aspectRatio
	cam.viewportHeight = 2.0 * h
	cam.viewportWidth = cam.aspectRatio * cam.viewportHeight
	cam.lensRadius = aperture / 2.0
	cam.origin = lookFrom

	cam.w = vector_unit(vector_sub(Vect3(lookFrom), Vect3(lookAt)))
	cam.u = vector_unit(vector_cross(vup, cam.w))
	cam.v = vector_cross(cam.w, cam.u)

	cam.horizontal = cam.u.vector_scalar_mul(cam.viewportWidth * focusDist)
	cam.vertical = cam.v.vector_scalar_mul(cam.viewportHeight * focusDist)
	cam.lowerLeftCorner = vector_sub(
		Vect3(cam.origin),
		vector_add(
			cam.horizontal.vector_scalar_div(2),
			vector_add(
				cam.vertical.vector_scalar_div(2),
				cam.w.vector_scalar_mul(focusDist),
			),
		),
	)
	return cam
}

func (cam Camera) get_ray(s float64, t float64, maxBounces int) Ray {
	rd := random_in_unit_disk().vector_scalar_mul(cam.lensRadius)
	offset := vector_add(cam.u.vector_scalar_mul(rd.x()), cam.v.vector_scalar_mul(rd.y()))

	return Ray{
		Point3(vector_add(Vect3(cam.origin), offset)),
		vector_add(
			cam.lowerLeftCorner,
			vector_add(
				cam.horizontal.vector_scalar_mul(s),
				vector_sub(
					cam.vertical.vector_scalar_mul(t),
					vector_add(Vect3(cam.origin), offset),
				),
			),
		),
		maxBounces,
		color_init(0, 0, 0),
		color_init(1, 1, 1),
	}
}
