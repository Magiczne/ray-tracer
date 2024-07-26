package material

import (
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/texture"
	"ray-tracer/vector"
)

type Lambertian struct {
	texture core.Texture
}

func NewLambertian(albedo *color.Color) *Lambertian {
	return &Lambertian{
		texture: texture.NewSolidColor(albedo),
	}
}

func NewTexturedLambertian(texture core.Texture) *Lambertian {
	return &Lambertian{
		texture: texture,
	}
}

func (l *Lambertian) Emitted(u, v float64, point *vector.Point3) *color.Color {
	return color.Black()
}

func (l *Lambertian) Scatter(rayIn *core.Ray, hitRecord *core.HitRecord, attenuation *color.Color, scattered *core.Ray) bool {
	scatterDirection := hitRecord.Normal.Add(vector.RandomUnitVector())

	// Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection.CopyFrom(hitRecord.Normal)
	}

	scattered.CopyFrom(core.NewTimedRay(hitRecord.Point, scatterDirection, rayIn.Time))
	attenuation.CopyFrom(l.texture.Value(hitRecord.U, hitRecord.V, hitRecord.Point))

	return true
}
