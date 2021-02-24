package rectangle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArea(t *testing.T) {
	tests := []struct {
		width  float64
		height float64
	}{
		{10, 10},
		{2, 5},
		{3, 12},
	}

	for _, tc := range tests {
		_, err := Area(tc.width, tc.height)
		assert.Nil(t, err)
	}
}
