package core

import (
	"ray-tracer/util"
)

type Hittable interface {
	BoundingBox() *AABB
	Hit(ray *Ray, rayT *util.Interval, hitRecord *HitRecord) bool
}
