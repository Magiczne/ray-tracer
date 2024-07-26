package main

import (
	"math/rand"
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

func main() {
	switch 6 {
	case 1:
		materialTester()
		break

	case 2:
		bouncingSpheres()
		break

	case 3:
		checkeredSpheres()
		break

	case 4:
		earth()
		break

	case 5:
		perlinSpheres()
		break

	case 6:
		quads()
		break
	}
}

func materialTester() {
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

func bouncingSpheres() {
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

	camera.VerticalFieldOfView = 20
	camera.LookFrom = vector.NewPoint3(13, 2, 3)
	// camera.LookFrom = vector.NewPoint3(1, 15, 1)
	camera.LookAt = vector.NewPoint3(0, 0, 0)
	camera.VectorUp = vector.NewVector3(0, 1, 0)

	camera.DefocusAngle = 0.6
	camera.FocusDistance = 10.0

	camera.Render(bvhWorld, writer)
}

func checkeredSpheres() {
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

func earth() {
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

func perlinSpheres() {
	world := core.EmptyHittableList()

	noiseTexture := texture.NewNoise(4)
	noiseMaterial := material.NewTexturedLambertian(noiseTexture)
	world.Add(object.NewSphere(vector.NewPoint3(0, -1000, 0), 1000, noiseMaterial))
	world.Add(object.NewSphere(vector.NewPoint3(0, 2, 0), 2, noiseMaterial))

	writer := writer.NewWriter(os.Args[1])

	// Camera
	camera := core.NewCamera()
	camera.AspectRatio = 16.0 / 9
	camera.ImageWidth = 400
	camera.SamplesPerPixel = 100
	camera.MaxDepth = 50

	camera.VerticalFieldOfView = 20
	camera.LookFrom = vector.NewPoint3(13, 2, 3)
	camera.LookAt = vector.NewPoint3(0, 0, 0)
	camera.VectorUp = vector.NewVector3(0, 1, 0)

	camera.DefocusAngle = 0
	camera.FocusDistance = 10

	camera.Render(world, writer)
}

func quads() {
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
