package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0

	if got != want {
		t.Errorf("expected %.2f but got %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(Rectangle{10.0, 10.0})
	want := 100.0

	if got != want {
		t.Errorf("expected %.2f but got %.2f", got, want)
	}
}
