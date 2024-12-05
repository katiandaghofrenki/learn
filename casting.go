package learn

// import "math"

func Casting(n float64) int {
//	result := int(math.Round(n))
//	return result
	a := 0
	if n >= 0 {
		a = int(n + 0.4999999999999999)
	} else {
		a = int(n - 0.4999999999999999)
	}
	return a
}