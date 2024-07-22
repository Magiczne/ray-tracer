package util

import (
	"math"
)

func LinearToGamma(linearComponent float64) float64 {
	if linearComponent > 0 {
		return math.Sqrt(linearComponent)
	}

	return 0
}

func DegToRad(deg float64) float64 {
	return deg * math.Pi / 180
}

func Reflectance(cos float64, refractionIndex float64) float64 {
	// Use Schlick's approximation for reflectance.
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 = r0 * r0

	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
