package vector

import "ray-tracer/random"

func RandomVector3(min float64, max float64) *Vector3 {
	return &Vector3{
		x: random.Float64(min, max),
		y: random.Float64(min, max),
		z: random.Float64(min, max),
	}
}

func RandomVector3InUnitSphere() *Vector3 {
	for {
		vec := RandomVector3(-1, 1)

		if vec.LengthSquared() < 1 {
			return vec
		}
	}
}

func RandomUnitVector() *Vector3 {
	return UnitVector(RandomVector3InUnitSphere())
}

func RandomVector3OnHemisphere(normal *Vector3) *Vector3 {
	onUnitSphere := RandomUnitVector()

	if DotProduct(onUnitSphere, normal) > 0.0 {
		return onUnitSphere
	}

	return UnitVector(onUnitSphere).MultiplyBy(-1)
}

func RandomVector3InUnitDisk() *Vector3 {
	for {
		vec := NewVector3(random.Float64(-1, 1), random.Float64(-1, 1), 0)

		if vec.LengthSquared() < 1 {
			return vec
		}
	}
}
