package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"ray-tracer/core"
	"ray-tracer/object"
	"ray-tracer/util"
)

// TODO: Zaczynać od 6.8 (ale nie działa to co robilem od 6.3)

func rayColor(ray *core.Ray, world core.Hittable) *util.Color {
	hitRecord := core.NewHitRecord()

	if world.Hit(ray, 0, math.Inf(1), hitRecord) {
		return hitRecord.Normal().Add(*util.NewColor(1, 1, 1)).MultiplyBy(0.5)
	}

	unitDirection := ray.Direction().UnitVector()
	a := 0.5 * (unitDirection.Y() + 1.0)

	return util.NewColor(1, 1, 1).MultiplyBy(1.0 - a).Add(*util.NewColor(0.5, 0.7, 1.0).MultiplyBy(a))
}

func main() {
	// Logging
	stdErr := log.New(os.Stderr, "", 0)
	stdOut := log.New(os.Stdout, "", 0)

	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)

	// World
	world := core.HittableList{}
	world.Add(object.NewSphere(*util.NewPoint3(0, 0, -1), 0.5))
	world.Add(object.NewSphere(*util.NewPoint3(0, -100.5, -1), 100))

	// Camera
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := util.NewPoint3(0, 0, 0)

	viewportU := util.NewVec3(viewportWidth, 0, 0)
	viewportV := util.NewVec3(0, -viewportHeight, 0)

	pixelDeltaU := viewportU.Divide(float64(imageWidth))
	pixelDeltaV := viewportV.Divide(float64(imageHeight))

	viewportUpperLeft := cameraCenter.Substract(*util.NewVec3(0, 0, focalLength)).Substract(*viewportU.Divide(2)).Substract(*viewportV.Divide(2))
	pixel00Location := viewportUpperLeft.Add(*pixelDeltaU.Add(*pixelDeltaV).MultiplyBy(0.5))

	// Render
	fmt.Println("P3")
	fmt.Printf("%d %d\n", imageWidth, imageHeight)
	fmt.Println(255)

	for j := range imageHeight {
		stdErr.Printf("Scanline remaining: %d", imageHeight-j)

		for i := range imageWidth {
			pixelCenter := pixel00Location.Add(*pixelDeltaU.MultiplyBy(float64(i))).Add(*pixelDeltaV.MultiplyBy(float64(j)))
			rayDirection := pixelCenter.Substract(*cameraCenter)
			ray := core.NewRay(*cameraCenter, *rayDirection)

			pixelColor := rayColor(ray, world)
			util.WriteColorTo(pixelColor, stdOut)
		}
	}

	stdErr.Println("Done")
}
