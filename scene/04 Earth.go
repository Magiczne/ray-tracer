package scene

import (
	"os"
	"ray-tracer/core"
	"ray-tracer/material"
	"ray-tracer/object"
	"ray-tracer/texture"
	"ray-tracer/vector"
	"ray-tracer/writer"
)

func Earth() {
	world := core.EmptyHittableList()

	earthTexture := texture.NewImage("assets/earthmap.jpg")
	earthSurface := material.NewTexturedLambertian(earthTexture)
	globe := object.NewSphere(vector.NewPoint3(0, 0, 0), 2, earthSurface)

	world.Add(globe)

	writer := writer.NewWriter(os.Args[1])

	// Camera
	camera := core.NewCamera()
	camera.AspectRatio = 16.0 / 9
	camera.ImageWidth = 400
	camera.SamplesPerPixel = 100
	camera.MaxDepth = 50

	camera.VerticalFieldOfView = 20
	camera.LookFrom = vector.NewPoint3(0, 0, 12)
	camera.LookAt = vector.NewPoint3(0, 0, 0)
	camera.VectorUp = vector.NewVector3(0, 1, 0)

	camera.DefocusAngle = 0
	camera.FocusDistance = 10

	camera.Render(world, writer)
}
