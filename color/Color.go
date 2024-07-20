package color

import (
	"fmt"
	"os"
	"ray-tracer/math"
	"ray-tracer/util"
)

var intensity = util.NewInterval(0.0, 0.999)

type Color struct {
	r, g, b float64
}

func NewColor(r float64, g float64, b float64) *Color {
	return &Color{r, g, b}
}

func (c *Color) CopyFrom(other *Color) {
	c.r = other.r
	c.g = other.g
	c.b = other.b
}

func (c *Color) Write() {
	r := math.LinearToGamma(c.r)
	g := math.LinearToGamma(c.g)
	b := math.LinearToGamma(c.b)

	rByte := int(256 * intensity.Clamp(r))
	gByte := int(256 * intensity.Clamp(g))
	bByte := int(256 * intensity.Clamp(b))

	fmt.Fprintf(os.Stdout, "%d %d %d\n", rByte, gByte, bByte)
}

func (c *Color) Add(other *Color) *Color {
	return NewColor(c.r+other.r, c.g+other.g, c.b+other.b)
}

func (c *Color) Multiply(other *Color) *Color {
	return NewColor(c.r*other.r, c.g*other.g, c.b*other.b)
}

func (c *Color) AddInPlace(other *Color) {
	c.r += other.r
	c.g += other.g
	c.b += other.b
}

func (c *Color) MultiplyBy(multiplier float64) *Color {
	return NewColor(c.r*multiplier, c.g*multiplier, c.b*multiplier)
}
