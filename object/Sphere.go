package object

import (
	"fmt"
	"math"
	"ray-tracer/core"
	"ray-tracer/util"
	"ray-tracer/vector"
)

type Sphere struct {
	centerStart  *vector.Point3
	CenterVector *vector.Vector3
	Radius       float64
	IsMoving     bool
	material     core.Material
	boundingBox  *core.AABB
}

func NewSphere(center *vector.Point3, radius float64, material core.Material) *Sphere {
	radiusVector := vector.NewVector3(radius, radius, radius)

	return &Sphere{
		centerStart:  center,
		CenterVector: nil,
		Radius:       radius,
		IsMoving:     false,
		material:     material,
		boundingBox:  core.NewAABBFromPoints(center.Substract(radiusVector), center.Add(radiusVector)),
	}
}

func NewMovingSphere(center1 *vector.Point3, center2 *vector.Point3, radius float64, material core.Material) *Sphere {
	radiusVector := vector.NewVector3(radius, radius, radius)
	box1 := core.NewAABBFromPoints(center1.Substract(radiusVector), center1.Add(radiusVector))
	box2 := core.NewAABBFromPoints(center2.Substract(radiusVector), center2.Add(radiusVector))

	return &Sphere{
		centerStart:  center1,
		CenterVector: center2.Substract(center1),
		Radius:       radius,
		IsMoving:     true,
		material:     material,
		boundingBox:  core.SurroundingAABB(box1, box2),
	}
}

func (s *Sphere) BoundingBox() *core.AABB {
	return s.boundingBox
}

func (s *Sphere) Display() {
	fmt.Printf("Sphere(c=%v, cv=%v, r=%f)", s.centerStart, s.CenterVector, s.Radius)
}

func (s *Sphere) Hit(ray *core.Ray, rayTime *util.Interval) *core.HitRecord {
	center := s.centerStart
	if s.IsMoving {
		center = s.Center(ray.Time)
	}

	oc := center.Substract(ray.Origin)
	a := ray.Direction.LengthSquared()
	h := vector.DotProduct(ray.Direction, oc)
	c := oc.LengthSquared() - s.Radius*s.Radius
	discriminant := h*h - a*c

	if discriminant < 0 {
		return nil
	}

	discriminantSqrt := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (h - discriminantSqrt) / a

	if !rayTime.Surrounds(root) {
		root = (h + discriminantSqrt) / a

		if !rayTime.Surrounds(root) {
			return nil
		}
	}

	hitRecord := core.EmptyHitRecord()
	hitRecord.Time = root
	hitRecord.Point = ray.At(hitRecord.Time)

	outwardNormal := hitRecord.Point.Substract(center).Divide(s.Radius)
	hitRecord.SetFaceNormal(ray, outwardNormal)
	hitRecord.Material = s.material

	return hitRecord
}

func (s *Sphere) Center(time float64) *vector.Point3 {
	return s.centerStart.Add(s.CenterVector.MultiplyBy(time))
}
