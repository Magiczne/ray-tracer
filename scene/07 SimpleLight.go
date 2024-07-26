package scene

import (
	"os"
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/material"
	"ray-tracer/object"
	"ray-tracer/texture"
	"ray-tracer/vector"
	"ray-tracer/writer"
)

func SimpleLight() {
	world := core.EmptyHittableList()

	noiseTexture := texture.NewNoise(4)
	noiseMaterial := material.NewTexturedLambertian(noiseTexture)
	world.Add(object.NewSphere(vector.NewPoint3(0, -1000, 0), 1000, noiseMaterial))
	world.Add(object.NewSphere(vector.NewPoint3(0, 2, 0), 2, noiseMaterial))

	diffuseLight := material.NewDiffuseLight(color.NewColor(4, 4, 4))
	world.Add(object.NewSphere(vector.NewPoint3(0, 7, 0), 2, diffuseLight))
	world.Add(object.NewQuad(vector.NewPoint3(3, 1, -2), vector.NewVector3(2, 0, 0), vector.NewVector3(0, 2, 0), diffuseLight))

	writer := writer.NewWriter(os.Args[1])

	// Camera
	camera := core.NewCamera()
	camera.AspectRatio = 16.0 / 9
	camera.ImageWidth = 400
	camera.SamplesPerPixel = 100
	camera.MaxDepth = 50
	camera.Background = color.NewColor(0, 0, 0)

	camera.VerticalFieldOfView = 20
	camera.LookFrom = vector.NewPoint3(26, 3, 6)
	camera.LookAt = vector.NewPoint3(0, 2, 0)
	camera.VectorUp = vector.NewVector3(0, 1, 0)

	camera.DefocusAngle = 0
	camera.FocusDistance = 10

	camera.Render(world, writer)
}
