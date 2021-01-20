package rectangle

import "errors"

// Area Calculate area of a rectangle
func Area(width, length float64) (float64, error) {
	if width == 0 {
		return 0, errors.New("Ширина = 0")
	}

	if length == 0 {
		return 0, errors.New("Длина = 0")
	}

	return width * length, nil
}
