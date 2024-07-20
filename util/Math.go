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
