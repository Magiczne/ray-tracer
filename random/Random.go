package random

import "math/rand/v2"

func Float64(min float64, max float64) float64 {
	return min + (max-min)*rand.Float64()
}
