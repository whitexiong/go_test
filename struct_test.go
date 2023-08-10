package main

import "testing"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func TestPerimeter(t *testing.T) {
	t.Run("结构体的测试", func(t *testing.T) {
		rec := Rectangle{10.0, 10.0}
		got := Rectangle.Area(rec)
		want := 40.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func Perimeter(rec Rectangle) float64 {
	return (rec.Width + rec.Width) * 2
}

// 方法
func (r Rectangle) Area() float64 {
	return (r.Height + r.Width) * 2
}

func (c Circle) Area() float64 {
	return 0
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}
