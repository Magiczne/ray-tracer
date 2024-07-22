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
	h.boundingBox = SurroundingAABB(h.boundingBox, hittable.BoundingBox())
}

func (h *HittableList) BoundingBox() *AABB {
	return h.boundingBox
}

func (h *HittableList) Objects() []Hittable {
	return h.objects
}

func (h *HittableList) Hit(ray *Ray, rayT *util.Interval) *HitRecord {
	var hitRecord *HitRecord
	hitAnything := false
	closestSoFar := rayT.Max

	for _, object := range h.objects {
		tempRecord := object.Hit(ray, util.NewInterval(rayT.Min, closestSoFar))

		if tempRecord != nil {
			hitAnything = true
			hitRecord = tempRecord
			closestSoFar = tempRecord.T
		}
	}

	if hitAnything {
		return hitRecord
	}

	return nil
}
