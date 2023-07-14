package main

type Material interface {
	scatter(*Ray, HitRecord, *Ray) bool
}

type LambertianMaterial struct {
	reflectance Color
}

func (lm LambertianMaterial) scatter(rayIn *Ray, hit HitRecord, rayScattered *Ray) bool {
	randomVect := vector_unit(vector_random(-1, 1))
	if vector_dot(randomVect, hit.normal) > 0 {
		randomVect = randomVect.vector_opsite()
	}
	*rayScattered = Ray{
		hit.p,
		vector_sub(Vect3(hit.p),
			vector_add(hit.normal, randomVect)),
		rayIn.remainingBounces - 1,
		rayIn.incomingLight,
		rayIn.color.mult(lm.reflectance),
	}
	return true
}

type DiffuselightMaterial struct {
	c Color
}

func (dm DiffuselightMaterial) scatter(rayIn *Ray, hit HitRecord, rayScattered *Ray) bool {
	(*rayIn).incomingLight = (*rayIn).incomingLight.add((dm.c.mult(rayIn.color)))
	return false
}

type MirroredlightMaterial struct {
	reflectance Color
}

func (mm MirroredlightMaterial) scatter(rayIn *Ray, hit HitRecord, rayScattered *Ray) bool {
	v := vector_unit(rayIn.direction)
	*rayScattered = Ray{
		hit.p,
		vector_sub(v, hit.normal.vector_scalar_mul(2*vector_dot(v, hit.normal))),
		rayIn.remainingBounces - 1,
		rayIn.incomingLight,
		rayIn.color.mult(mm.reflectance),
	}
	return true
}

type FuzzyMaterial struct {
	reflectance Color
	fuzz        float64
}

func (fm FuzzyMaterial) scatter(rayIn *Ray, hit HitRecord, rayScattered *Ray) bool {
	fuzz := fm.fuzz
	if fuzz > 1 {
		fuzz = 1
	}
	randomVect := vector_unit(vector_random(-1, 1))
	if vector_dot(randomVect, hit.normal) > 0 {
		randomVect = randomVect.vector_opsite()
	}
	v := vector_add(vector_unit(rayIn.direction), randomVect.vector_scalar_mul(fuzz))
	*rayScattered = Ray{
		hit.p,
		vector_sub(v, hit.normal.vector_scalar_mul(2*vector_dot(v, hit.normal))),
		rayIn.remainingBounces - 1,
		rayIn.incomingLight,
		rayIn.color.mult(fm.reflectance),
	}
	return true
}
