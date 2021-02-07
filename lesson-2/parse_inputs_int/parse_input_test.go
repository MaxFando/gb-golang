package parse_inputs_int

import "testing"

func TestParse(t *testing.T) {
	want := "В числе 234 - 2 сотен, 3 десятков, 4 единиц"
	if got, _ := Parse(234); got != want {
		t.Errorf("Parse(234) = %s, want %s", got, want)
	}
}
