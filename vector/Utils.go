package vector

import "math"

func DotProduct(v1, v2 *Vector3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func CrossProduct(v1, v2 *Vector3) *Vector3 {
	return NewVector3(
		v1.Y*v2.Z-v1.Z*v2.Y,
		v1.Z*v2.X-v1.X*v2.Z,
		v1.X*v2.Y-v1.Y*v2.X,
	)
}

func UnitVector(v *Vector3) *Vector3 {
	return v.Divide(v.Length())
}

func Reflect(v, n *Vector3) *Vector3 {
	return v.Substract(n.MultiplyBy(2 * DotProduct(v, n)))
}

func Refract(uv, n *Vector3, etaiOverEtat float64) *Vector3 {
	cosTheta := math.Min(DotProduct(uv.Negate(), n), 1.0)
	rOutPerpendicular := uv.Add(n.MultiplyBy(cosTheta)).MultiplyBy(etaiOverEtat)
	rOutParallel := n.MultiplyBy(-math.Sqrt(math.Abs(1 - rOutPerpendicular.LengthSquared())))

	return rOutPerpendicular.Add(rOutParallel)
}
