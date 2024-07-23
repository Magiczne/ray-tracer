package color

import "ray-tracer/random"

func RandomColor(min float64, max float64) *Color {
	return &Color{
		R: random.Float64(min, max),
		G: random.Float64(min, max),
		B: random.Float64(min, max),
	}
}
