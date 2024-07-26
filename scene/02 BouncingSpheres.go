package scene

import (
	"math/rand/v2"
	"os"
	"ray-tracer/bvh"
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/material"
	"ray-tracer/object"
	"ray-tracer/random"
	"ray-tracer/texture"
	"ray-tracer/vector"
	"ray-tracer/writer"
)

func BouncingSpheres() {
	smallSphereCoefficient := 11
	world := core.EmptyHittableList()

	groundTexture := texture.NewColoredChecker(0.32, color.NewColor(0.2, 0.3, 0.1), color.NewColor(0.9, 0.9, 0.9))
	groundMaterial := material.NewTexturedLambertian(groundTexture)
	world.Add(object.NewSphere(vector.NewPoint3(0, -1000, 0), 1000, groundMaterial))

	for a := -smallSphereCoefficient; a < smallSphereCoefficient; a++ {
		for b := -smallSphereCoefficient; b < smallSphereCoefficient; b++ {
			chooseMat := rand.Float64()
			center := vector.NewPoint3(float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64())

			if center.Substract(vector.NewPoint3(4, 0.2, 0)).Length() > 0.9 {
				if chooseMat < 0.8 {
					// diffuse
					albedo := color.RandomColor(0, 1).Multiply(color.RandomColor(0, 1))
					material := material.NewLambertian(albedo)
					center2 := center.Add(vector.NewVector3(0, random.Float64(0, 0.5), 0))
					world.Add(object.NewMovingSphere(center, center2, 0.2, material))
				} else if chooseMat < 0.95 {
					// metal
					albedo := color.RandomColor(0.5, 1)
					fuzz := random.Float64(0, 0.5)
					material := material.NewMetal(albedo, fuzz)
					world.Add(object.NewSphere(center, 0.2, material))
				} else {
					// glass
					material := material.NewDielectric(1.5)
					world.Add(object.NewSphere(center, 0.2, material))
				}
			}
		}
	}

	material1 := material.NewDielectric(1.5)
	world.Add(object.NewSphere(vector.NewPoint3(0, 1, 0), 1.0, material1))

	material2 := material.NewLambertian(color.NewColor(0.4, 0.2, 0.1))
	world.Add(object.NewSphere(vector.NewPoint3(-4, 1, 0), 1.0, material2))

	material3 := material.NewMetal(color.NewColor(0.7, 0.6, 0.5), 0.0)
	world.Add(object.NewSphere(vector.NewPoint3(4, 1, 0), 1.0, material3))

	bvhWorld := core.EmptyHittableList()
	bvhWorld.Add(bvh.NewBVHNode(world.Objects()))

	// world.Display()

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
	camera.Background = color.NewColor(0.7, 0.8, 1.0)

	camera.VerticalFieldOfView = 20
	camera.LookFrom = vector.NewPoint3(13, 2, 3)
	// camera.LookFrom = vector.NewPoint3(1, 15, 1)
	camera.LookAt = vector.NewPoint3(0, 0, 0)
	camera.VectorUp = vector.NewVector3(0, 1, 0)

	camera.DefocusAngle = 0.6
	camera.FocusDistance = 10.0

	camera.Render(bvhWorld, writer)
}
