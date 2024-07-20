package vector

type Point3 = Vector3

func NewPoint3(x float64, y float64, z float64) *Vector3 {
	return &Point3{x, y, z}
}
