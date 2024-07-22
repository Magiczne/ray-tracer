package bvh

import (
	"ray-tracer/core"
	"ray-tracer/util"
	"sort"
)

type BVHNode struct {
	boundingBox *core.AABB
	left        core.Hittable
	right       core.Hittable
}

func NewBVHNode(objects []core.Hittable) *BVHNode {
	boundingBox := core.EmptyAABB()
	for _, object := range objects {
		boundingBox = core.SurroundingAABB(boundingBox, object.BoundingBox())
	}

	axis := boundingBox.LongestAxis()

	sort.Slice(objects, func(i, j int) bool {
		return core.AABBSortByAxis(objects[i].BoundingBox(), objects[j].BoundingBox(), axis)
	})

	var left, right core.Hittable
	if len(objects) == 1 {
		left = objects[0]
		right = objects[0]
	} else if len(objects) == 2 {
		left = objects[0]
		right = objects[1]
	} else {
		left = NewBVHNode(objects[:len(objects)/2])
		right = NewBVHNode(objects[len(objects)/2:])
	}

	return &BVHNode{
		left:        left,
		right:       right,
		boundingBox: core.SurroundingAABB(left.BoundingBox(), right.BoundingBox()),
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
		if hitRecordLeft.Time < hitRecordRight.Time {
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
