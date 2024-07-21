package color

import "ray-tracer/random"

func RandomColor(min float64, max float64) *Color {
	return &Color{
		r: random.Float64(min, max),
		g: random.Float64(min, max),
		b: random.Float64(min, max),
	}
}
