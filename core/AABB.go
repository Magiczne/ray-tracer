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

func NewAABBFromAABB(box1, box2 *AABB) *AABB {
	return &AABB{
		X: util.NewIntervalFromIntervals(box1.X, box2.X),
		Y: util.NewIntervalFromIntervals(box1.Y, box2.Y),
		Z: util.NewIntervalFromIntervals(box1.Z, box2.Z),
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
	if axis == constants.AxisX {
		return aabb.X
	}

	if axis == constants.AxisY {
		return aabb.Y
	}

	return aabb.Z
}

func (aabb *AABB) Hit(ray *Ray, rayT *util.Interval) bool {
	for axis := range 3 {
		constantAxis := constants.Axis(axis)
		axisInterval := aabb.AxisInterval(constantAxis)
		directionAxisInverted := 1.0 / ray.Direction.Axis(constantAxis)

		t0 := (axisInterval.Min - ray.Origin.Axis(constantAxis)) * directionAxisInverted
		t1 := (axisInterval.Max - ray.Origin.Axis(constantAxis)) * directionAxisInverted

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
