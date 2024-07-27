package core

import (
	"ray-tracer/color"
	"ray-tracer/vector"
)

type Material interface {
	Emitted(u, v float64, point *vector.Point3) *color.Color
	// returns: scatter, scattered ray, attenuation
	Scatter(ray *Ray, hitRecord *HitRecord) (bool, *Ray, *color.Color)
}
