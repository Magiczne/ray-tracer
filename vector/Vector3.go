package vector

import (
	"fmt"
	"math"
	"ray-tracer/constants"
)

type Vector3 struct {
	X, Y, Z float64
}

func EmptyVec3() *Vector3 {
	return &Vector3{}
}

func NewVector3(x float64, y float64, z float64) *Vector3 {
	return &Vector3{x, y, z}
}

func (vec3 *Vector3) CopyFrom(other *Vector3) {
	vec3.X = other.X
	vec3.Y = other.Y
	vec3.Z = other.Z
}

func (vec3 *Vector3) Axis(axis constants.Axis) float64 {
	if axis == constants.AxisX {
		return vec3.X
	}

	if axis == constants.AxisY {
		return vec3.Y
	}

	return vec3.Z
}

func (vec3 *Vector3) Length() float64 {
	return math.Sqrt(vec3.LengthSquared())
}

func (vec3 *Vector3) LengthSquared() float64 {
	return vec3.X*vec3.X + vec3.Y*vec3.Y + vec3.Z*vec3.Z
}

func (vec3 *Vector3) String() string {
	return fmt.Sprintf("%f, %f, %f", vec3.X, vec3.Y, vec3.Z)
}

// Vector utility functions
func (vec3 *Vector3) Add(other *Vector3) *Vector3 {
	return NewVector3(vec3.X+other.X, vec3.Y+other.Y, vec3.Z+other.Z)
}

func (vec3 *Vector3) AddInPlace(other *Vector3) {
	vec3.X += other.X
	vec3.Y += other.Y
	vec3.Z += other.Z
}

func (vec3 *Vector3) Substract(other *Vector3) *Vector3 {
	return NewVector3(vec3.X-other.X, vec3.Y-other.Y, vec3.Z-other.Z)
}

func (vec3 *Vector3) Multiply(other *Vector3) *Vector3 {
	return NewVector3(vec3.X*other.X, vec3.Y*other.Y, vec3.Z*other.Z)
}

func (vec3 *Vector3) MultiplyBy(multiplier float64) *Vector3 {
	return NewVector3(vec3.X*multiplier, vec3.Y*multiplier, vec3.Z*multiplier)
}

func (vec3 *Vector3) Divide(divisor float64) *Vector3 {
	return vec3.MultiplyBy(1.0 / divisor)
}

func (vec3 *Vector3) Negate() *Vector3 {
	return NewVector3(-vec3.X, -vec3.Y, -vec3.Z)
}

func (vec3 *Vector3) NearZero() bool {
	epsilon := 1e-8

	return math.Abs(vec3.X) < epsilon && math.Abs(vec3.Y) < epsilon && math.Abs(vec3.Z) < epsilon
}
