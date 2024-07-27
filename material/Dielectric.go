package material

import (
	"math"
	"math/rand/v2"
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/util"
	"ray-tracer/vector"
)

type Dielectic struct {
	// Refractive index in vacuum or air, or the ratio of the material's refractive index over
	// the refractive index of the enclosing media
	refractionIndex float64
}

func NewDielectric(refractionIndex float64) *Dielectic {
	return &Dielectic{
		refractionIndex: refractionIndex,
	}
}

func (d *Dielectic) Emitted(u, v float64, point *vector.Point3) *color.Color {
	return color.Black()
}

func (d *Dielectic) Scatter(rayIn *core.Ray, hitRecord *core.HitRecord) (bool, *core.Ray, *color.Color) {
	attenuation := color.White()

	refractionIndex := d.refractionIndex
	if hitRecord.FrontFace {
		refractionIndex = 1 / d.refractionIndex
	}

	unitDirection := vector.UnitVector(rayIn.Direction)
	cosTheta := math.Min(vector.DotProduct(unitDirection.Negate(), hitRecord.Normal), 1)
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)

	cannotRefract := refractionIndex*sinTheta > 1.0

	var direction *vector.Vector3
	if cannotRefract || util.Reflectance(cosTheta, refractionIndex) > rand.Float64() {
		direction = vector.Reflect(unitDirection, hitRecord.Normal)
	} else {
		direction = vector.Refract(unitDirection, hitRecord.Normal, refractionIndex)
	}

	scattered := core.NewTimedRay(hitRecord.Point, direction, rayIn.Time)

	return true, scattered, attenuation
}
