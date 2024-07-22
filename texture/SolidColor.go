package texture

import (
	"ray-tracer/color"
	"ray-tracer/vector"
)

type SolidColor struct {
	albedo *color.Color
}

func NewSolidColor(albedo *color.Color) *SolidColor {
	return &SolidColor{
		albedo: albedo,
	}
}

func (sc *SolidColor) Value(u, v float64, point *vector.Point3) *color.Color {
	return sc.albedo
}
