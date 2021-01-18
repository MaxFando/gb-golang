package circle

import "math"

func Length(area int) float64 {
	return 2 * math.Pi * Diameter(area)
}
