package core

import (
	"ray-tracer/constants"
	"ray-tracer/util"
	"ray-tracer/vector"
)

type AABB struct {
	X, Y, Z *util.Interval
}

func EmptyAABB() *AABB {
	return &AABB{
		X: util.EmptyInterval(),
		Y: util.EmptyInterval(),
		Z: util.EmptyInterval(),
	}
}

func NewAABBFromIntervals(x, y, z *util.Interval) *AABB {
	return &AABB{
		X: x,
		Y: y,
		Z: z,
	}
}

// Treat the two points a and b as extrema for the bounding box, so we don't require a
// particular minimum/maximum coordinate order.
func NewAABBFromPoints(a, b *vector.Point3) *AABB {
	var x, y, z *util.Interval

	if a.X() <= b.X() {
		x = util.NewInterval(a.X(), b.X())
	} else {
		x = util.NewInterval(b.X(), a.X())
	}

	if a.Y() <= b.Y() {
		y = util.NewInterval(a.Y(), b.Y())
	} else {
		y = util.NewInterval(b.Y(), a.Y())
	}

	if a.Z() <= b.Z() {
		z = util.NewInterval(a.Z(), b.Z())
	} else {
		z = util.NewInterval(b.Z(), a.Z())
	}

	return &AABB{
		X: x,
		Y: y,
		Z: z,
	}
}

func (aabb *AABB) AxisInterval(axis constants.Axis) *util.Interval {
	if axis == constants.AxisY {
		return aabb.Y
	}

	if axis == constants.AxisZ {
		return aabb.Z
	}

	return aabb.X
}

func (aabb *AABB) Hit(ray *Ray, rayT *util.Interval) bool {
	rayOrigin := ray.Origin
	rayDirection := ray.Direction

	for axis := range 3 {
		axisInterval := aabb.AxisInterval(constants.Axis(axis))
		adinv := 1.0 / rayDirection.Axis(constants.Axis(axis))

		t0 := (axisInterval.Min - rayOrigin.Axis(constants.Axis(axis))) * adinv
		t1 := (axisInterval.Max - rayOrigin.Axis(constants.Axis(axis))) * adinv

		if t0 < t1 {
			if t0 > rayT.Min {
				rayT.Min = t0
			}

			if t1 < rayT.Max {
				rayT.Max = t1
			}
		} else {
			if t1 > rayT.Min {
				rayT.Min = t1
			}

			if t0 < rayT.Max {
				rayT.Max = t0
			}
		}

		if rayT.Max <= rayT.Min {
			return false
		}
	}

	return true
}
