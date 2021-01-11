package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0

	if got != want {
		t.Errorf("expected %.2f but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		hasArea  float64
	}{
		{shape: Rectangle{Width:12, Height: 6}, hasArea: 72.0},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v expected %.2f but got %.2f", tt.shape, tt.hasArea, got)
		}
	}
}
