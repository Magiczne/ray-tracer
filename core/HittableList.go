package core

import (
	"ray-tracer/util"
)

type HittableList struct {
	boundingBox *AABB
	objects     []Hittable
}

func EmptyHittableList() *HittableList {
	return &HittableList{
		boundingBox: EmptyAABB(),
		objects:     make([]Hittable, 0),
	}
}

func (h *HittableList) Clear() {
	h.boundingBox = EmptyAABB()
	h.objects = make([]Hittable, 0)
}

func (h *HittableList) Add(hittable Hittable) {
	h.objects = append(h.objects, hittable)
	h.boundingBox = NewAABBFromAABB(h.boundingBox, hittable.BoundingBox())
}

func (h *HittableList) BoundingBox() *AABB {
	return h.boundingBox
}

func (h *HittableList) Objects() []Hittable {
	return h.objects
}

func (h *HittableList) Hit(ray *Ray, rayTime *util.Interval, hitRecord *HitRecord) bool {
	tempRecord := NewHitRecord()
	hitAnything := false
	closestSoFar := rayTime.Max

	for _, object := range h.objects {
		if object.Hit(ray, util.NewInterval(rayTime.Min, closestSoFar), tempRecord) {
			hitAnything = true
			closestSoFar = tempRecord.T
			hitRecord.CopyFrom(tempRecord)
		}
	}

	return hitAnything
}
