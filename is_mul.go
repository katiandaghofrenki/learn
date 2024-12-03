package learn

import (
    "regexp"
    "strconv"
)

func IsMul(s string) (int, int) {
    total1 := 0 // for part1
	total2 := 0 // for part2
	doSeen := true
	reMul := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
    reDo := regexp.MustCompile(`do\(\)`)
    reDont := regexp.MustCompile(`don't\(\)`)

    // Find all matches and their submatches
    found := reMul.FindAllStringSubmatch(s, -1)

    for _, match := range found {
        // Extract the numbers from the match
        x, _ := strconv.Atoi(match[1])
        y, _ := strconv.Atoi(match[2])

        // Calculate (x * y) and add to total
        total1 += x * y
    }

	segments := regexp.MustCompile(`do\(\)|mul\(\d{1,3},\d{1,3}\)|don't\(\)`).FindAllString(s, -1)

	for _, segment := range segments {
        if reDont.MatchString(segment) {
            // Reset if 'don't()' is found
            doSeen = false
        } else if reDo.MatchString(segment) {
            // Enable counting after 'do()' is found
            doSeen = true
        } else if doSeen && reMul.MatchString(segment) {
            match := reMul.FindStringSubmatch(segment)
            x, _ := strconv.Atoi(match[1])
            y, _ := strconv.Atoi(match[2])
            total2 += x * y
        }
    }

    return total1, total2
}