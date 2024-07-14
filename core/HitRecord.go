package core

import (
	"ray-tracer/util"
)

type HitRecord struct {
	point     util.Point3
	normal    util.Vec3
	t         float64
	frontFace bool
}

func NewHitRecord() *HitRecord {
	return &HitRecord{
		point:     *util.NewPoint3(0, 0, 0),
		normal:    *util.EmptyVec3(),
		t:         0,
		frontFace: false,
	}
}

func (hitRecord HitRecord) Point() util.Point3 {
	return hitRecord.point
}

func (hitRecord HitRecord) SetPoint(point util.Point3) {
	hitRecord.point = point
}

func (hitRecord HitRecord) Normal() util.Vec3 {
	return hitRecord.normal
}

func (hitRecord HitRecord) SetNormal(normal util.Vec3) {
	hitRecord.normal = normal
}

func (hitRecord HitRecord) T() float64 {
	return hitRecord.t
}

func (hitRecord HitRecord) SetT(t float64) {
	hitRecord.t = t
}

func (hitRecord HitRecord) SetFaceNormal(ray *Ray, outwardNormal *util.Vec3) {
	frontFace := ray.direction.DotProduct(*outwardNormal) < 0

	if frontFace {
		hitRecord.normal = *outwardNormal
	} else {
		hitRecord.normal = *outwardNormal.MultiplyBy(-1)
	}
}

func (hitRecord HitRecord) GetFrontFace() bool {
	return hitRecord.frontFace
}
