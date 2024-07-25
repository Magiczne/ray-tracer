package object

import (
	"math"
	"ray-tracer/vector"
)

func GetSphereUV(point *vector.Point3) (float64, float64) {
	theta := math.Acos(-point.Y)
	phi := math.Atan2(-point.Z, point.X) + math.Pi
	u := phi / (2 * math.Pi)
	v := theta / math.Pi

	return u, v
}
