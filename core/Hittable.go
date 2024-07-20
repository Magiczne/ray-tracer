package core

import (
	"ray-tracer/util"
)

type Hittable interface {
	Display()
	Hit(ray *Ray, rayT *util.Interval, hitRecord *HitRecord) bool
}
