package texture

import (
	"ray-tracer/color"
	"ray-tracer/perlin"
	"ray-tracer/vector"
)

type Noise struct {
	perlin *perlin.Perlin
	scale  float64
}

func NewNoise(scale float64) *Noise {
	return &Noise{
		perlin: perlin.NewPerlin(),
		scale:  scale,
	}
}

func (n *Noise) Value(u, v float64, point *vector.Point3) *color.Color {
	// Perlin interpolation can return [-1, 1], we need to map that to [0, 1]
	// return color.White().MultiplyBy(0.5).MultiplyBy(n.perlin.Noise(point.MultiplyBy(n.scale)) + 1)

	// Turbulence
	return color.White().MultiplyBy(n.perlin.Turbulence(point, 7))
}
