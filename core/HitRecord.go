package core

import (
	"ray-tracer/vector"
)

type HitRecord struct {
	Point     *vector.Point3
	Normal    *vector.Vector3
	Material  Material
	T         float64
	FrontFace bool
}

func EmptyHitRecord() *HitRecord {
	return &HitRecord{
		Point:     vector.NewPoint3(0, 0, 0),
		Normal:    vector.EmptyVec3(),
		Material:  nil,
		T:         0,
		FrontFace: false,
	}
}

func (hr *HitRecord) SetFaceNormal(ray *Ray, outwardNormal *vector.Vector3) {
	hr.FrontFace = vector.DotProduct(ray.Direction, outwardNormal) < 0

	if hr.FrontFace {
		hr.Normal = outwardNormal
	} else {
		hr.Normal = outwardNormal.MultiplyBy(-1)
	}
}

func (hr *HitRecord) GetFrontFace() bool {
	return hr.FrontFace
}
