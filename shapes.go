package main

import "math"

// Rectangle has the dimensions of a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

// Area returns the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle represents a circle...
type Circle struct {
	Radius float64
}

// Area returns the area of a circle.
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
