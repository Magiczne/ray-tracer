package core

import (
	"ray-tracer/util"
)

type Hittable interface {
	BoundingBox() *AABB
	Hit(ray *Ray, rayTime *util.Interval) *HitRecord
}
