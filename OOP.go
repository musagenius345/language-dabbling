package main

import "fmt"

// Define a struct (similar to a class)
type Circle struct {
	radius float64
}

// Define a method for the Circle struct
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// Interface declaration
type Shape interface {
	area() float64
}

// Another struct
type Rectangle struct {
	width  float64
	height float64
}

// Method for Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Embedding: A struct embedding another struct
type ColoredCircle struct {
	Circle
	color string
}

func main() {
	// Creating objects
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 4, height: 6}
	coloredCircle := ColoredCircle{
		Circle: Circle{radius: 3},
		color:  "red",
	}

	// Using methods
	fmt.Printf("Circle Area: %.2f\n", circle.area())
	fmt.Printf("Rectangle Area: %.2f\n", rectangle.area())
	fmt.Printf("Colored Circle Area: %.2f\n", coloredCircle.area())

	// Using interfaces
	shapes := []Shape{circle, rectangle, coloredCircle}
	for _, shape := range shapes {
		fmt.Printf("Shape Area: %.2f\n", shape.area())
	}
}
