package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func main() {
	rect := Rectangle{
		width:  5,
		height: 10,
	}
	circ := Circle{
		radius: 3,
	}
	fmt.Println("Rectanlge Perimeter:", Shape.Perimeter(&rect))
	fmt.Println("Circle Permiter:", Shape.Perimeter(&circ))
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
