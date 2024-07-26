package object

import (
	"math"
	"ray-tracer/core"
	"ray-tracer/util"
	"ray-tracer/vector"
)

type Ellipse struct {
	center      *vector.Point3
	sideA       *vector.Vector3
	sideB       *vector.Vector3
	w           *vector.Vector3
	normal      *vector.Vector3
	D           float64
	material    core.Material
	boundingBox *core.AABB
}

func NewEllipse(center *vector.Point3, sideA, sideB *vector.Vector3, material core.Material) *Ellipse {
	normalVector := vector.CrossProduct(sideA, sideB)
	normal := vector.UnitVector(normalVector)
	D := vector.DotProduct(normal, center)

	boundingBox := core.NewAABBFromPoints(
		center.Substract(sideA).Substract(sideB),
		center.Add(sideA).Add(sideB),
	)

	return &Ellipse{
		center:      center,
		sideA:       sideA,
		sideB:       sideB,
		w:           normalVector.Divide(vector.DotProduct(normalVector, normalVector)),
		normal:      normal,
		D:           D,
		material:    material,
		boundingBox: boundingBox,
	}
}

func (e *Ellipse) BoundingBox() *core.AABB {
	return e.boundingBox
}

func (e *Ellipse) Hit(ray *core.Ray, rayTime *util.Interval) *core.HitRecord {
	denominator := vector.DotProduct(e.normal, ray.Direction)

	// Ray parallel to the plane
	if math.Abs(denominator) < 1e-8 {
		return nil
	}

	// Hit point parameter t is outside ray interval
	t := (e.D - vector.DotProduct(e.normal, ray.Origin)) / denominator
	if !rayTime.Contains(t) {
		return nil
	}

	hitRecord := core.EmptyHitRecord()

	// Determine the hit point lies within the planar shape using its plane coordinated
	intersection := ray.At(t)
	planarHitpointVector := intersection.Substract(e.center)
	alpha := vector.DotProduct(e.w, vector.CrossProduct(planarHitpointVector, e.sideB))
	beta := vector.DotProduct(e.w, vector.CrossProduct(e.sideA, planarHitpointVector))

	if !e.isInterior(alpha, beta) {
		return nil
	}

	hitRecord.Time = t
	hitRecord.Point = intersection
	hitRecord.Material = e.material
	hitRecord.U = alpha/2 + 0.5
	hitRecord.V = beta/2 + 0.5
	hitRecord.SetFaceNormal(ray, e.normal)

	return hitRecord
}

func (e *Ellipse) isInterior(a, b float64) bool {
	return a*a+b*b <= 1
}
