package learn

import (
    "strings"
    "fmt"
    "strconv"
)

func Calibrations(input string) int {
    lhsValues, rhsValues := ParseInput(input) // function on separate file
    total := CalculateTotal(lhsValues, rhsValues) // function on separate file
    return total
}

func ParseInput(input string) ([]int, [][]int) {
    lines := strings.Split(input, "\n")
    var lhsValues []int
    var rhsValues [][]int

    for _, line := range lines {
        parts := strings.Split(line, ": ")
        lhs := ToInt(parts[0])
        rhs := strings.Fields(parts[1])
        var rhsNums []int
        for _, num := range rhs {
            rhsNums = append(rhsNums, ToInt(num))
        }
        lhsValues = append(lhsValues, lhs)
        rhsValues = append(rhsValues, rhsNums)
    }
    return lhsValues, rhsValues
}

func CalculateTotal(lhsValues []int, rhsValues [][]int) int {
    total := 0

    for i := 0; i < len(lhsValues); i++ {
        lhs := lhsValues[i]
        rhs := rhsValues[i]
        combinations := GenerateCombinations(rhs)

        matchFound := false
        for _, combination := range combinations {
            results := ApplyOperators(combination)
            for _, result := range results {
                if lhs == result {
                    total += lhs
                    matchFound = true
                    break
                }
            }
            if matchFound {
                break
            }
        }
    }
    return total
}

func ApplyOperators(nums []int) []int {
    var results []int
    if len(nums) == 0 {
        return results
    }
    var helper func(int, int)
    helper = func(index int, currentValue int) {
        if index == len(nums) {
            results = append(results, currentValue)
            return
        }
        helper(index+1, currentValue+nums[index])
        helper(index+1, currentValue*nums[index])
    }
    helper(0, 0)
    return results
}

func GenerateCombinations(arr []int) [][]int {
    var combinations [][]int
    var helper func(int, []int)
    helper = func(start int, current []int) {
        temp := make([]int, len(current))
        copy(temp, current)
        combinations = append(combinations, temp)
        for i := start; i < len(arr); i++ {
            helper(i+1, append(current, arr[i]))
        }
    }
    helper(0, []int{})
    return combinations
}

func CalculateTotal2(input string) int {
    lines := strings.Split(input, "\n")
    total := 0

    for _, line := range lines {
        parts := strings.Split(line, ": ")
        target, _ := strconv.Atoi(parts[0])
        numStrs := strings.Split(parts[1], " ")
        var nums []int
        for _, str := range numStrs {
            num, _ := strconv.Atoi(str)
            nums = append(nums, num)
        }

        operators := []string{"+", "*", "|"}
        combinations := GenerateCombinations2(len(nums) - 1)

        for _, combo := range combinations {
            if TestCombination(nums, combo, operators) == target {
                total += target
                break
            }
        }
    }

    return total
}

func GenerateCombinations2(n int) [][]int {
    if n == 0 {
        return [][]int{}
    }

    combinations := [][]int{}
    current := make([]int, n)
    for {
        copyCombo := make([]int, n)
        copy(copyCombo, current)
        combinations = append(combinations, copyCombo)

        i := 0
        for ; i < n; i++ {
            if current[i] < 2 {
                current[i]++
                break
            }
            current[i] = 0
        }
        if i == n {
            break
        }
    }
    return combinations
}

func TestCombination(nums []int, combo []int, operators []string) int {
    result := nums[0]
    for i := 1; i < len(nums); i++ {
        switch operators[combo[i-1]] {
        case "+":
            result += nums[i]
        case "*":
            result *= nums[i]
        case "|":
            combinedValue, _ := strconv.Atoi(fmt.Sprintf("%d%d", result, nums[i]))
            result = combinedValue
        }
    }
    return result
}

/*
usage : on your func main
input := ``
input2 := ``
fmt.Println(learn.Calibrations(input))
fmt.Println(learn.Calibrations(input2))
fmt.Println(learn.CalculateTotal2(input))
fmt.Println(learn.CalculateTotal2(input2))
*/