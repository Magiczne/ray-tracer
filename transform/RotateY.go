package transform

import (
	"math"
	"ray-tracer/core"
	"ray-tracer/util"
	"ray-tracer/vector"
)

type RotateY struct {
	object      core.Hittable
	sinTheta    float64
	cosTheta    float64
	boundingBox *core.AABB
}

func NewRotateY(object core.Hittable, angle float64) *RotateY {
	theta := util.DegToRad(angle)
	sinTheta := math.Sin(theta)
	cosTheta := math.Cos(theta)
	boundingBox := object.BoundingBox()

	// A little bit different approach than in the tutorial, using bounding box points
	p1 := toWorldSpace(object.BoundingBox().Min(), cosTheta, sinTheta)
	p2 := toWorldSpace(object.BoundingBox().Max(), cosTheta, sinTheta)

	min := vector.NewPoint3(math.Min(p1.X, p2.X), p1.Y, math.Min(p1.Z, p2.Z))
	max := vector.NewPoint3(math.Max(p1.X, p2.X), p2.Y, math.Max(p1.Z, p2.Z))

	boundingBox = core.NewAABBFromPoints(min, max)

	return &RotateY{
		object:      object,
		sinTheta:    sinTheta,
		cosTheta:    cosTheta,
		boundingBox: boundingBox,
	}
}

func (t *RotateY) BoundingBox() *core.AABB {
	return t.boundingBox
}

func (rot *RotateY) Hit(ray *core.Ray, rayTime *util.Interval) *core.HitRecord {
	// Change the ray from world space to object space
	origin := toObjectSpace(ray.Origin, rot.cosTheta, rot.sinTheta)
	direction := toObjectSpace(ray.Direction, rot.cosTheta, rot.sinTheta)
	rotatedRay := core.NewTimedRay(origin, direction, ray.Time)

	// Determine whether an intersection exists in object space (and if so, where)
	hit := rot.object.Hit(rotatedRay, rayTime)
	if hit == nil {
		return nil
	}

	// Change the intersection point and normal from object space to world space
	hit.Point = toWorldSpace(hit.Point, rot.cosTheta, rot.sinTheta)
	hit.Normal = toWorldSpace(hit.Normal, rot.cosTheta, rot.sinTheta)

	return hit
}

func toObjectSpace(v *vector.Vector3, cosTheta, sinTheta float64) *vector.Vector3 {
	return vector.NewVector3(
		cosTheta*v.X-sinTheta*v.Z,
		v.Y,
		sinTheta*v.X+cosTheta*v.Z,
	)
}

func toWorldSpace(v *vector.Vector3, cosTheta, sinTheta float64) *vector.Vector3 {
	return vector.NewVector3(
		cosTheta*v.X+sinTheta*v.Z,
		v.Y,
		-sinTheta*v.X+cosTheta*v.Z,
	)
}
