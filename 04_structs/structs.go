package structs

import "math"


type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width	float64
	Height 	float64
}

type Triangle struct {
	Width	float64
	Height	float64
}

type Circle struct {
	Radius	float64
}

func (t Triangle) Area() float64 {
	return (t.Width * t.Height) / 2
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height) 
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
