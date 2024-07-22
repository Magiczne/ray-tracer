package bvh

import (
	"math/rand"
	"ray-tracer/constants"
	"ray-tracer/core"
	"ray-tracer/util"
	"slices"
)

type BVHNode struct {
	boundingBox *core.AABB
	left        core.Hittable
	right       core.Hittable
}

func NewBVHNodeFromHittableList(hittableList core.HittableList) *BVHNode {
	return NewBVHNodeFromListOfObjects(hittableList.Objects(), 0, len(hittableList.Objects()))
}

func NewBVHNodeFromListOfObjects(objects []core.Hittable, start int, end int) *BVHNode {
	axis := constants.Axis(rand.Int31n(3))
	objectSpan := end - start

	var left, right core.Hittable
	if objectSpan == 1 {
		left = objects[start]
		right = objects[start]
	} else if objectSpan == 2 {
		left = objects[start]
		right = objects[start+1]
	} else {
		slices.SortFunc(objects[start:end], func(a, b core.Hittable) int {
			aAxisInterval := a.BoundingBox().AxisInterval(axis)
			bAxisInterval := b.BoundingBox().AxisInterval(axis)

			// TODO: Handle equals?
			if aAxisInterval.Min < bAxisInterval.Min {
				return -1
			}

			if aAxisInterval.Min > bAxisInterval.Min {
				return 1
			}

			return 0
		})

		mid := start + objectSpan/2
		left = NewBVHNodeFromListOfObjects(objects, start, mid)
		right = NewBVHNodeFromListOfObjects(objects, mid, end)
	}

	return &BVHNode{
		left:        left,
		right:       right,
		boundingBox: core.NewAABBFromAABB(left.BoundingBox(), right.BoundingBox()),
	}
}

func (node *BVHNode) BoundingBox() *core.AABB {
	return node.boundingBox
}

func (node *BVHNode) Hit(ray *core.Ray, rayTime *util.Interval) *core.HitRecord {
	if !node.boundingBox.Hit(ray, rayTime) {
		return nil
	}

	hitRecordLeft := node.left.Hit(ray, rayTime)
	hitRecordRight := node.right.Hit(ray, rayTime)

	if hitRecordLeft != nil && hitRecordRight != nil {
		if hitRecordLeft.T < hitRecordRight.T {
			return hitRecordLeft
		}

		return hitRecordRight
	}

	if hitRecordLeft != nil {
		return hitRecordLeft
	}

	if hitRecordRight != nil {
		return hitRecordRight
	}

	return nil
}
