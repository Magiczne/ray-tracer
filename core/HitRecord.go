package core

import (
	"ray-tracer/vector"
)

type HitRecord struct {
	Point     *vector.Point3
	Normal    *vector.Vector3
	Material  Material
	Time      float64
	U         float64
	V         float64
	FrontFace bool
}

func EmptyHitRecord() *HitRecord {
	return &HitRecord{
		Point:     vector.NewPoint3(0, 0, 0),
		Normal:    vector.EmptyVec3(),
		Material:  nil,
		Time:      0,
		U:         0,
		V:         0,
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
