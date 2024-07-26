package scene

import (
	"os"
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/material"
	"ray-tracer/object"
	"ray-tracer/vector"
	"ray-tracer/writer"
)

func Quads() {
	world := core.EmptyHittableList()

	leftRed := material.NewLambertian(color.NewColor(1.0, 0.2, 0.2))
	backGreen := material.NewLambertian(color.NewColor(0.2, 1.0, 0.2))
	rightBlue := material.NewLambertian(color.NewColor(0.2, 0.2, 1.0))
	upperOrange := material.NewLambertian(color.NewColor(1.0, 0.5, 0.0))
	lowerTeal := material.NewLambertian(color.NewColor(0.2, 0.8, 0.8))

	world.Add(object.NewQuad(vector.NewPoint3(-3, -2, 5), vector.NewVector3(0, 0, -4), vector.NewVector3(0, 4, 0), leftRed))
	world.Add(object.NewQuad(vector.NewPoint3(-2, -2, 0), vector.NewVector3(4, 0, 0), vector.NewVector3(0, 4, 0), 0.5, backGreen))
	world.Add(object.NewQuad(vector.NewPoint3(3, -2, 1), vector.NewVector3(0, 0, 4), vector.NewVector3(0, 4, 0), rightBlue))
	world.Add(object.NewQuad(vector.NewPoint3(-2, 3, 1), vector.NewVector3(4, 0, 0), vector.NewVector3(0, 0, 4), upperOrange))
	world.Add(object.NewQuad(vector.NewPoint3(-2, -3, 5), vector.NewVector3(4, 0, 0), vector.NewVector3(0, 0, -4), lowerTeal))

	writer := writer.NewWriter(os.Args[1])

	// Camera
	camera := core.NewCamera()
	camera.AspectRatio = 1.0
	camera.ImageWidth = 400
	camera.SamplesPerPixel = 20
	camera.MaxDepth = 50

	camera.VerticalFieldOfView = 80
	camera.LookFrom = vector.NewPoint3(0, 0, 9)
	camera.LookAt = vector.NewPoint3(0, 0, 0)
	camera.VectorUp = vector.NewVector3(0, 1, 0)

	camera.DefocusAngle = 0

	camera.Render(world, writer)
}
