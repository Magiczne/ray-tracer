package material

import (
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/texture"
	"ray-tracer/vector"
)

type DiffuseLight struct {
	texture core.Texture
}

func NewDiffuseLight(emit *color.Color) *DiffuseLight {
	return &DiffuseLight{
		texture: texture.NewSolidColor(emit),
	}
}

func NewTexturedDiffuseLight(texture core.Texture) *DiffuseLight {
	return &DiffuseLight{
		texture: texture,
	}
}

func (dl *DiffuseLight) Emitted(u, v float64, point *vector.Point3) *color.Color {
	return dl.texture.Value(u, v, point)
}

func (dl *DiffuseLight) Scatter(rayIn *core.Ray, hitRecord *core.HitRecord) (bool, *core.Ray, *color.Color) {
	return false, nil, nil
}
