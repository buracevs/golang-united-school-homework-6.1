package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (r Rectangle) CalcPerimeter() float64 {
	result := r.Height*2 + r.Weight*2
	return result
}

func (r Rectangle) CalcArea() float64 {
	result := r.Height * r.Weight
	return result
}
