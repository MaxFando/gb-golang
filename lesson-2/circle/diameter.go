package circle

import "math"

func Diameter(area int) float64 {
	return 2 * math.Sqrt(float64(area)/math.Pi)
}
