package rectangle

import "testing"

func TestArea(t *testing.T) {
	want := 20.0
	if got, _ := Area(10.0, 2.0); got != want {
		t.Errorf("Area() = %f, want %f", got, want)
	}
}
