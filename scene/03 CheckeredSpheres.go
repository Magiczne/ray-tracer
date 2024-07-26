package scene

import (
	"os"
	"ray-tracer/bvh"
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/material"
	"ray-tracer/object"
	"ray-tracer/texture"
	"ray-tracer/vector"
	"ray-tracer/writer"
)

func CheckeredSpheres() {
	world := core.EmptyHittableList()

	checkerTexture := texture.NewColoredChecker(0.32, color.NewColor(0.2, 0.3, 0.1), color.NewColor(0.9, 0.9, 0.9))
	checkerMaterial := material.NewTexturedLambertian(checkerTexture)

	world.Add(object.NewSphere(vector.NewPoint3(0.0, -10, 0), 10, checkerMaterial))
	world.Add(object.NewSphere(vector.NewPoint3(0.0, 10, 0), 10, checkerMaterial))

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
	camera.ImageWidth = 400      // 1200/400
	camera.SamplesPerPixel = 100 // 500/10
	camera.MaxDepth = 50

	camera.VerticalFieldOfView = 20
	camera.LookFrom = vector.NewPoint3(13, 2, 3)
	camera.LookAt = vector.NewPoint3(0, 0, 0)
	camera.VectorUp = vector.NewVector3(0, 1, 0)

	camera.DefocusAngle = 0
	camera.FocusDistance = 10

	camera.Render(bvhWorld, writer)
}
