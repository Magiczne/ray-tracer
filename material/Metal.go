package material

import (
	"math"
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/vector"
)

type Metal struct {
	albedo *color.Color
	fuzz   float64
}

func NewMetal(albedo *color.Color, fuzz float64) *Metal {
	return &Metal{
		albedo: albedo,
		fuzz:   math.Min(fuzz, 1),
	}
}

func (m *Metal) Emitted(u, v float64, point *vector.Point3) *color.Color {
	return color.Black()
}

func (m *Metal) Scatter(rayIn *core.Ray, hitRecord *core.HitRecord) (bool, *core.Ray, *color.Color) {
	reflected := vector.Reflect(rayIn.Direction, hitRecord.Normal)
	reflected = vector.UnitVector(reflected).Add(vector.RandomUnitVector().MultiplyBy(m.fuzz))

	scattered := core.NewTimedRay(hitRecord.Point, reflected, rayIn.Time)
	attenuation := m.albedo

	if vector.DotProduct(scattered.Direction, hitRecord.Normal) > 0 {
		return true, scattered, attenuation
	}

	return false, nil, nil
}
