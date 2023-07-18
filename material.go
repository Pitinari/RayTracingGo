package main

import (
	"math"
	"math/rand"
)

type Material interface {
	scatter(*Ray, HitRecord, *Ray) bool
}

type LambertianMaterial struct {
	color Color
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
		rayIn.color.mult(lm.color),
	}
	return true
}

type DiffuselightMaterial struct {
	color Color
}

func (dm DiffuselightMaterial) scatter(rayIn *Ray, hit HitRecord, rayScattered *Ray) bool {
	(*rayIn).incomingLight = (*rayIn).incomingLight.add(dm.color.mult((*rayIn).color))
	return false
}

type MirroredlightMaterial struct {
	color Color
}

func reflect(direction Vect3, normal Vect3) Vect3 {
	return vector_sub(direction, normal.vector_scalar_mul(2*vector_dot(direction, normal)))
}

func (mm MirroredlightMaterial) scatter(rayIn *Ray, hit HitRecord, rayScattered *Ray) bool {
	v := vector_unit(rayIn.direction)
	*rayScattered = Ray{
		hit.p,
		reflect(v, hit.normal),
		rayIn.remainingBounces - 1,
		rayIn.incomingLight,
		rayIn.color.mult(mm.color),
	}
	return true
}

type FuzzyMaterial struct {
	color Color
	fuzz  float64
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
		rayIn.color.mult(fm.color),
	}
	return true
}

func refract(direction Vect3, normal Vect3, etaiOverEtat float64) Vect3 {
	cosTheta := vector_dot(direction.vector_opsite(), normal)
	rOutPerp := vector_add(direction, normal.vector_scalar_mul(cosTheta)).vector_scalar_mul(etaiOverEtat)
	rOutParallel := normal.vector_scalar_mul(-1 * math.Sqrt(math.Abs(1.0-rOutPerp.vector_squared_length())))
	return vector_add(rOutPerp, rOutParallel)
}

func reflectance_aproximation(cosine float64, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}

type DielectricMaterial struct {
	color           Color
	refractionIndex float64
}

func (dm DielectricMaterial) scatter(rayIn *Ray, hit HitRecord, rayScattered *Ray) bool {
	refractionRatio := dm.refractionIndex
	if !hit.frontFace {
		refractionRatio = 1 / refractionRatio
	}
	unitDirection := vector_unit(rayIn.direction)

	cosTheta := vector_dot(unitDirection.vector_opsite(), hit.normal)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	var direction Vect3
	if refractionRatio*sinTheta > 1.0 {
		direction = reflect(unitDirection, hit.normal)
	} else {
		// if reflectance_aproximation(cosTheta, refractionRatio) > rand.Float64() {
		// 	direction = reflect(unitDirection, hit.normal)
		// } else {
		direction = refract(unitDirection, hit.normal, refractionRatio)
		// }
	}
	*rayScattered = Ray{
		hit.p,
		direction,
		rayIn.remainingBounces - 1,
		rayIn.incomingLight,
		rayIn.color.mult(dm.color),
	}
	return true
}

type MatteMaterial struct {
	reflectance float64
	color       Color
}

func (mm MatteMaterial) scatter(rayIn *Ray, hit HitRecord, rayScattered *Ray) bool {

	if rand.Float64() < mm.reflectance {
		direction := vector_add(hit.normal, vector_unit(vector_random(-1, 1)))
		*rayScattered = Ray{
			hit.p,
			direction,
			rayIn.remainingBounces - 1,
			rayIn.incomingLight,
			rayIn.color.mult(mm.color),
		}
		return true
	}
	(*rayIn).color = (*&rayIn.color).mult(mm.color)
	return false
}
