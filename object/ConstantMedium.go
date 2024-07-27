package object

import (
	"math"
	"math/rand/v2"
	"ray-tracer/color"
	"ray-tracer/core"
	"ray-tracer/material"
	"ray-tracer/util"
	"ray-tracer/vector"
)

type ConstantMedium struct {
	boundary               core.Hittable
	negatedInvertedDensity float64
	phaseFunction          core.Material
}

func NewConstantMedium(boundary core.Hittable, density float64, albedo *color.Color) *ConstantMedium {
	return &ConstantMedium{
		boundary:               boundary,
		negatedInvertedDensity: -1 / density,
		phaseFunction:          material.NewIsotropic(albedo),
	}
}

func NewTexturedConstantMedium(boundary core.Hittable, density float64, texture core.Texture) *ConstantMedium {
	return &ConstantMedium{
		boundary:               boundary,
		negatedInvertedDensity: -1 / density,
		phaseFunction:          material.NewTexturedIsotropic(texture),
	}
}

func (cm *ConstantMedium) BoundingBox() *core.AABB {
	return cm.boundary.BoundingBox()
}

func (cm *ConstantMedium) Hit(ray *core.Ray, rayTime *util.Interval) *core.HitRecord {
	hit1 := cm.boundary.Hit(ray, util.UniverseInterval())
	if hit1 == nil {
		return nil
	}

	hit2 := cm.boundary.Hit(ray, util.NewInterval(hit1.Time+0.0001, math.Inf(1)))
	if hit2 == nil {
		return nil
	}

	if hit1.Time < rayTime.Min {
		hit1.Time = rayTime.Min
	}

	if hit2.Time > rayTime.Max {
		hit2.Time = rayTime.Max
	}

	if hit1.Time >= hit2.Time {
		return nil
	}

	if hit1.Time < 0 {
		hit1.Time = 0
	}

	rayLength := ray.Direction.Length()
	distanceInsideBoundary := (hit2.Time - hit1.Time) * rayLength
	hitDistance := cm.negatedInvertedDensity * math.Log(rand.Float64())

	if hitDistance > distanceInsideBoundary {
		return nil
	}

	hit := core.EmptyHitRecord()
	hit.Time = hit1.Time + hitDistance/rayLength
	hit.Point = ray.At(hit.Time)
	hit.Normal = vector.NewVector3(1, 0, 0) // arbitrary value
	hit.FrontFace = true                    // arbitrary value
	hit.Material = cm.phaseFunction

	return hit
}
