package object

import (
	"math"
	"ray-tracer/core"
	"ray-tracer/util"
	"ray-tracer/vector"
)

type Quad struct {
	Q           *vector.Point3
	u           *vector.Vector3
	v           *vector.Vector3
	w           *vector.Vector3
	normal      *vector.Vector3
	D           float64
	material    core.Material
	boundingBox *core.AABB
}

func NewQuad(q *vector.Point3, u, v *vector.Vector3, material core.Material) *Quad {
	normal := vector.UnitVector(vector.CrossProduct(u, v))
	D := vector.DotProduct(normal, q)

	boundingBoxDiagonal1 := core.NewAABBFromPoints(q, q.Add(u).Add(v))
	boundingBoxDiagonal2 := core.NewAABBFromPoints(q.Add(u), q.Add(v))

	return &Quad{
		Q:           q,
		u:           u,
		v:           v,
		w:           normal.Divide(vector.DotProduct(normal, normal)),
		normal:      normal,
		D:           D,
		material:    material,
		boundingBox: core.SurroundingAABB(boundingBoxDiagonal1, boundingBoxDiagonal2),
	}
}

func (q *Quad) BoundingBox() *core.AABB {
	return q.boundingBox
}

func (q *Quad) Hit(ray *core.Ray, rayTime *util.Interval) *core.HitRecord {
	denominator := vector.DotProduct(q.normal, ray.Direction)

	// Ray parallel to the plane
	if math.Abs(denominator) < 1e-8 {
		return nil
	}

	// Hit point parameter t is outside ray interval
	t := (q.D - vector.DotProduct(q.normal, ray.Origin)) / denominator
	if !rayTime.Contains(t) {
		return nil
	}

	hitRecord := core.EmptyHitRecord()

	// Determine the hit point lies within the planar shape using its plane coordinated
	intersection := ray.At(t)
	planarHitpointVector := intersection.Substract(q.Q)
	alpha := vector.DotProduct(q.w, vector.CrossProduct(planarHitpointVector, q.v))
	beta := vector.DotProduct(q.w, vector.CrossProduct(q.u, planarHitpointVector))

	if !q.isInterior(alpha, beta) {
		return nil
	}

	hitRecord.Time = t
	hitRecord.Point = intersection
	hitRecord.Material = q.material
	hitRecord.U = alpha
	hitRecord.V = beta
	hitRecord.SetFaceNormal(ray, q.normal)

	return hitRecord
}

func (q *Quad) isInterior(a, b float64) bool {
	unitInterval := util.NewInterval(0, 1)

	return unitInterval.Contains(a) && unitInterval.Contains(b)
}
