package scene

import (
	"os"
	"ray-tracer/bvh"
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/material"
	"ray-tracer/object"
	"ray-tracer/vector"
	"ray-tracer/writer"
)

func MaterialTester() {
	world := core.EmptyHittableList()

	groundMaterial := material.NewLambertian(color.NewColor(0.8, 0.8, 0))
	centerMaterial := material.NewLambertian(color.NewColor(0.1, 0.2, 0.5))
	leftMaterial := material.NewDielectric(1.5)
	bubbleMaterial := material.NewDielectric(1.0 / 1.5)
	rightMaterial := material.NewMetal(color.NewColor(0.8, 0.6, 0.2), 1.0)

	world.Add(object.NewSphere(vector.NewPoint3(0.0, -100.5, -1.0), 100.0, groundMaterial))
	world.Add(object.NewSphere(vector.NewPoint3(0.0, 0.0, -1.2), 0.5, centerMaterial))
	world.Add(object.NewSphere(vector.NewPoint3(-1.0, 0.0, -1.0), 0.5, leftMaterial))
	world.Add(object.NewSphere(vector.NewPoint3(-1.0, 0.0, -1.0), 0.4, bubbleMaterial))
	world.Add(object.NewSphere(vector.NewPoint3(1.0, 0.0, -1.0), 0.5, rightMaterial))

	bvhWorld := core.EmptyHittableList()
	bvhWorld.Add(bvh.NewBVHNode(world.Objects()))

	// Writer
	if len(os.Args) < 2 {
		panic("You need to supply file name")
	}

	writer := writer.NewWriter(os.Args[1])

	// Camera
	camera := core.NewCamera()
	camera.AspectRatio = 16.0 / 9
	camera.ImageWidth = 400     // 1200/400
	camera.SamplesPerPixel = 50 // 500/10
	camera.MaxDepth = 50

	camera.VerticalFieldOfView = 30
	camera.LookFrom = vector.NewPoint3(-2, 2, 1)
	camera.LookAt = vector.NewPoint3(0, 0, -1)
	camera.VectorUp = vector.NewVector3(0, 1, 0)

	camera.DefocusAngle = 0
	camera.FocusDistance = 10

	camera.Render(bvhWorld, writer)
}
