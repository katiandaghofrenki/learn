// advent of code 2024 day 1 part 1
package learn

import (
	"strings"
)

func SideBySide(s string) (int, int) {
	// split string s new line
	lines := strings.Split(s, "\n")

	var sL []int
	var sR []int

	for _, line := range lines {
		parts := strings.Fields(line) // split by spaces
		lhs, rhs := 0, 0
		for i := 0; i < len(parts[0]) && i < len(parts[1]); i++ {
			lhs = lhs*10 + int(parts[0][i] - '0')
			rhs = rhs*10 + int(parts[1][i] - '0')
		}
		sL = append(sL, lhs)
		sR = append(sR, rhs)
	}
	
	for i := 0; i < len(sL)-1 && i < len(sR)-1; i++ { // to sort the sL and sR
		for j := i + 1; j < len(sL) && j < len(sR); j++ {
			if sL[i] > sL[j] {
				sL[i], sL[j] = sL[j], sL[i]
			}
			if sR[i] > sR[j] {
				sR[i], sR[j] = sR[j], sR[i]
			}
		}
	}
//  return sL, sR // this to confirm that the sL and sR has been sorted out properly, must change the output to be ([]int, []int)
	
	num := 0 // part1

	for i := 0; i < len(sL) && i < len(sR); i++ {
		if sL[i] > sR[i] {
			num += (sL[i] - sR[i])
		} else {
			num += (sR[i] - sL[i])
		}
	}
	
	total := 0 // part2
	sum := 0
	match := false
	for i := 0; i < len(sL); i++ {
		for j := 0; j < len(sR); j++ {
			if sL[i] == sR[j] {
				match = true
				if match {
					sum++
					total += sL[i]*sum
				}
				sum = 0
			}
		}
	}
	return num, total
}