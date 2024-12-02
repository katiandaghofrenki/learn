package learn

import (
	"strings"
)

func IsSafe(s string) int {
    lines := strings.Split(s, "\n")
    var datas [][]int

    // Convert string input to 2D slice of integers
    for _, line := range lines {
        report := strings.Fields(line)
        var data []int
        for _, item := range report {
            tointeg := 0
            for _, char := range item {
                tointeg = tointeg*10 + int(char-'0')
            }
            data = append(data, tointeg)
        }
        datas = append(datas, data)
    }

    countSequences := 0

    // Iterate through each line of data
    for _, level := range datas {
        if len(level) < 2 {
            continue
        }

        // Check if the sequence is increasing or decreasing
        isIncreasing := level[1] > level[0]
        isSequence := true

        for i := 0; i < len(level)-1; i++ {
            diff := level[i+1] - level[i]
            if diff < -3 || diff > 3 || diff == 0 || (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
                isSequence = false
                break
            }
        }

        if isSequence {
            countSequences++
        }
    }

    return countSequences
}