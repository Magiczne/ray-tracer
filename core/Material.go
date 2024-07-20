package core

import (
	"ray-tracer/color"
)

type Material interface {
	Scatter(rayIn *Ray, hitRecord *HitRecord, attenuation *color.Color, scattered *Ray) bool
}
