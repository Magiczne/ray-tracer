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
	return color.White().MultiplyBy(n.perlin.Noise(point.MultiplyBy(n.scale)))
}
