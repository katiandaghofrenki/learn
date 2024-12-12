package learn

import (
	"strings"
)

// Repeats2 processes the input string 's' for 'n' iterations and returns the total number of resulting elements
func Blinks(s string, n int) int {
    // Split the input string into slices (words)
    slices := strings.Fields(s)
    // Map to keep track of the counts of each unique slice
    counts := make(map[string]int)
    // Initialize the counts map with the initial slices
    for _, slice := range slices {
        counts[slice]++
    }
    
    // Perform 'n' iterations
    for i := 0; i < n; i++ {
        // Map to store the counts of the new slices after processing
        newCounts := make(map[string]int)
        // Process each slice based on its length and content
        for slice, count := range counts {
            if len(slice)%2 == 0 {
                // If the slice length is even, split it into two halves
                half := len(slice) / 2
                num := ToInt(slice[half:]) // Convert the second half to an integer
                half2 := IntToString(num)  // Convert the integer back to a string
                // Update the newCounts map with the two halves
                newCounts[slice[:half]] += count
                newCounts[half2] += count
            } else if slice == "0" {
                // If the slice is "0", change it to "1"
                newCounts["1"] += count
            } else {
                // For other slices, multiply the number by 2024
                mul := ToInt(slice) * 2024 // Convert to integer and multiply
                newCounts[IntToString(mul)] += count // Convert back to string and update the count
            }
        }
        // Update counts for the next iteration
        counts = newCounts
    }

    // Calculate the total number of slices after all iterations
    total := 0
    for _, count := range counts {
        total += count
    }

    // Convert the total to int64 and return
    return total
}

func IntToString(num int) string {
	if num == 0 {
		return "0"
	}
	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}

	var result []byte

	for num > 0 {
		digit := num % 10
		result = append([]byte{byte(digit + '0')}, result...)
		num /= 10
	}

	if isNegative {
		result = append([]byte{'-'}, result...)
	}
	return string(result)
}