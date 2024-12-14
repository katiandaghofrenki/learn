package learn

func ToInt(s string) int {
    num := 0
    for i := 0; i < len(s); i++ {
        num = num*10 + int(s[i]-'0') // Convert character to int and add to num
    }
    return num
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