package learn

import (
    "strings"
)

func Printing(s1 string, s2 string) int {
    lines1 := strings.Split(s1, "\n")  // Split s1 by newlines into lines1
    lines2 := strings.Split(s2, "\n")  // Split s2 by newlines into lines2
    sequenceMap := make(map[int][]int) // Create a map to hold sequences

    // Parse s1 to create a sequence map
    for _, line := range lines1 {
        parts := strings.Split(line, "|")  // Split each line by "|"
        lhs := ToInt(parts[0])             // Convert left-hand side to integer
        rhs := ToInt(parts[1])             // Convert right-hand side to integer
        sequenceMap[lhs] = append(sequenceMap[lhs], rhs) // Append rhs to lhs in sequenceMap
    }

    // Iterate through lines in s2 and check for valid sequences
    sumMiddleNumbers := 0
    for _, line := range lines2 {
        parts := strings.Split(line, ",")        // Split each line by ","
        sequence := make([]int, len(parts))      // Create a slice of ints with the same length as parts
        for i, part := range parts {             // Iterate over parts
            sequence[i] = ToInt(part)           // Convert each part to int and store in sequence
        }

        if IsValidSequence(sequence, sequenceMap) { // Check if the sequence is valid
            middleIndex := len(sequence) / 2        // Find the middle index
            middleNumber := sequence[middleIndex]   // Get the middle number
            sumMiddleNumbers += middleNumber        // Add the middle number to the sum
        }
    }

    return sumMiddleNumbers
}

func ToInt2(s string) int {
    num := 0
    for i := 0; i < len(s); i++ {
        num = num*10 + int(s[i]-'0') // Convert character to int and add to num
    }
    return num
}

func IsValidSequence(sequence []int, sequenceMap map[int][]int) bool {
    for i := 0; i < len(sequence)-1; i++ {
        lhs := sequence[i]  // Get the current number
        rhs := sequence[i+1] // Get the next number
        valid := false
        if nextNums, found := sequenceMap[lhs]; found { // Check if lhs has a valid next number
            for _, next := range nextNums {
                if next == rhs { // If rhs is a valid next number
                    valid = true
                    break
                }
            }
        }
        if !valid { // If no valid next number is found, return false
            return false
        }
    }
    return true // All numbers are in valid sequence
}

func Printing2(s1 string, s2 string) int {
    lines1 := strings.Split(s1, "\n")
    lines2 := strings.Split(s2, "\n")
    precedenceMap := make(map[int][]int)


    // Parse s1 to create a precedence map
    for _, line := range lines1 {
        parts := strings.Split(line, "|")
        lhs := ToInt(parts[0])
        rhs := ToInt(parts[1])
        precedenceMap[lhs] = append(precedenceMap[lhs], rhs)
    }


    // Sum of middle numbers from corrected sequences
    sumMiddleNumbers := 0

    for _, line := range lines2 {
        parts := strings.Split(line, ",")
        sequence := make([]int, len(parts))
        for i, part := range parts {
            sequence[i] = ToInt(part)
        }

        if !IsValidSequence(sequence, precedenceMap) {
            sequence = FixSequence(sequence, precedenceMap)
            middleIndex := len(sequence) / 2
            middleNumber := sequence[middleIndex]
            sumMiddleNumbers += middleNumber
        }
    }

    return sumMiddleNumbers
}

func Contains(slice []int, value int) bool {
    for _, v := range slice {
        if v == value {
            return true
        }
    }
    return false
}

func FixSequence(sequence []int,  precedenceMap map[int][]int) []int {
    n := len(sequence)
        fixed := make([]int, n)
        copy(fixed, sequence)

        for i := 0; i < n-1; i++ {
            for j := i + 1; j < n; j++ {
                // If sequence[i] should come after sequence[j] based on precedenceMap
                if Contains(precedenceMap[fixed[j]], fixed[i]) {
                    // Swap the elements
                    fixed[i], fixed[j] = fixed[j], fixed[i]
                }
            }
        }
        return fixed
}