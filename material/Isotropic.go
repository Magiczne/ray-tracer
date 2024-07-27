package material

import (
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/texture"
	"ray-tracer/vector"
)

type Isotropic struct {
	texture core.Texture
}

func NewIsotropic(albedo *color.Color) *Isotropic {
	return NewTexturedIsotropic(texture.NewSolidColor(albedo))
}

func NewTexturedIsotropic(texture core.Texture) *Isotropic {
	return &Isotropic{
		texture: texture,
	}
}

func (i *Isotropic) Emitted(u, v float64, point *vector.Point3) *color.Color {
	return color.Black()
}

func (i *Isotropic) Scatter(ray *core.Ray, hitRecord *core.HitRecord, attenuation *color.Color, scattered *core.Ray) bool {
	scattered.CopyFrom(core.NewTimedRay(hitRecord.Point, vector.RandomUnitVector(), ray.Time))
	attenuation.CopyFrom(i.texture.Value(hitRecord.U, hitRecord.V, hitRecord.Point))

	return true
}
