package main

import "ray-tracer/scene"

func main() {
	switch 7 {
	case 1:
		scene.MaterialTester()
		break

	case 2:
		scene.BouncingSpheres()
		break

	case 3:
		scene.CheckeredSpheres()
		break

	case 4:
		scene.Earth()
		break

	case 5:
		scene.PerlinSpheres()
		break

	case 6:
		scene.Quads()
		break

	case 7:
		scene.SimpleLight()
		break
	}
}
