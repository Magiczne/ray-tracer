package vector

import "ray-tracer/random"

func RandomVec3(min float64, max float64) *Vector3 {
	return &Vector3{
		x: random.Float64(min, max),
		y: random.Float64(min, max),
		z: random.Float64(min, max),
	}
}

func RandomVec3InUnitSphere() *Vector3 {
	for {
		vec := RandomVec3(-1, 1)

		if vec.LengthSquared() < 1 {
			return vec
		}
	}
}

func RandomUnitVector() *Vector3 {
	return UnitVector(RandomVec3InUnitSphere())
}

func RandomVec3OnHemisphere(normal *Vector3) *Vector3 {
	onUnitSphere := RandomUnitVector()

	if DotProduct(onUnitSphere, normal) > 0.0 {
		return onUnitSphere
	}

	return UnitVector(onUnitSphere).MultiplyBy(-1)
}

func RandomVec3InUnitDisk() *Vector3 {
	for {
		vec := NewVec3(random.Float64(-1, 1), random.Float64(-1, 1), 0)

		if vec.LengthSquared() < 1 {
			return vec
		}
	}
}
