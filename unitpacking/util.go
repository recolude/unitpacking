package unitpacking

import (
	"math"

	"github.com/EliCDavis/vector"
)

func floorVec2(v vector.Vector2) vector.Vector2 {
	return vector.NewVector2(
		math.Floor(v.X()),
		math.Floor(v.Y()),
	)
}

func clampVec2(v vector.Vector2, min float64, max float64) vector.Vector2 {
	return vector.NewVector2(
		Clamp(v.X(), min, max),
		Clamp(v.Y(), min, max),
	)
}

// Clamp clamps the given value between the given minimum float and maximum
// float values. Returns the given value if it is within the min and max range.
func Clamp(num float64, min float64, max float64) float64 {
	if num < min {
		return min
	}
	if num > max {
		return max
	}
	return num
}
