package util

import "math"

type Interval struct {
	Min float64
	Max float64
}

func EmptyInterval() *Interval {
	return &Interval{
		Min: math.Inf(1),
		Max: math.Inf(-1),
	}
}

func UniverseInterval() *Interval {
	return &Interval{
		Min: math.Inf(-1),
		Max: math.Inf(1),
	}
}

func NewInterval(min float64, max float64) *Interval {
	return &Interval{
		Min: min,
		Max: max,
	}
}

func NewIntervalFromIntervals(a, b *Interval) *Interval {
	var min float64
	var max float64

	if a.Min <= b.Min {
		min = a.Min
	} else {
		min = b.Min
	}

	if a.Max >= b.Max {
		max = a.Max
	} else {
		max = b.Max
	}

	return &Interval{
		Min: min,
		Max: max,
	}
}

func (i *Interval) Size() float64 {
	return i.Max - i.Min
}

func (i *Interval) Contains(x float64) bool {
	return x >= i.Min && x <= i.Max
}

func (i *Interval) Surrounds(x float64) bool {
	return x > i.Min && x < i.Max
}

func (i *Interval) Clamp(x float64) float64 {
	if x < i.Min {
		return i.Min
	}

	if x > i.Max {
		return i.Max
	}

	return x
}

func (i *Interval) Expand(delta float64) *Interval {
	padding := delta / 2

	return NewInterval(i.Min-padding, i.Max+padding)
}
