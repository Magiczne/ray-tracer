package vector

import (
	"fmt"
	"math"
	"ray-tracer/constants"
)

type Vector3 struct {
	x, y, z float64
}

func EmptyVec3() *Vector3 {
	return &Vector3{}
}

func NewVector3(x float64, y float64, z float64) *Vector3 {
	return &Vector3{x, y, z}
}

func (vec3 *Vector3) CopyFrom(other *Vector3) {
	vec3.x = other.x
	vec3.y = other.y
	vec3.z = other.z
}

func (vec3 *Vector3) X() float64 {
	return vec3.x
}

func (vec3 *Vector3) Y() float64 {
	return vec3.y
}

func (vec3 *Vector3) Z() float64 {
	return vec3.z
}

func (vec3 *Vector3) Axis(axis constants.Axis) float64 {
	if axis == constants.AxisY {
		return vec3.y
	}

	if axis == constants.AxisZ {
		return vec3.z
	}

	return vec3.x
}

func (vec3 *Vector3) Length() float64 {
	return math.Sqrt(vec3.LengthSquared())
}

func (vec3 *Vector3) LengthSquared() float64 {
	return vec3.x*vec3.x + vec3.y*vec3.y + vec3.z*vec3.z
}

func (vec3 *Vector3) String() string {
	return fmt.Sprintf("%f, %f, %f", vec3.x, vec3.y, vec3.z)
}

// Vector utility functions
func (vec3 *Vector3) Add(other *Vector3) *Vector3 {
	return NewVector3(vec3.x+other.x, vec3.y+other.y, vec3.z+other.z)
}

func (vec3 *Vector3) AddInPlace(other *Vector3) {
	vec3.x += other.x
	vec3.y += other.y
	vec3.z += other.z
}

func (vec3 *Vector3) Substract(other *Vector3) *Vector3 {
	return NewVector3(vec3.x-other.x, vec3.y-other.y, vec3.z-other.z)
}

func (vec3 *Vector3) Multiply(other *Vector3) *Vector3 {
	return NewVector3(vec3.x*other.x, vec3.y*other.y, vec3.z*other.z)
}

func (vec3 *Vector3) MultiplyBy(multiplier float64) *Vector3 {
	return NewVector3(vec3.x*multiplier, vec3.y*multiplier, vec3.z*multiplier)
}

func (vec3 *Vector3) Divide(divisor float64) *Vector3 {
	return vec3.MultiplyBy(1 / divisor)
}

func (vec3 *Vector3) NearZero() bool {
	epsilon := 1e-8

	return (math.Abs(vec3.x) < epsilon) && (math.Abs(vec3.y) < epsilon) && (math.Abs(vec3.z) < epsilon)
}
