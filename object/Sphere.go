package object

import (
	"math"
	"ray-tracer/core"
	"ray-tracer/util"
)

type Sphere struct {
	center util.Point3
	radius float64
}

func NewSphere(center util.Point3, radius float64) *Sphere {
	return &Sphere{center, radius}
}

func (sphere Sphere) Center() util.Point3 {
	return sphere.center
}

func (shpere Sphere) Radius() float64 {
	return shpere.radius
}

func (sphere Sphere) Hit(ray *core.Ray, rayTMin float64, rayTmax float64, hitRecord *core.HitRecord) bool {
	oc := sphere.Center().Substract(ray.Origin())
	a := ray.Direction().LengthSquared()
	h := ray.Direction().DotProduct(*oc)
	c := oc.LengthSquared() - sphere.radius*sphere.radius
	discriminant := h*h - a*c

	if discriminant < 0 {
		println("Delta < 0")

		return false
	}

	discriminantSqrt := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (h - discriminantSqrt) / a

	if root <= rayTMin || rayTmax <= root {
		root = (h + discriminantSqrt) / a

		if root <= rayTMin || rayTmax <= root {
			return false
		}
	}

	hitRecord.SetT(root)
	hitRecord.SetPoint(*ray.At(hitRecord.T()))
	hitRecord.SetNormal(*hitRecord.Point().Substract(sphere.center).Divide(sphere.radius))

	outwardNormal := hitRecord.Point().Substract(sphere.center).Divide(sphere.Radius())
	hitRecord.SetFaceNormal(ray, outwardNormal)

	return true
}
