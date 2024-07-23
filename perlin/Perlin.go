package perlin

import (
	"math"
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
	u := point.X() - math.Floor(point.X())
	v := point.Y() - math.Floor(point.Y())
	w := point.Z() - math.Floor(point.Z())

	// Hermitian smoothing
	u = u * u * (3 - 2*u)
	v = v * v * (3 - 2*v)
	w = w * w * (3 - 2*w)

	i := int(math.Floor(point.X()))
	j := int(math.Floor(point.Y()))
	k := int(math.Floor(point.Z()))
	c := [2][2][2]float64{}

	for di := range 2 {
		for dj := range 2 {
			for dk := range 2 {
				c[di][dj][dk] = p.randomFloat[p.permutationsX[(i+di)&255]^p.permutationsY[(j+dj)&255]^p.permutationsZ[(k+dk)&255]]
			}
		}
	}

	return trilinearInterpolation(c, u, v, w)
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

func trilinearInterpolation(c [2][2][2]float64, u, v, w float64) float64 {
	sum := 0.0

	for i := range 2 {
		for j := range 2 {
			for k := range 2 {
				sum += (float64(i)*u + float64(1-i)*(1-u)) * (float64(j)*v + float64(1-j)*(1-v)) * (float64(k)*w + float64(1-k)*(1-w)) * c[i][j][k]
			}
		}
	}

	return sum
}
