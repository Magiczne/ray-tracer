package perlin

import (
	"math/rand"
	"ray-tracer/vector"
)

const pointCount = 256

type Perlin struct {
	randomFloat   []float64
	permutationsX []int
	permutationsY []int
	permutationsZ []int
}

func NewPerlin() *Perlin {
	randomFloat := make([]float64, pointCount)
	for i := 0; i < pointCount; i++ {
		randomFloat[i] = rand.Float64()
	}

	return &Perlin{
		randomFloat:   randomFloat,
		permutationsX: generatePerlinPermutation(),
		permutationsY: generatePerlinPermutation(),
		permutationsZ: generatePerlinPermutation(),
	}
}

func (p *Perlin) Noise(point *vector.Point3) float64 {
	i := int(4*point.X()) & 255
	j := int(4*point.Y()) & 255
	k := int(4*point.Z()) & 255

	return p.randomFloat[p.permutationsX[i]^p.permutationsY[j]^p.permutationsZ[k]]
}

func generatePerlinPermutation() []int {
	p := make([]int, pointCount)

	for i := 0; i < pointCount; i++ {
		p[i] = i
	}

	// Perform permutation
	for i := pointCount - 1; i > 0; i-- {
		target := rand.Int31n(int32(i))

		p[i], p[target] = p[target], p[i]
	}

	return p
}
