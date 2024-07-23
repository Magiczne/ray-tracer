package texture

import (
	"ray-tracer/color"
	"ray-tracer/image"
	"ray-tracer/util"
	"ray-tracer/vector"
)

type Image struct {
	image *image.Image
}

func NewImage(filename string) *Image {
	return &Image{
		image: image.NewImage(filename),
	}
}

func (img *Image) Value(u, v float64, point *vector.Point3) *color.Color {
	if img.image.Height <= 0 {
		// Debugging basically returning some color which will always be visible
		return color.Cyan()
	}

	u = util.NewInterval(0, 1).Clamp(u)
	v = 1.0 - util.NewInterval(0, 1).Clamp(v)

	i := int(u * float64(img.image.Width))
	j := int(v * float64(img.image.Height))
	pixel := img.image.PixelAt(i, j)

	colorScale := 1.0 / 255.0

	return color.NewColor(
		colorScale*pixel.R,
		colorScale*pixel.G,
		colorScale*pixel.B,
	)
}
