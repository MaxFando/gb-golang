package rectangle

import "errors"

// Calculate area of a rectangle
func Area(width float64, length float64) (float64, error) {
	if width == 0 {
		return 0, errors.New("Width is 0")
	}

	if length == 0 {
		return 0, errors.New("Length is 0")
	}

	return width * length, nil
}
