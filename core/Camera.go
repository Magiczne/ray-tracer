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
	AspectRatio     float64
	ImageWidth      int
	imageHeight     int
	SamplesPerPixel int
	MaxDepth        int

	VerticalFieldOfView int
	LookFrom            vector.Point3
	LookAt              vector.Point3
	VectorUp            vector.Vector3

	// Camera frame basis vectors
	u vector.Vector3
	v vector.Vector3
	w vector.Vector3

	center            vector.Point3
	pixel00Location   vector.Point3
	pixelDeltaU       vector.Vector3 // Offset to pixel to the right
	pixelDeltaV       vector.Vector3 // Offset to pixel below
	pixelSamplesScale float64        // Color scale factor for a sum of pixel samples
}

func NewCamera() *Camera {
	return &Camera{
		AspectRatio:         1,
		ImageWidth:          100,
		SamplesPerPixel:     10,
		MaxDepth:            10,
		VerticalFieldOfView: 90,
		LookFrom:            *vector.NewPoint3(0, 0, 0),
		LookAt:              *vector.NewPoint3(0, 0, -1),
		VectorUp:            *vector.NewVec3(0, 1, 0),
	}
}

func (c *Camera) Render(world Hittable) {
	c.initialize()

	fmt.Println("P3")
	fmt.Printf("%d %d\n", c.ImageWidth, c.imageHeight)
	fmt.Println(255)

	for j := range c.imageHeight {
		fmt.Fprintf(os.Stderr, "Scanline remaining: %d\n", c.imageHeight-j)

		for i := range c.ImageWidth {
			pixelColor := color.NewColor(0, 0, 0)

			for range c.SamplesPerPixel {
				ray := c.getRay(i, j)
				pixelColor.AddInPlace(c.rayColor(ray, c.MaxDepth, world))
			}

			pixelColor.MultiplyBy(c.pixelSamplesScale).Write()
		}
	}

	fmt.Fprintf(os.Stderr, "Done\n")
}

func (c *Camera) initialize() {
	c.imageHeight = int(float64(c.ImageWidth) / c.AspectRatio)

	c.pixelSamplesScale = 1.0 / float64(c.SamplesPerPixel)
	c.center.CopyFrom(&c.LookFrom)

	// Determine viewport dimensions
	focalLength := c.LookFrom.Substract(&c.LookAt).Length()
	theta := util.DegToRad(float64(c.VerticalFieldOfView))
	h := math.Tan(theta / 2)
	viewportHeight := 2 * h * focalLength
	viewportWidth := viewportHeight * (float64(c.ImageWidth) / float64(c.imageHeight))

	// Calculate the uvw unit basis vectors for the camera coordinate frame
	c.w = *vector.UnitVector(c.LookFrom.Substract(&c.LookAt))
	c.u = *vector.UnitVector(vector.CrossProduct(&c.VectorUp, &c.w))
	c.v = *vector.CrossProduct(&c.w, &c.u)

	// Edge vectors
	viewportU := c.u.MultiplyBy(viewportWidth)
	viewportV := c.v.MultiplyBy(-1).MultiplyBy(viewportHeight)

	// Delta vectors from pixel to pixel
	c.pixelDeltaU = *viewportU.Divide(float64(c.ImageWidth))
	c.pixelDeltaV = *viewportV.Divide(float64(c.imageHeight))

	// Location of the left upper pixel
	viewportUpperLeft := c.center.Substract(c.w.MultiplyBy(focalLength)).Substract(viewportU.Divide(2)).Substract(viewportV.Divide(2))
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
