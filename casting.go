package learn

// import "math"

func Casting(n float64) int {
//	result := int(math.Round(n))
//	return result
	a := int(n)
	if n >= 0 {
		a = int(n + 0.5)
	} else {
		a = int(n - 0.5)
	}
	return a
}