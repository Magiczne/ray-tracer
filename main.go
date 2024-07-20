package main

import (
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/material"
	"ray-tracer/object"
	"ray-tracer/vector"
)

// TODO: Zaczynać od 12

func main() {
	world := core.NewHittableList()

	groundMaterial := material.NewLambertian(color.NewColor(0.8, 0.8, 0))
	centerMaterial := material.NewLambertian(color.NewColor(0.1, 0.2, 0.5))
	leftMaterial := material.NewDielectric(1.5)
	bubbleMaterial := material.NewDielectric(1.0 / 1.5)
	rightMaterial := material.NewMetal(color.NewColor(0.8, 0.6, 0.2), 1.0)

	world.Add(object.NewSphere(*vector.NewPoint3(0.0, -100.5, -1.0), 100.0, groundMaterial))
	world.Add(object.NewSphere(*vector.NewPoint3(0.0, 0.0, -1.2), 0.5, centerMaterial))
	world.Add(object.NewSphere(*vector.NewPoint3(-1.0, 0.0, -1.0), 0.5, leftMaterial))
	world.Add(object.NewSphere(*vector.NewPoint3(-1.0, 0.0, -1.0), 0.4, bubbleMaterial))
	world.Add(object.NewSphere(*vector.NewPoint3(1.0, 0.0, -1.0), 0.5, rightMaterial))
	// world.Display()

	camera := core.NewCamera()
	camera.SetAspectRatio(16.0 / 9)
	camera.SetImageWidth(400)
	camera.SetSamplesPerPixel(100)
	camera.SetMaxDepth(50)
	camera.Render(world)
}
