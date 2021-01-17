package unitpacking

func clamp(num float64, min float64, max float64) float64 {
	if num < min {
		return min
	}

	if num > max {
		return max
	}

	return num
}
