package object

import (
	"fmt"
	"math"
	"ray-tracer/core"
	"ray-tracer/util"
	"ray-tracer/vector"
)

type Sphere struct {
	Center   vector.Point3
	Radius   float64
	material core.Material // TODO: Pointer?
}

func NewSphere(center vector.Point3, radius float64, material core.Material) *Sphere {
	return &Sphere{center, radius, material}
}

func (s Sphere) Display() {
	fmt.Printf("Sphere(c=%v r=%f)", s.Center, s.Radius)
}

func (s Sphere) Hit(ray *core.Ray, rayT *util.Interval, hitRecord *core.HitRecord) bool {
	oc := s.Center.Substract(&ray.Origin)
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

	outwardNormal := hitRecord.Point.Substract(&s.Center).Divide(s.Radius)
	hitRecord.SetFaceNormal(ray, outwardNormal)
	hitRecord.Material = s.material

	return true
}
