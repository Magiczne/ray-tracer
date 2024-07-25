package color

import "ray-tracer/vector"

func FromVector3(vec3 *vector.Vector3) *Color {
	return NewColor(vec3.X, vec3.Y, vec3.Z)
}
