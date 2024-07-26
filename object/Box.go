package object

import (
	"math"
	"ray-tracer/core"
	"ray-tracer/vector"
)

func NewBox(a, b *vector.Point3, material core.Material) *core.HittableList {
	box := core.EmptyHittableList()

	// Construct the two opposite vertices with the minimum and maximum coordinates.
	min := vector.NewPoint3(math.Min(a.X, b.X), math.Min(a.Y, b.Y), math.Min(a.Z, b.Z))
	max := vector.NewPoint3(math.Max(a.X, b.X), math.Max(a.Y, b.Y), math.Max(a.Z, b.Z))

	dx := vector.NewVector3(max.X-min.X, 0, 0)
	dy := vector.NewVector3(0, max.Y-min.Y, 0)
	dz := vector.NewVector3(0, 0, max.Z-min.Z)

	box.Add(NewQuad(vector.NewPoint3(min.X, min.Y, max.Z), dx, dy, material))                // front
	box.Add(NewQuad(vector.NewPoint3(max.X, min.Y, max.Z), dz.MultiplyBy(-1), dy, material)) // right
	box.Add(NewQuad(vector.NewPoint3(max.X, min.Y, min.Z), dx.MultiplyBy(-1), dy, material)) // back
	box.Add(NewQuad(vector.NewPoint3(min.X, min.Y, min.Z), dz, dy, material))                // left
	box.Add(NewQuad(vector.NewPoint3(min.X, max.Y, max.Z), dx, dz.MultiplyBy(-1), material)) // top
	box.Add(NewQuad(vector.NewPoint3(min.X, min.Y, min.Z), dx, dz, material))                // bottom

	return box
}
