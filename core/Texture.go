package core

import (
	"ray-tracer/color"
	"ray-tracer/vector"
)

type Texture interface {
	Value(u, v float64, point *vector.Point3) *color.Color
}
