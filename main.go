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

	scenes[7]()
}
