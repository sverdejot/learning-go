package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	
	
	t.Run("circle area", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793 

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
	
}

func TestAreaInterfaces(t *testing.T) {
	areaTest := []struct{
		name	string
		shape	Shape
		want	float64
	}{
		{name: "Rectangle",	shape: Rectangle{10.0, 10.0},	want: 100.0},
		{name: "Circle",	shape: Circle{10.0},		want: 314.1592653589793},
		{name: "Triangle",	shape: Triangle{10.0, 10.0},	want: 50.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt, got, tt.want)
			}
		})
	}
}
