package util

import (
	"math"
)

type Vec3 struct {
	e [3]float64
}

func EmptyVec3() *Vec3 {
	return &Vec3{[3]float64{0, 0, 0}}
}

func NewVec3(e0 float64, e1 float64, e2 float64) *Vec3 {
	return &Vec3{[3]float64{e0, e1, e2}}
}

func (vec3 Vec3) X() float64 {
	return vec3.e[0]
}

func (vec3 Vec3) Y() float64 {
	return vec3.e[1]
}

func (vec3 Vec3) Z() float64 {
	return vec3.e[2]
}

func (vec3 Vec3) Length() float64 {
	return math.Sqrt(vec3.LengthSquared())
}

func (vec3 Vec3) LengthSquared() float64 {
	return vec3.e[0]*vec3.e[0] + vec3.e[1]*vec3.e[1] + vec3.e[2]*vec3.e[2]
}

func (vec3 Vec3) Print() {
	print("%d %d %d", vec3.e[0], vec3.e[1], vec3.e[2])
}

// Point3 alias
type Point3 = Vec3

func NewPoint3(e0 float64, e1 float64, e2 float64) *Vec3 {
	return &Point3{[3]float64{e0, e1, e2}}
}

// Vector utility functions
func (vec3 Vec3) Add(other Vec3) *Vec3 {
	return NewVec3(vec3.e[0]+other.e[0], vec3.e[1]+other.e[1], vec3.e[2]+other.e[2])
}

func (vec3 Vec3) Substract(other Vec3) *Vec3 {
	return NewVec3(vec3.e[0]-other.e[0], vec3.e[1]-other.e[1], vec3.e[2]-other.e[2])
}

func (vec3 Vec3) Multiply(other Vec3) *Vec3 {
	return NewVec3(vec3.e[0]*other.e[0], vec3.e[1]*other.e[1], vec3.e[2]*other.e[2])
}

func (vec3 Vec3) MultiplyBy(multiplier float64) *Vec3 {
	return NewVec3(vec3.e[0]*multiplier, vec3.e[1]*multiplier, vec3.e[2]*multiplier)
}

func (vec3 Vec3) Divide(divisor float64) *Vec3 {
	return vec3.MultiplyBy(1 / divisor)
}

func (vec3 Vec3) DotProduct(other Vec3) float64 {
	return vec3.e[0]*other.e[0] + vec3.e[1]*other.e[1] + vec3.e[2]*other.e[2]
}

func (vec3 Vec3) CrossProduct(other Vec3) *Vec3 {
	return NewVec3(
		vec3.e[1]*other.e[2]-vec3.e[2]*other.e[1],
		vec3.e[2]*other.e[0]-vec3.e[0]*other.e[2],
		vec3.e[0]*other.e[1]-vec3.e[1]*other.e[0],
	)
}

func (vec3 Vec3) UnitVector() *Vec3 {
	return vec3.Divide(vec3.Length())
}
