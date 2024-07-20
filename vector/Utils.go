package vector

import "math"

func DotProduct(v1, v2 *Vector3) float64 {
	return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
}

func CrossProduct(v1, v2 *Vector3) *Vector3 {
	return NewVector3(
		v1.y*v2.z-v1.z*v2.y,
		v1.z*v2.x-v1.x*v2.z,
		v1.x*v2.y-v1.y*v2.x,
	)
}

func UnitVector(v *Vector3) *Vector3 {
	return v.Divide(v.Length())
}

func Reflect(v, n *Vector3) *Vector3 {
	return v.Substract(n.MultiplyBy(2 * DotProduct(v, n)))
}

func Refract(uv, n *Vector3, etaiOverEtat float64) *Vector3 {
	cosTheta := math.Min(DotProduct(uv.MultiplyBy(-1), n), 1.0)
	rOutPerpendicular := uv.Add(n.MultiplyBy(cosTheta)).MultiplyBy(etaiOverEtat)
	rOutParallel := n.MultiplyBy(-math.Sqrt(math.Abs(1 - rOutPerpendicular.LengthSquared())))

	return rOutPerpendicular.Add(rOutParallel)
}
