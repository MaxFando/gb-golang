package parse_inputs_int

import (
	"fmt"
	"testing"
)

func ExampleParse() {
	message, _ := Parse(234)
	fmt.Println(message)
	// Output: В числе 234 - 2 сотен, 3 десятков, 4 единиц
}

func TestParse(t *testing.T) {
	tests := []struct {
		num  uint32
		want string
	}{
		{num: 234, want: "В числе 234 - 2 сотен, 3 десятков, 4 единиц"},
		{num: 582, want: "В числе 582 - 5 сотен, 8 десятков, 2 единиц"},
		{num: 102, want: "В числе 102 - 1 сотен, 0 десятков, 2 единиц"},
	}

	for _, tc := range tests {
		if got, _ := Parse(tc.num); got != tc.want {
			t.Errorf("Parse(234) = %s, want %s", got, tc.want)
		}
	}
}
