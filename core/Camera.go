package core

import (
	"fmt"
	"math"
	"math/rand/v2"
	"ray-tracer/color"
	"ray-tracer/util"
	"ray-tracer/vector"
	"ray-tracer/writer"
	"time"
)

type Camera struct {
	AspectRatio     float64
	ImageWidth      int
	imageHeight     int
	SamplesPerPixel int
	MaxDepth        int
	Background      *color.Color

	VerticalFieldOfView float64
	LookFrom            *vector.Point3
	LookAt              *vector.Point3
	VectorUp            *vector.Vector3

	// Defocus (Depth of field)
	DefocusAngle  float64
	FocusDistance float64
	defocusDiskU  *vector.Vector3
	defocusDiskV  *vector.Vector3

	// Camera frame basis vectors
	u *vector.Vector3
	v *vector.Vector3
	w *vector.Vector3

	center            *vector.Point3
	pixel00Location   *vector.Point3
	pixelDeltaU       *vector.Vector3 // Offset to pixel to the right
	pixelDeltaV       *vector.Vector3 // Offset to pixel below
	pixelSamplesScale float64         // Color scale factor for a sum of pixel samples
}

func NewCamera() *Camera {
	return &Camera{
		AspectRatio:         1,
		ImageWidth:          100,
		SamplesPerPixel:     10,
		MaxDepth:            10,
		VerticalFieldOfView: 90,
		LookFrom:            vector.NewPoint3(0, 0, 0),
		LookAt:              vector.NewPoint3(0, 0, -1),
		VectorUp:            vector.NewVector3(0, 1, 0),
		DefocusAngle:        0,
		FocusDistance:       10,
	}
}

func (c *Camera) Render(world Hittable, writer *writer.Writer) {
	start := time.Now()

	c.initialize()

	writer.WriteHeader(c.ImageWidth, c.imageHeight)

	for j := range c.imageHeight {
		fmt.Printf("Scanline remaining: %d\n", c.imageHeight-j)

		for i := range c.ImageWidth {
			pixelColor := color.NewColor(0, 0, 0)

			for range c.SamplesPerPixel {
				ray := c.getRay(i, j)
				pixelColor.AddInPlace(c.rayColor(ray, c.MaxDepth, world))
			}

			writer.WriteColor(pixelColor.MultiplyBy(c.pixelSamplesScale))
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)
}

func (c *Camera) initialize() {
	c.imageHeight = int(float64(c.ImageWidth) / c.AspectRatio)

	c.pixelSamplesScale = 1.0 / float64(c.SamplesPerPixel)
	c.center = c.LookFrom

	// Determine viewport dimensions
	theta := util.DegToRad(c.VerticalFieldOfView)
	h := math.Tan(theta / 2)
	viewportHeight := 2 * h * c.FocusDistance
	viewportWidth := viewportHeight * (float64(c.ImageWidth) / float64(c.imageHeight))

	// Calculate the uvw unit basis vectors for the camera coordinate frame
	c.w = vector.UnitVector(c.LookFrom.Substract(c.LookAt))
	c.u = vector.UnitVector(vector.CrossProduct(c.VectorUp, c.w))
	c.v = vector.CrossProduct(c.w, c.u)

	// Edge vectors
	viewportU := c.u.MultiplyBy(viewportWidth)
	viewportV := c.v.Negate().MultiplyBy(viewportHeight)

	// Delta vectors from pixel to pixel
	c.pixelDeltaU = viewportU.Divide(float64(c.ImageWidth))
	c.pixelDeltaV = viewportV.Divide(float64(c.imageHeight))

	// Location of the left upper pixel
	viewportUpperLeft := c.center.Substract(c.w.MultiplyBy(c.FocusDistance)).Substract(viewportU.Divide(2)).Substract(viewportV.Divide(2))
	c.pixel00Location = viewportUpperLeft.Add(c.pixelDeltaU.Add(c.pixelDeltaV).MultiplyBy(0.5))

	// Calculate the camera defocus disk basis vectors.
	defocusRadius := c.FocusDistance * math.Tan(util.DegToRad(c.DefocusAngle/2))
	c.defocusDiskU = c.u.MultiplyBy(defocusRadius)
	c.defocusDiskV = c.v.MultiplyBy(defocusRadius)
}

func (c *Camera) getRay(i int, j int) *Ray {
	offset := c.sampleSquare()

	pixelSample := c.pixel00Location.Add(c.pixelDeltaU.MultiplyBy(float64(i) + offset.X)).Add(c.pixelDeltaV.MultiplyBy(float64(j) + offset.Y))

	rayOrigin := c.center
	if c.DefocusAngle > 0 {
		rayOrigin = c.defocusDiskSample()
	}

	rayDirection := pixelSample.Substract(rayOrigin)

	// We're forcing the ray tracer to render frame as it was starting in t=0 and ending in t=1,
	// so we just generate random times between 0 and 1 for our rays.
	rayTime := rand.Float64()

	return NewTimedRay(rayOrigin, rayDirection, rayTime)
}

func (c *Camera) rayColor(ray *Ray, depth int, world Hittable) *color.Color {
	if depth <= 0 {
		return color.Black()
	}

	hitRecord := world.Hit(ray, util.NewInterval(0.001, math.Inf(1)))

	if hitRecord != nil {
		color := hitRecord.Material.Emitted(hitRecord.U, hitRecord.V, hitRecord.Point)

		if ok, scattered, attenuation := hitRecord.Material.Scatter(ray, hitRecord); ok {
			scatterColor := c.rayColor(scattered, depth-1, world)
			color = color.Add(scatterColor.Multiply(attenuation))
		}

		return color
	}

	return c.Background
}

func (c *Camera) sampleSquare() *vector.Vector3 {
	return vector.NewVector3(
		rand.Float64()-0.5,
		rand.Float64()-0.5,
		0,
	)
}

func (c *Camera) defocusDiskSample() *vector.Point3 {
	p := vector.RandomVector3InUnitDisk()

	return c.center.Add(c.defocusDiskU.MultiplyBy(p.X)).Add(c.defocusDiskV.MultiplyBy(p.Y))
}
