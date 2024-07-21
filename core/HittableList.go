package core

import (
	"fmt"
	"ray-tracer/util"
)

type HittableList struct {
	Hittable
	boundingBox *AABB
	objects     []Hittable
}

func NewHittableList() *HittableList {
	return &HittableList{
		boundingBox: EmptyAABB(),
		objects:     make([]Hittable, 0),
	}
}

func (h *HittableList) Clear() {
	// TODO: clear bounding box when clearing?
	h.objects = make([]Hittable, 0)
}

func (h *HittableList) Add(hittable Hittable) {
	h.objects = append(h.objects, hittable)
	h.boundingBox = NewAABBFromAABB(h.boundingBox, hittable.BoundingBox())
}

func (h *HittableList) BoundingBox() *AABB {
	return h.boundingBox
}

func (h *HittableList) Display() {
	fmt.Printf("HittableList(objects=%d)", len(h.objects))

	for _, object := range h.objects {
		object.Display()
	}
}

func (h *HittableList) Hit(ray *Ray, rayT *util.Interval, hitRecord *HitRecord) bool {
	tempRecord := NewHitRecord()
	hitAnything := false
	closestSoFar := rayT.Max

	for _, object := range h.objects {
		if object.Hit(ray, util.NewInterval(rayT.Min, closestSoFar), tempRecord) {
			hitAnything = true
			closestSoFar = tempRecord.T
			hitRecord.CopyFrom(*tempRecord)
		}
	}

	return hitAnything
}
