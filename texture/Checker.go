package texture

import (
	"math"
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/vector"
)

type Checker struct {
	invertedScale float64
	even          core.Texture
	odd           core.Texture
}

func NewTexturedChecker(scale float64, even, odd core.Texture) *Checker {
	return &Checker{
		invertedScale: 1.0 / scale,
		even:          even,
		odd:           odd,
	}
}

func NewColoredChecker(scale float64, color1, color2 *color.Color) *Checker {
	return &Checker{
		invertedScale: 1.0 / scale,
		even:          NewSolidColor(color1),
		odd:           NewSolidColor(color2),
	}
}

func (c *Checker) Value(u, v float64, point *vector.Point3) *color.Color {
	xInt := int(math.Floor(c.invertedScale * point.X()))
	yInt := int(math.Floor(c.invertedScale * point.Y()))
	zInt := int(math.Floor(c.invertedScale * point.Z()))
	isEven := (xInt+yInt+zInt)%2 == 0

	if isEven {
		return c.even.Value(u, v, point)
	}

	return c.odd.Value(u, v, point)
}
