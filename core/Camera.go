package core

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"ray-tracer/color"
	"ray-tracer/util"
	"ray-tracer/vector"
)

type Camera struct {
	aspectRatio       float64
	imageWidth        int
	imageHeight       int
	samplesPerPixel   int
	maxDepth          int
	center            vector.Point3
	pixel00Location   vector.Point3
	pixelDeltaU       vector.Vector3 // Offset to pixel to the right
	pixelDeltaV       vector.Vector3 // Offset to pixel below
	pixelSamplesScale float64        // Color scale factor for a sum of pixel samples
}

func NewCamera() *Camera {
	return &Camera{
		aspectRatio:     1,
		imageWidth:      100,
		samplesPerPixel: 10,
		maxDepth:        10,
	}
}

func (c *Camera) SetAspectRatio(aspectRatio float64) {
	c.aspectRatio = aspectRatio
}

func (c *Camera) SetImageWidth(imageWidth int) {
	c.imageWidth = imageWidth
}

func (c *Camera) SetSamplesPerPixel(samplesPerPixel int) {
	c.samplesPerPixel = samplesPerPixel
}

func (c *Camera) SetMaxDepth(depth int) {
	c.maxDepth = depth
}

func (c *Camera) Render(world Hittable) {
	c.initialize()

	fmt.Println("P3")
	fmt.Printf("%d %d\n", c.imageWidth, c.imageHeight)
	fmt.Println(255)

	for j := range c.imageHeight {
		fmt.Fprintf(os.Stderr, "Scanline remaining: %d\n", c.imageHeight-j)

		for i := range c.imageWidth {
			pixelColor := color.NewColor(0, 0, 0)

			for range c.samplesPerPixel {
				ray := c.getRay(i, j)
				pixelColor.AddInPlace(c.rayColor(ray, c.maxDepth, world))
			}

			pixelColor.MultiplyBy(c.pixelSamplesScale).Write()
		}
	}

	fmt.Fprintf(os.Stderr, "Done\n")
}

func (c *Camera) initialize() {
	c.imageHeight = int(float64(c.imageWidth) / c.aspectRatio)
	c.center = *vector.NewPoint3(0, 0, 0)

	c.pixelSamplesScale = 1.0 / float64(c.samplesPerPixel)

	// Viewport
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(c.imageWidth) / float64(c.imageHeight))

	// Edge vectors
	viewportU := vector.NewVec3(viewportWidth, 0, 0)
	viewportV := vector.NewVec3(0, -viewportHeight, 0)

	// Delta vectors from pixel to pixel
	c.pixelDeltaU = *viewportU.Divide(float64(c.imageWidth))
	c.pixelDeltaV = *viewportV.Divide(float64(c.imageHeight))

	// Location of the left upper pixel
	viewportUpperLeft := c.center.Substract(vector.NewVec3(0, 0, focalLength)).Substract(viewportU.Divide(2)).Substract(viewportV.Divide(2))
	c.pixel00Location = *viewportUpperLeft.Add(c.pixelDeltaU.Add(&c.pixelDeltaV).MultiplyBy(0.5))
}

func (c *Camera) getRay(i int, j int) *Ray {
	offset := c.sampleSquare()

	pixelSample := c.pixel00Location.Add(c.pixelDeltaU.MultiplyBy(float64(i) + offset.X())).Add(c.pixelDeltaV.MultiplyBy(float64(j) + offset.Y()))
	rayOrigin := c.center
	rayDirection := pixelSample.Substract(&rayOrigin)

	return NewRay(rayOrigin, *rayDirection)
}

func (c *Camera) rayColor(ray *Ray, depth int, world Hittable) *color.Color {
	if depth <= 0 {
		return color.NewColor(0, 0, 0)
	}

	hitRecord := NewHitRecord()

	if world.Hit(ray, util.NewInterval(0.001, math.Inf(1)), hitRecord) {
		scattered := EmptyRay()
		attenuation := color.Black()

		if hitRecord.Material.Scatter(ray, hitRecord, attenuation, scattered) {
			return attenuation.Multiply(c.rayColor(scattered, depth-1, world))
		}

		return color.Black()
	}

	unitDirection := vector.UnitVector(&ray.Direction)
	a := 0.5 * (unitDirection.Y() + 1.0)

	// TODO: AddInPlace?
	return color.NewColor(1, 1, 1).MultiplyBy(1.0 - a).Add(color.NewColor(0.5, 0.7, 1.0).MultiplyBy(a))
}

func (c *Camera) sampleSquare() *vector.Vector3 {
	return vector.NewVec3(
		rand.Float64()-0.5,
		rand.Float64()-0.5,
		0,
	)
}
