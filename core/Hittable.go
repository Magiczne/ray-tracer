package core

type Hittable interface {
	Hit(ray *Ray, rayTMin float64, rayTmax float64, hitRecord *HitRecord) bool
}
