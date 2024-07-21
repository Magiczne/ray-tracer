package core

import (
	"ray-tracer/vector"
)

type Ray struct {
	Origin    vector.Point3
	Direction vector.Vector3
	Time      float64
}

func NewRay(origin vector.Point3, direction vector.Vector3) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: direction,
		Time:      0,
	}
}

func NewTimedRay(origin vector.Point3, direction vector.Vector3, time float64) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: direction,
		Time:      time,
	}
}

func EmptyRay() *Ray {
	return &Ray{
		Origin:    *vector.NewPoint3(0, 0, 0),
		Direction: *vector.EmptyVec3(),
		Time:      0,
	}
}

func (r *Ray) CopyFrom(other *Ray) {
	r.Origin = other.Origin
	r.Direction = other.Direction
}

func (ray *Ray) At(t float64) *vector.Point3 {
	return ray.Origin.Add(ray.Direction.MultiplyBy(t))
}
