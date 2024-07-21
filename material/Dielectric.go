package material

import (
	"math"
	"math/rand"
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/vector"
)

type Dielectic struct {
	core.Material

	// Refractive index in vacuum or air, or the ratio of the material's refractive index over
	// the refractive index of the enclosing media
	refractionIndex float64
}

func NewDielectric(refractionIndex float64) *Dielectic {
	return &Dielectic{
		refractionIndex: refractionIndex,
	}
}

func (d *Dielectic) Scatter(rayIn *core.Ray, hitRecord *core.HitRecord, attenuation *color.Color, scattered *core.Ray) bool {
	attenuation.CopyFrom(color.White())

	refractionIndex := d.refractionIndex
	if hitRecord.FrontFace {
		refractionIndex = 1 / d.refractionIndex
	}

	unitDirection := vector.UnitVector(&rayIn.Direction)
	cosTheta := math.Min(vector.DotProduct(unitDirection.MultiplyBy(-1), &hitRecord.Normal), 1)
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)

	cannotRefract := refractionIndex*sinTheta > 1.0

	var direction *vector.Vector3
	if cannotRefract || reflectance(cosTheta, refractionIndex) > rand.Float64() {
		direction = vector.Reflect(unitDirection, &hitRecord.Normal)
	} else {
		direction = vector.Refract(unitDirection, &hitRecord.Normal, refractionIndex)
	}

	scattered.CopyFrom(core.NewTimedRay(hitRecord.Point, *direction, rayIn.Time))

	return true
}

func reflectance(cos float64, refractionIndex float64) float64 {
	// Use Schlick's approximation for reflectance.
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 = r0 * r0

	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
