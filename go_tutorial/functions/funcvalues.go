package functions

import "math"

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func FuncValues() float64 {
	hypo := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	compute(hypo)
	return compute(hypo)
}
