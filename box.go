package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	totalItems := len(b.shapes)
	if totalItems >= b.shapesCapacity {
		return errors.New("Maximum elements is reached")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= b.shapesCapacity {
		return nil, errors.New("Shape not found")
	}

	return b.shapes[i], nil

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= b.shapesCapacity {
		return nil, errors.New("Out of range error")
	}
	shape := b.shapes[i]
	copy(b.shapes[i:], b.shapes[i+1:])
	b.shapes = b.shapes[:len(b.shapes)-1]
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {

	if i >= len(b.shapes) {
		return nil, errors.New("Out of range error")
	}
	removedShape := b.shapes[i]
	b.shapes[i] = shape
	return removedShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var result float64
	for _, value := range b.shapes {
		result += value.CalcPerimeter()
	}
	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var result float64
	for _, value := range b.shapes {
		result += value.CalcArea()
	}
	return result

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	removed := 0
	var final []Shape
	for _, value := range b.shapes {
		typof := fmt.Sprintf("%T", value)
		switch typof {
		case "*golang_united_school_homework.Circle":

			removed++
		default:
			final = append(final, value)
			continue
		}
	}
	b.shapes = b.shapes[:0]
	b.shapes = append(b.shapes, final...)

	if removed == 0 {
		return errors.New("No circles found")
	} else {
		return nil
	}
}
