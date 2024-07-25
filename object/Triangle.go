package object

import (
	"math"
	"ray-tracer/core"
	"ray-tracer/util"
	"ray-tracer/vector"
)

type Triangle struct {
	o           *vector.Point3
	aa          *vector.Vector3
	ab          *vector.Vector3
	w           *vector.Vector3
	normal      *vector.Vector3
	D           float64
	material    core.Material
	boundingBox *core.AABB
}

func NewTriangle(o *vector.Point3, aa, ab *vector.Vector3, material core.Material) *Triangle {
	normalVector := vector.CrossProduct(aa, ab)
	normal := vector.UnitVector(normalVector)
	D := vector.DotProduct(normal, o)

	boundingBoxDiagonal1 := core.NewAABBFromPoints(o, o.Add(aa).Add(ab))
	boundingBoxDiagonal2 := core.NewAABBFromPoints(o.Add(aa), o.Add(ab))

	return &Triangle{
		o:           o,
		aa:          aa,
		ab:          ab,
		w:           normalVector.Divide(vector.DotProduct(normalVector, normalVector)),
		normal:      normal,
		D:           D,
		material:    material,
		boundingBox: core.SurroundingAABB(boundingBoxDiagonal1, boundingBoxDiagonal2),
	}
}

func (t *Triangle) BoundingBox() *core.AABB {
	return t.boundingBox
}

func (triangle *Triangle) Hit(ray *core.Ray, rayTime *util.Interval) *core.HitRecord {
	denominator := vector.DotProduct(triangle.normal, ray.Direction)

	// Ray parallel to the plane
	if math.Abs(denominator) < 1e-8 {
		return nil
	}

	// Hit point parameter t is outside ray interval
	t := (triangle.D - vector.DotProduct(triangle.normal, ray.Origin)) / denominator
	if !rayTime.Contains(t) {
		return nil
	}

	hitRecord := core.EmptyHitRecord()

	// Determine the hit point lies within the planar shape using its plane coordinated
	intersection := ray.At(t)
	planarHitpointVector := intersection.Substract(triangle.o)
	alpha := vector.DotProduct(triangle.w, vector.CrossProduct(planarHitpointVector, triangle.ab))
	beta := vector.DotProduct(triangle.w, vector.CrossProduct(triangle.aa, planarHitpointVector))

	if !triangle.isInterior(alpha, beta) {
		return nil
	}

	hitRecord.Time = t
	hitRecord.Point = intersection
	hitRecord.Material = triangle.material
	hitRecord.U = alpha
	hitRecord.V = beta
	hitRecord.SetFaceNormal(ray, triangle.normal)

	return hitRecord
}

func (triangle *Triangle) isInterior(a, b float64) bool {
	return a >= 0 && b >= 0 && (a+b <= 1)
}
