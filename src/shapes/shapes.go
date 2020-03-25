package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() (area float64) {
	return r.Height * r.Width
}

func (c Circle) Area() (area float64) {
	return math.Pi * c.Radius * c.Radius
}
