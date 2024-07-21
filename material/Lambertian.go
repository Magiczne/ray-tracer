package material

import (
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/vector"
)

type Lambertian struct {
	core.Material
	albedo *color.Color
}

func NewLambertian(albedo *color.Color) *Lambertian {
	return &Lambertian{
		albedo: albedo,
	}
}

func (l *Lambertian) Scatter(rayIn *core.Ray, hitRecord *core.HitRecord, attenuation *color.Color, scattered *core.Ray) bool {
	scatterDirection := hitRecord.Normal.Add(vector.RandomUnitVector())

	// Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection.CopyFrom(&hitRecord.Normal)
	}

	scattered.CopyFrom(core.NewTimedRay(hitRecord.Point, *scatterDirection, rayIn.Time))
	attenuation.CopyFrom(l.albedo)

	return true
}
