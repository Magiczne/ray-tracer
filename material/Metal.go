package material

import (
	"math"
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/vector"
)

type Metal struct {
	core.Material
	albedo *color.Color
	fuzz   float64
}

func NewMetal(albedo *color.Color, fuzz float64) *Metal {
	return &Metal{
		albedo: albedo,
		fuzz:   math.Min(fuzz, 1),
	}
}

func (m *Metal) Scatter(rayIn *core.Ray, hitRecord *core.HitRecord, attenuation *color.Color, scattered *core.Ray) bool {
	// TODO: Coś z pointerami posrane
	reflected := vector.Reflect(&rayIn.Direction, &hitRecord.Normal)
	reflected = vector.UnitVector(reflected).Add(vector.RandomUnitVector().MultiplyBy(m.fuzz))

	scattered.CopyFrom(core.NewRay(hitRecord.Point, *reflected))
	attenuation.CopyFrom(m.albedo)

	return vector.DotProduct(&scattered.Direction, &hitRecord.Normal) > 0
}
