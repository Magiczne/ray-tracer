package image

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"ray-tracer/color"
)

type Image struct {
	data   image.Image
	Width  int
	Height int
}

func NewImage(filename string) *Image {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return &Image{
		Width:  image.Bounds().Dx(),
		Height: image.Bounds().Dy(),
		data:   image,
	}
}

func (img *Image) PixelAt(i, j int) *color.Color {
	pixel := img.data.At(i, j)
	rByte, gByte, bByte, _ := pixel.RGBA()
	r := float64(rByte) / 255.0
	g := float64(gByte) / 255.0
	b := float64(bByte) / 255.0

	return color.NewColor(r, g, b)
}
