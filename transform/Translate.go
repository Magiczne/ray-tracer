package transform

import (
	"ray-tracer/core"
	"ray-tracer/util"
	"ray-tracer/vector"
)

type Translate struct {
	object      core.Hittable
	offset      *vector.Vector3
	boundingBox *core.AABB
}

func NewTranslate(object core.Hittable, offset *vector.Vector3) *Translate {
	return &Translate{
		object:      object,
		offset:      offset,
		boundingBox: object.BoundingBox().Add(offset),
	}
}

func (t *Translate) BoundingBox() *core.AABB {
	return t.boundingBox
}

func (t *Translate) Hit(ray *core.Ray, rayTime *util.Interval) *core.HitRecord {
	offsetRay := core.NewTimedRay(ray.Origin.Substract(t.offset), ray.Direction, ray.Time)
	hit := t.object.Hit(offsetRay, rayTime)

	if hit == nil {
		return nil
	}

	hit.Point.AddInPlace(t.offset)

	return hit
}
