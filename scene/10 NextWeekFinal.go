package scene

import (
	"os"
	"ray-tracer/bvh"
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/material"
	"ray-tracer/object"
	"ray-tracer/random"
	"ray-tracer/texture"
	"ray-tracer/transform"
	"ray-tracer/vector"
	"ray-tracer/writer"
)

func CreateNextWeekFinalScene(imageWidth, samplesPerPixel, maxDepth int) func() {
	return func() {
		boxes1 := core.EmptyHittableList()
		ground := material.NewLambertian(color.NewColor(0.48, 0.83, 0.53))

		boxesPerSide := 20

		for i := range boxesPerSide {
			for j := range boxesPerSide {
				w := 100.0
				x0 := -1000.0 + float64(i)*w
				z0 := -1000.0 + float64(j)*w
				y0 := 0.0
				x1 := x0 + w
				y1 := random.Float64(1, 101)
				z1 := z0 + w

				boxes1.Add(object.NewBox(vector.NewPoint3(x0, y0, z0), vector.NewPoint3(x1, y1, z1), ground))
			}
		}

		world := core.EmptyHittableList()
		world.Add(bvh.NewBVHNode(boxes1.Objects()))

		light := material.NewDiffuseLight(color.NewColor(7, 7, 7))
		world.Add(object.NewQuad(vector.NewPoint3(123, 554, 147), vector.NewVector3(300, 0, 0), vector.NewVector3(0, 0, 265), light))

		center1 := vector.NewPoint3(400, 400, 200)
		center2 := center1.Add(vector.NewVector3(30, 0, 0))

		sphereMaterial := material.NewLambertian(color.NewColor(0.7, 0.3, 0.1))
		world.Add(object.NewMovingSphere(center1, center2, 50, sphereMaterial))

		world.Add(object.NewSphere(vector.NewPoint3(260, 150, 45), 50, material.NewDielectric(1.5)))
		world.Add(object.NewSphere(vector.NewPoint3(0, 150, 145), 50, material.NewMetal(color.NewColor(0.8, 0.8, 0.9), 1.0)))

		boundary := object.NewSphere(vector.NewPoint3(360, 150, 145), 70, material.NewDielectric(1.5))
		world.Add(boundary)
		world.Add(object.NewConstantMedium(boundary, 0.2, color.NewColor(0.2, 0.4, 0.9)))
		boundary = object.NewSphere(vector.NewPoint3(0, 0, 0), 5000, material.NewDielectric(1.5))
		world.Add(object.NewConstantMedium(boundary, .0001, color.NewColor(1, 1, 1)))

		emat := material.NewTexturedLambertian(texture.NewImage("assets/earthmap.jpg"))
		world.Add(object.NewSphere(vector.NewPoint3(400, 200, 400), 100, emat))
		pertext := texture.NewNoise(0.2)
		world.Add(object.NewSphere(vector.NewPoint3(220, 280, 300), 80, material.NewTexturedLambertian(pertext)))

		boxes2 := core.EmptyHittableList()
		white := material.NewLambertian(color.NewColor(.73, .73, .73))
		ns := 1000
		for range ns {
			boxes2.Add(object.NewSphere(vector.RandomVector3(0, 165), 10, white))
		}

		world.Add(
			transform.NewTranslate(
				transform.NewRotateY(
					bvh.NewBVHNode(boxes2.Objects()),
					15,
				),
				vector.NewVector3(-100, 270, 395),
			),
		)

		writer := writer.NewWriter(os.Args[1])

		// Camera
		camera := core.NewCamera()
		camera.AspectRatio = 1.0
		camera.ImageWidth = imageWidth
		camera.SamplesPerPixel = samplesPerPixel
		camera.MaxDepth = maxDepth
		camera.Background = color.NewColor(0, 0, 0)

		camera.VerticalFieldOfView = 40
		camera.LookFrom = vector.NewPoint3(478, 278, -600)
		camera.LookAt = vector.NewPoint3(278, 278, 0)
		camera.VectorUp = vector.NewVector3(0, 1, 0)

		camera.DefocusAngle = 0
		camera.FocusDistance = 10

		camera.Render(world, writer)
	}
}
