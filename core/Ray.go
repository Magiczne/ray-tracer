package core

import (
	"ray-tracer/util"
)

type Ray struct {
	origin    util.Point3
	direction util.Vec3
}

func NewRay(origin util.Point3, direction util.Vec3) *Ray {
	return &Ray{origin, direction}
}

func (ray Ray) At(t float64) *util.Point3 {
	return ray.origin.Add(*ray.direction.MultiplyBy(t))
}

func (ray Ray) Origin() util.Point3 {
	return ray.origin
}

func (ray Ray) Direction() util.Vec3 {
	return ray.direction
}
