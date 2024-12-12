package learn

import (
	"strings"
)

func ParseInput1(input string) ([]int, [][]string) {
    lines := strings.Split(input, "\n")
    var lhsValues []int
    var rhsValues [][]string

    for _, line := range lines {
        parts := strings.Split(line, ": ")
        lhs := ToInt(parts[0])
        rhs := strings.Fields(parts[1])
        lhsValues = append(lhsValues, lhs)
        rhsValues = append(rhsValues, rhs)
    }
    return lhsValues, rhsValues
}