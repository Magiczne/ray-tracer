package object

import (
	"fmt"
	"math"
	"ray-tracer/core"
	"ray-tracer/util"
	"ray-tracer/vector"
)

type Sphere struct {
	centerStart  vector.Point3
	CenterVector vector.Vector3
	Radius       float64
	IsMoving     bool
	material     core.Material // TODO: Pointer?
}

func NewSphere(center vector.Point3, radius float64, material core.Material) *Sphere {
	return &Sphere{
		centerStart:  center,
		CenterVector: *vector.EmptyVec3(), // TODO: nil
		Radius:       radius,
		IsMoving:     false,
		material:     material,
	}
}

func NewMovingSphere(center1 vector.Point3, center2 vector.Point3, radius float64, material core.Material) *Sphere {
	return &Sphere{
		centerStart:  center1,
		CenterVector: *center2.Substract(&center1),
		Radius:       radius,
		IsMoving:     true,
		material:     material,
	}
}

func (s *Sphere) Display() {
	fmt.Printf("Sphere(c=%v, cv=%v, r=%f)", s.centerStart, s.CenterVector, s.Radius)
}

func (s *Sphere) Hit(ray *core.Ray, rayT *util.Interval, hitRecord *core.HitRecord) bool {
	center := s.centerStart
	if s.IsMoving {
		center = *s.Center(ray.Time)
	}

	oc := center.Substract(&ray.Origin)
	a := ray.Direction.LengthSquared()
	h := vector.DotProduct(&ray.Direction, oc)
	c := oc.LengthSquared() - s.Radius*s.Radius
	discriminant := h*h - a*c

	if discriminant < 0 {
		return false
	}

	discriminantSqrt := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (h - discriminantSqrt) / a

	if !rayT.Surrounds(root) {
		root = (h + discriminantSqrt) / a

		if !rayT.Surrounds(root) {
			return false
		}
	}

	hitRecord.T = root
	hitRecord.Point = *ray.At(hitRecord.T)

	outwardNormal := hitRecord.Point.Substract(&center).Divide(s.Radius)
	hitRecord.SetFaceNormal(ray, outwardNormal)
	hitRecord.Material = s.material

	return true
}

func (s *Sphere) Center(time float64) *vector.Point3 {
	return s.centerStart.Add(s.CenterVector.MultiplyBy(time))
}
