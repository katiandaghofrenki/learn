package learn

import (
    "math"
    "strings"
)

// Check if a float64 is an integer
func IsIntegers(f float64) bool {
    return f == math.Floor(f)
}

// Convert a string to a 2D slice of strings
func Str2Slice(input string) [][]string {
    lines := strings.Split(input, "\n")
    var newLines [][]string

    for _, line := range lines {
		var newLine []string
        if len(line) > 0 {
            parts := strings.Split(line, ", ")
            if len(parts) < 2 {
                continue
            }
            lhs := parts[0]
            rhs := parts[1]
            part := strings.Split(lhs, ": ")
            if len(part) < 2 {
                continue
            }
            lh := part[1]
            newLine = append(newLine, lh+rhs)
			newLines = append(newLines, newLine)
        } else {
			newLine = append(newLine, "\n")
			newLines = append(newLines, newLine)
		}
    }
    return newLines
}

// Convert a 2D slice of strings back to a single string
func NewString(newLines [][]string) string {
    var newString []string
    for _, slice := range newLines {
        newString = append(newString, slice...)
    }
    return strings.Join(newString, " ")
}

// Perform calculations on the input neString
func Token(s string) int {
	newLines := Str2Slice(s)
	newString := NewString(newLines)
    costA := 0.00
    costB := 0.00
    costTokenA := 3.00
    costTokenB := 1.00
    total := 0.00
    lines := strings.Split(newString, "\n")

    for i, line := range lines {
        parts := strings.Fields(line)
        if len(parts) < 3 {
            continue
        }
        parts1 := strings.Split(string(parts[0]), "Y+")
        parts2 := strings.Split(string(parts[1]), "Y+")
        parts3 := strings.Split(string(parts[2]), "Y=")

        if len(parts1) < 2 || len(parts2) < 2 || len(parts3) < 2 {
            continue
        }

        lhsparts1 := strings.Split(string(parts1[0]), "X+")
        lhsparts2 := strings.Split(string(parts2[0]), "X+")
        lhsparts3 := strings.Split(string(parts3[0]), "X=")

        if len(lhsparts1) < 2 || len(lhsparts2) < 2 || len(lhsparts3) < 2 {
            continue
        }

        xA := float64(ToInt(lhsparts1[1])) // func ToInt is on helps.go
        yA := float64(ToInt(parts1[1]))
        xB := float64(ToInt(lhsparts2[1]))
        yB := float64(ToInt(parts2[1]))
        priceX := float64(ToInt(lhsparts3[1]))
        priceY := float64(ToInt(parts3[1]))

        timeA := ((priceY*xB) - (priceX*yB)) / ((yA*xB) - (xA*yB))
        timeB := ((priceY*xA) - (priceX*yA)) / ((yB*xA) - (xB*yA))


        if IsIntegers(timeA) && IsIntegers(timeB) {
			priceXresult := (timeA*xA) + (timeB*xB)
			priceYresult := (timeA*yA) + (timeB*yB)
			if  priceX == priceXresult && priceY == priceYresult {
				costA = costTokenA * timeA
				costB = costTokenB * timeB
				total += (costA + costB)
				i++
			}
        }
    }
    return int(total)
}

// usage --> func Main()
// input := `copy your input to here`
// fmt.Println("result part1: \n", learn.Token(input))