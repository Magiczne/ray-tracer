package core

type HittableList struct {
	Hittable
	objects []Hittable
}

func (h HittableList) Clear() {
	h.objects = make([]Hittable, 1, 1)
}

func (h HittableList) Add(hittable Hittable) {
	h.objects = append(h.objects, hittable)
}

func (h HittableList) Hit(ray *Ray, rayTMin float64, rayTmax float64, hitRecord *HitRecord) bool {
	tempRecord := NewHitRecord()
	hitAnything := false
	closestSoFar := rayTmax

	for _, object := range h.objects {
		if object.Hit(ray, rayTMin, closestSoFar, tempRecord) {
			hitAnything = true
			closestSoFar = tempRecord.T()
			hitRecord = tempRecord
		}
	}

	return hitAnything
}
