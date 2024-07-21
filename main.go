package main

import (
	"math/rand"
	"os"
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/material"
	"ray-tracer/object"
	"ray-tracer/random"
	"ray-tracer/vector"
	"ray-tracer/writer"
)

// TODO: Zaczynać od 12

func main() {
	world := core.NewHittableList()

	// groundMaterial := material.NewLambertian(color.NewColor(0.8, 0.8, 0))
	// centerMaterial := material.NewLambertian(color.NewColor(0.1, 0.2, 0.5))
	// leftMaterial := material.NewDielectric(1.5)
	// bubbleMaterial := material.NewDielectric(1.0 / 1.5)
	// rightMaterial := material.NewMetal(color.NewColor(0.8, 0.6, 0.2), 1.0)

	// world.Add(object.NewSphere(*vector.NewPoint3(0.0, -100.5, -1.0), 100.0, groundMaterial))
	// world.Add(object.NewSphere(*vector.NewPoint3(0.0, 0.0, -1.2), 0.5, centerMaterial))
	// world.Add(object.NewSphere(*vector.NewPoint3(-1.0, 0.0, -1.0), 0.5, leftMaterial))
	// world.Add(object.NewSphere(*vector.NewPoint3(-1.0, 0.0, -1.0), 0.4, bubbleMaterial))
	// world.Add(object.NewSphere(*vector.NewPoint3(1.0, 0.0, -1.0), 0.5, rightMaterial))

	groundMaterial := material.NewLambertian(color.NewColor(0.5, 0.5, 0.5))
	world.Add(object.NewSphere(*vector.NewPoint3(0, -1000, 0), 1000, groundMaterial))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := vector.NewPoint3(float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64())

			if center.Substract(vector.NewPoint3(4, 0.2, 0)).Length() > 0.9 {
				if chooseMat < 0.8 {
					// diffuse
					albedo := color.RandomColor(0, 1).Multiply(color.RandomColor(0, 1))
					material := material.NewLambertian(albedo)
					world.Add(object.NewSphere(*center, 0.2, material))
				} else if chooseMat < 0.95 {
					// metal
					albedo := color.RandomColor(0.5, 1)
					fuzz := random.Float64(0, 0.5)
					material := material.NewMetal(albedo, fuzz)
					world.Add(object.NewSphere(*center, 0.2, material))
				} else {
					// glass
					material := material.NewDielectric(1.5)
					world.Add(object.NewSphere(*center, 0.2, material))
				}
			}
		}
	}

	material1 := material.NewDielectric(1.5)
	world.Add(object.NewSphere(*vector.NewPoint3(0, 1, 0), 1.0, material1))

	material2 := material.NewLambertian(color.NewColor(0.4, 0.2, 0.1))
	world.Add(object.NewSphere(*vector.NewPoint3(-4, 1, 0), 1.0, material2))

	material3 := material.NewMetal(color.NewColor(0.7, 0.6, 0.5), 0.0)
	world.Add(object.NewSphere(*vector.NewPoint3(4, 1, 0), 1.0, material3))

	// world.Display()

	// Writer
	if len(os.Args) < 2 {
		panic("You need to supply file name")
	}

	writer := writer.NewWriter(os.Args[1])

	// Camera
	camera := core.NewCamera()
	camera.AspectRatio = 16.0 / 9
	camera.ImageWidth = 150     // 1200
	camera.SamplesPerPixel = 50 // 500
	camera.MaxDepth = 50

	camera.VerticalFieldOfView = 20
	camera.LookFrom = *vector.NewPoint3(13, 2, 3)
	camera.LookAt = *vector.NewPoint3(0, 0, 0)
	camera.VectorUp = *vector.NewVector3(0, 1, 0)

	camera.DefocusAngle = 0.6
	camera.FocusDistance = 10.0

	camera.Render(world, writer)
}
