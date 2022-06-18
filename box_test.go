package golang_united_school_homework

import (
	"math"
	"testing"
)

func TestAddShapeReturnsError(t *testing.T) {
	tr := Triangle{
		Side: 11,
	}
	boxes := NewBox(1)
	err := boxes.AddShape(tr)
	tr1 := Triangle{
		Side: 11,
	}
	err = boxes.AddShape(tr1)
	if err == nil {
		t.Error("error should be returned")
	}
}

func TestAddShapeReturnsNoError(t *testing.T) {
	tr := Triangle{
		Side: 11,
	}
	boxes := NewBox(1)
	err := boxes.AddShape(tr)
	if err != nil {
		t.Error("error should not be returned")
	}
}

func TestGetByIndexReturnsError(t *testing.T) {
	circle := Circle{
		Radius: 11,
	}
	boxes := NewBox(1)
	_ = boxes.AddShape(circle)
	_, err := boxes.GetByIndex(2)
	if err == nil {
		t.Error("Out of range error should be returned")
	}
}

func TestGetByIndexReturnsNoError(t *testing.T) {
	circle := Circle{
		Radius: 11,
	}
	boxes := NewBox(1)
	_ = boxes.AddShape(circle)
	b, err := boxes.GetByIndex(0)
	if err != nil {
		t.Error("Out of range error should not be returned")
	}
	if b == nil {
		t.Error("A shape should be returned")
	}
}

func TestExtractByIndexReturnsError(t *testing.T) {
	tr := Triangle{
		Side: 11,
	}
	boxes := NewBox(1)
	_ = boxes.AddShape(tr)
	_, err := boxes.ExtractByIndex(11)
	if err == nil {
		t.Error("Out of range error should be returned")
	}
}

func TestExtractByIndexReturnsNoError(t *testing.T) {
	tr := Triangle{
		Side: 11,
	}
	boxes := NewBox(1)
	_ = boxes.AddShape(tr)
	_, err := boxes.ExtractByIndex(0)
	if err != nil {
		t.Error("Out of range error should be returned")
	}
}

func TestReplaceByIndexReturnsError(t *testing.T) {
	tr := Triangle{
		Side: 11,
	}
	boxes := NewBox(1)
	_ = boxes.AddShape(tr)
	circle := Circle{
		Radius: 11,
	}
	_, err := boxes.ReplaceByIndex(11, circle)
	if err == nil {
		t.Error("Out of range error should be returned")
	}
}

func TestReplaceByIndexReturnsNoError(t *testing.T) {
	tr := Triangle{
		Side: 11,
	}
	boxes := NewBox(1)
	_ = boxes.AddShape(tr)
	circle := Circle{
		Radius: 11,
	}
	_, err := boxes.ReplaceByIndex(0, circle)
	if err != nil {
		t.Error("Out of range error should not be returned")
	}
}

func TestSumArea(t *testing.T) {
	circle := Circle{
		Radius: 11,
	}
	boxes := NewBox(1)
	_ = boxes.AddShape(circle)
	result := boxes.SumArea()
	if math.Round(result) != 380 {
		t.Error("Area is wrong")
	}
}

func TestSumPerimeter(t *testing.T) {
	circle := Circle{
		Radius: 11,
	}
	boxes := NewBox(1)
	_ = boxes.AddShape(circle)
	result := boxes.SumPerimeter()

	if math.Round(result) != 69 {
		t.Error("Perimeter is wrong")
	}
}

func TestRemoveCirclesErrors(t *testing.T) {
	tr := Triangle{
		Side: 11,
	}
	boxes := NewBox(1)
	_ = boxes.AddShape(tr)
	err := boxes.RemoveAllCircles()

	if err == nil {
		t.Error("Circles not found")
	}
}
