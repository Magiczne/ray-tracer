package util

import "math"

type Interval struct {
	min float64
	max float64
}

func EmptyInterval() *Interval {
	return &Interval{
		min: math.Inf(1),
		max: math.Inf(-1),
	}
}

func UniverseInterval() *Interval {
	return &Interval{
		min: math.Inf(-1),
		max: math.Inf(1),
	}
}

func NewInterval(min float64, max float64) *Interval {
	return &Interval{min, max}
}

func (i *Interval) Min() float64 {
	return i.min
}

func (i *Interval) Max() float64 {
	return i.max
}

func (i *Interval) Size() float64 {
	return i.max - i.min
}

func (i *Interval) Contains(x float64) bool {
	return i.min <= x && x <= i.max
}

func (i *Interval) Surrounds(x float64) bool {
	return i.min < x && x < i.max
}

func (i *Interval) Clamp(x float64) float64 {
	if x < i.min {
		return i.min
	}

	if x > i.max {
		return i.max
	}

	return x
}
