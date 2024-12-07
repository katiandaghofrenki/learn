package learn

func Calibration(input string) int {
    lhsValues, rhsValues := ParseInput(input) // function on separate file
    total := CalculateTotal(lhsValues, rhsValues) // function on separate file
    return total
}