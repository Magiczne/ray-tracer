package main

import "ray-tracer/scene"

func main() {
	scenes := make(map[int]func())
	scenes[1] = scene.MaterialTester
	scenes[2] = scene.BouncingSpheres
	scenes[3] = scene.CheckeredSpheres
	scenes[4] = scene.Earth
	scenes[5] = scene.PerlinSpheres
	scenes[6] = scene.Quads
	scenes[7] = scene.SimpleLight
	scenes[8] = scene.CornellBox
	scenes[9] = scene.SmokedCornellBox
	scenes[10] = scene.CreateNextWeekFinalScene(800, 10000, 40)
	scenes[11] = scene.CreateNextWeekFinalScene(400, 250, 4)

	scenes[11]()
}
