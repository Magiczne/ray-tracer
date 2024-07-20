package core

import (
	"ray-tracer/vector"
)

type Ray struct {
	Origin    vector.Point3
	Direction vector.Vector3
}

func NewRay(origin vector.Point3, direction vector.Vector3) *Ray {
	return &Ray{origin, direction}
}

func EmptyRay() *Ray {
	return &Ray{
		Origin:    *vector.NewPoint3(0, 0, 0),
		Direction: *vector.EmptyVec3(),
	}
}

func (r *Ray) CopyFrom(other *Ray) {
	r.Origin = other.Origin
	r.Direction = other.Direction
}

func (ray *Ray) At(t float64) *vector.Point3 {
	return ray.Origin.Add(ray.Direction.MultiplyBy(t))
}
