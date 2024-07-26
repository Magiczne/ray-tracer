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

func CornellBox() {
	world := core.EmptyHittableList()

	red := material.NewLambertian(color.NewColor(0.65, 0.05, 0.05))
	white := material.NewLambertian(color.NewColor(0.73, 0.73, 0.73))
	green := material.NewLambertian(color.NewColor(0.12, 0.45, 0.15))
	light := material.NewDiffuseLight(color.NewColor(15, 15, 15))

	world.Add(object.NewQuad(vector.NewPoint3(555, 0, 0), vector.NewVector3(0, 555, 0), vector.NewVector3(0, 0, 555), green))
	world.Add(object.NewQuad(vector.NewPoint3(0, 0, 0), vector.NewVector3(0, 555, 0), vector.NewVector3(0, 0, 555), red))
	world.Add(object.NewQuad(vector.NewPoint3(343, 554, 332), vector.NewVector3(-130, 0, 0), vector.NewVector3(0, 0, -105), light))
	world.Add(object.NewQuad(vector.NewPoint3(0, 0, 0), vector.NewVector3(555, 0, 0), vector.NewVector3(0, 0, 555), white))
	world.Add(object.NewQuad(vector.NewPoint3(555, 555, 555), vector.NewVector3(-555, 0, 0), vector.NewVector3(0, 0, -555), white))
	world.Add(object.NewQuad(vector.NewPoint3(0, 0, 555), vector.NewVector3(555, 0, 0), vector.NewVector3(0, 555, 0), white))

	writer := writer.NewWriter(os.Args[1])

	// Camera
	camera := core.NewCamera()
	camera.AspectRatio = 16.0 / 9
	camera.ImageWidth = 600
	camera.SamplesPerPixel = 200
	camera.MaxDepth = 50
	camera.Background = color.NewColor(0, 0, 0)

	camera.VerticalFieldOfView = 40
	camera.LookFrom = vector.NewPoint3(278, 278, -800)
	camera.LookAt = vector.NewPoint3(278, 278, 0)
	camera.VectorUp = vector.NewVector3(0, 1, 0)

	camera.DefocusAngle = 0
	camera.FocusDistance = 10

	camera.Render(world, writer)
}
