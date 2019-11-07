package functions

import (
	"math"
)

type Shapes interface {
	Area() float64
}

type Circles struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (c Circles) Area() float64 {
	return math.Pi * c.Radius*c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func GetArea(s Shapes) float64 {
	return s.Area()
}