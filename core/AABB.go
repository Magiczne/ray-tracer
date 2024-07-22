package core

import (
	"math"
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

func SurroundingAABB(box1, box2 *AABB) *AABB {
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

func (aabb *AABB) Hit(ray *Ray, rayTime *util.Interval) bool {
	for axis := range 3 {
		constantAxis := constants.Axis(axis)
		axisInterval := aabb.AxisInterval(constantAxis)
		directionAxisInverted := 1.0 / ray.Direction.Axis(constantAxis)

		t0 := (axisInterval.Min - ray.Origin.Axis(constantAxis)) * directionAxisInverted
		t1 := (axisInterval.Max - ray.Origin.Axis(constantAxis)) * directionAxisInverted

		if directionAxisInverted < 0.0 {
			t0, t1 = t1, t0
		}

		tMin := math.Max(t0, rayTime.Min)
		tMax := math.Min(t1, rayTime.Max)

		if tMax < tMin {
			return false
		}
	}

	return true
}

func (aabb *AABB) LongestAxis() constants.Axis {
	if aabb.X.Size() > aabb.Y.Size() {
		if aabb.X.Size() > aabb.Z.Size() {
			return constants.AxisX
		}

		return constants.AxisZ
	}

	if aabb.Y.Size() > aabb.Z.Size() {
		return constants.AxisY
	}

	return constants.AxisZ
}

func AABBSortByAxis(lhs, rhs *AABB, axis constants.Axis) bool {
	return lhs.AxisInterval(axis).Min < rhs.AxisInterval(axis).Min
}
