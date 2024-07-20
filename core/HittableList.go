package core

import (
	"fmt"
	"os"
	"ray-tracer/util"
)

type HittableList struct {
	Hittable
	objects []Hittable
}

func NewHittableList() *HittableList {
	return &HittableList{
		objects: make([]Hittable, 0),
	}
}

func (h *HittableList) Clear() {
	h.objects = make([]Hittable, 0)
}

func (h *HittableList) Add(hittable Hittable) {
	h.objects = append(h.objects, hittable)
}

func (h *HittableList) Display() {
	fmt.Fprintf(os.Stderr, "Hittable list: %d", len(h.objects))

	for _, object := range h.objects {
		object.Display()
	}
}

func (h *HittableList) Hit(ray *Ray, rayT *util.Interval, hitRecord *HitRecord) bool {
	tempRecord := NewHitRecord()
	hitAnything := false
	closestSoFar := rayT.Max()

	for _, object := range h.objects {
		if object.Hit(ray, util.NewInterval(rayT.Min(), closestSoFar), tempRecord) {
			hitAnything = true
			closestSoFar = tempRecord.T
			hitRecord.CopyFrom(*tempRecord)
		}
	}

	return hitAnything
}
