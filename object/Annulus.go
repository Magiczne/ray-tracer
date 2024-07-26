package object

import (
	"math"
	"ray-tracer/core"
	"ray-tracer/util"
	"ray-tracer/vector"
)

type Annulus struct {
	center      *vector.Point3
	sideA       *vector.Vector3
	sideB       *vector.Vector3
	inner       float64
	w           *vector.Vector3
	normal      *vector.Vector3
	D           float64
	material    core.Material
	boundingBox *core.AABB
}

func NewAnnulus(center *vector.Point3, sideA, sideB *vector.Vector3, inner float64, material core.Material) *Annulus {
	normalVector := vector.CrossProduct(sideA, sideB)
	normal := vector.UnitVector(normalVector)
	D := vector.DotProduct(normal, center)

	boundingBox := core.NewAABBFromPoints(
		center.Substract(sideA).Substract(sideB),
		center.Add(sideA).Add(sideB),
	)

	return &Annulus{
		center:      center,
		sideA:       sideA,
		sideB:       sideB,
		inner:       inner,
		w:           normalVector.Divide(vector.DotProduct(normalVector, normalVector)),
		normal:      normal,
		D:           D,
		material:    material,
		boundingBox: boundingBox,
	}
}

func (a *Annulus) BoundingBox() *core.AABB {
	return a.boundingBox
}

func (a *Annulus) Hit(ray *core.Ray, rayTime *util.Interval) *core.HitRecord {
	denominator := vector.DotProduct(a.normal, ray.Direction)

	// Ray parallel to the plane
	if math.Abs(denominator) < 1e-8 {
		return nil
	}

	// Hit point parameter t is outside ray interval
	t := (a.D - vector.DotProduct(a.normal, ray.Origin)) / denominator
	if !rayTime.Contains(t) {
		return nil
	}

	hitRecord := core.EmptyHitRecord()

	// Determine the hit point lies within the planar shape using its plane coordinated
	intersection := ray.At(t)
	planarHitpointVector := intersection.Substract(a.center)
	alpha := vector.DotProduct(a.w, vector.CrossProduct(planarHitpointVector, a.sideB))
	beta := vector.DotProduct(a.w, vector.CrossProduct(a.sideA, planarHitpointVector))

	if !a.isInterior(alpha, beta) {
		return nil
	}

	hitRecord.Time = t
	hitRecord.Point = intersection
	hitRecord.Material = a.material
	hitRecord.U = alpha/2 + 0.5
	hitRecord.V = beta/2 + 0.5
	hitRecord.SetFaceNormal(ray, a.normal)

	return hitRecord
}

func (ann *Annulus) isInterior(a, b float64) bool {
	centerDist := math.Sqrt(a*a + b*b)

	return centerDist >= ann.inner && centerDist <= 1
}
