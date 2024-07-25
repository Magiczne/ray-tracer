package perlin

import (
	"math"
	"math/rand"
	"ray-tracer/vector"
)

const pointCount = 256

type Perlin struct {
	randomVectors []*vector.Vector3
	permutationsX []int
	permutationsY []int
	permutationsZ []int
}

func NewPerlin() *Perlin {
	randomVectors := make([]*vector.Vector3, pointCount)
	for i := 0; i < pointCount; i++ {
		randomVectors[i] = vector.UnitVector(vector.RandomVector3(-1, 1))
	}

	return &Perlin{
		randomVectors: randomVectors,
		permutationsX: generatePerlinPermutation(),
		permutationsY: generatePerlinPermutation(),
		permutationsZ: generatePerlinPermutation(),
	}
}

func (p *Perlin) Noise(point *vector.Point3) float64 {
	u := point.X - math.Floor(point.X)
	v := point.Y - math.Floor(point.Y)
	w := point.Z - math.Floor(point.Z)

	i := int(math.Floor(point.X))
	j := int(math.Floor(point.Y))
	k := int(math.Floor(point.Z))
	c := [2][2][2]*vector.Vector3{}

	for di := range 2 {
		for dj := range 2 {
			for dk := range 2 {
				c[di][dj][dk] = p.randomVectors[p.permutationsX[(i+di)&255]^p.permutationsY[(j+dj)&255]^p.permutationsZ[(k+dk)&255]]
			}
		}
	}

	return perlinInterpolation(c, u, v, w)
}

func (p *Perlin) Turbulence(point *vector.Point3, depth int) float64 {
	sum := 0.0
	tempPoint := point
	weight := 1.0

	for _ = range depth {
		sum += weight * p.Noise(tempPoint)
		weight *= 0.5
		tempPoint = tempPoint.MultiplyBy(2)
	}

	return math.Abs(sum)
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

func perlinInterpolation(c [2][2][2]*vector.Vector3, u, v, w float64) float64 {
	uu := u * u * (3 - 2*u)
	vv := v * v * (3 - 2*v)
	ww := w * w * (3 - 2*w)
	sum := 0.0

	for i := range 2 {
		for j := range 2 {
			for k := range 2 {
				weightVector := vector.NewVector3(u-float64(i), v-float64(j), w-float64(k))

				sum += (float64(i)*uu + float64(1-i)*(1-uu)) *
					(float64(j)*vv + float64(1-j)*(1-vv)) *
					(float64(k)*ww + float64(1-k)*(1-ww)) *
					vector.DotProduct(c[i][j][k], weightVector)
			}
		}
	}

	return sum
}
