package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcPerimeter() float64 {
	result := 2 * math.Pi * c.Radius
	return result
}

// CalcArea returns calculation result of area
func (c Circle) CalcArea() float64 {
	result := math.Pow(c.Radius, 2) * math.Pi
	return result
}
