package core

import (
	"ray-tracer/color"
	"ray-tracer/vector"
)

type Material interface {
	Emitted(u, v float64, point *vector.Point3) *color.Color
	Scatter(ray *Ray, hitRecord *HitRecord, attenuation *color.Color, scattered *Ray) bool
}
