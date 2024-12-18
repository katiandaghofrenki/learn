package learn

func ToInt(s string) int {
    num := 0
    sign := 1
    startIndex := 0

    if len(s) > 0 && s[0] == '-' {
        sign = -1
        startIndex = 1
    } else if s[0] == '+' {
        startIndex = 1
    }

    for i := startIndex; i < len(s); i++ {
        if s[i] < '0' || s[i] > '9' {
            return 0
        }
        num = num*10 + int(s[i] - '0')
    }
    return num*sign
}

func Found(grid [][]rune, char rune) (int, int, bool) {
    for i, row := range grid {
        for j, col := range row {
            if col == char {
                return i, j, true
            }
        }
    }
    return -1, -1, false
}

func FoundX(grid [][]rune, char rune) int {
    sum := 0
    for _, row := range grid {
        for _, col := range row {
            if col == char {
                sum++
            }
        }
    }
    return sum
}

func TryMe() string {
    return "you can try what ever you want in here. This has similar function with func main(). Just change the output as this is string so the output is string"
}