package color

import (
	"ray-tracer/util"
)

var intensity = util.NewInterval(0.0, 0.999)

type Color struct {
	R, G, B float64
}

func NewColor(r float64, g float64, b float64) *Color {
	return &Color{r, g, b}
}

func (c *Color) CopyFrom(other *Color) {
	c.R = other.R
	c.G = other.G
	c.B = other.B
}

func (c *Color) ToRgbBytes() (int, int, int) {
	r := util.LinearToGamma(c.R)
	g := util.LinearToGamma(c.G)
	b := util.LinearToGamma(c.B)

	rByte := int(256 * intensity.Clamp(r))
	gByte := int(256 * intensity.Clamp(g))
	bByte := int(256 * intensity.Clamp(b))

	return rByte, gByte, bByte
}

func (c *Color) Add(other *Color) *Color {
	return NewColor(c.R+other.R, c.G+other.G, c.B+other.B)
}

func (c *Color) Multiply(other *Color) *Color {
	return NewColor(c.R*other.R, c.G*other.G, c.B*other.B)
}

func (c *Color) AddInPlace(other *Color) {
	c.R += other.R
	c.G += other.G
	c.B += other.B
}

func (c *Color) MultiplyBy(multiplier float64) *Color {
	return NewColor(c.R*multiplier, c.G*multiplier, c.B*multiplier)
}
