package learn

import (
    "strings"
)

func FindXmas(s string, f string) (string, int, string) {
    // Split the input string by lines
    lines := strings.Split(strings.TrimSpace(s), "\n")

    // Convert the lines to a 2D grid of runes
    grid := make([][]rune, len(lines))
    for i, line := range lines {
        grid[i] = []rune(line)
    }

    count := 0
    directions := [][]int{
        {0, 1},  // horizontal right
        {0, -1}, // horizontal left
        {1, 0},  // vertical down
        {-1, 0}, // vertical up
        {1, 1},  // diagonal down-right
        {-1, -1},// diagonal up-left
        {1, -1}, // diagonal down-left
        {-1, 1}, // diagonal up-right
    }

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            for _, dir := range directions {
				search := true
                for k := 0; k < len(f); k++ {
					nx, ny := i+k*dir[0], j+k*dir[1]
					if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) || grid[nx][ny] != rune(f[k]) {
						search = false
					}
				}
				if search {
					count++
				}
				
            }
        }
    }

    return "XMAS Appears : ", count, "times"
}

// =============== Part Two ===================
func FindXMASPattern(s string, f string) (string, int, string) {
    // Split the input string by lines
    lines := strings.Split(strings.TrimSpace(s), "\n")

    // Convert the lines to a 2D grid of runes
    grid := make([][]rune, len(lines))
    for i, line := range lines {
        grid[i] = []rune(line)
    }
    count := 0
    length := len(f)


    // Iterate over each cell in the grid
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if IsWithinBounds(grid, i, j, 1, 1, length) && CheckPattern(grid, f, i, j, 1, 1) {
                if (IsWithinBounds(grid, i, j+length-1, 1, -1, length) && CheckPattern(grid, f, i, j+length-1, 1, -1)) ||
                (IsWithinBounds(grid, i+length-1, j, -1, 1, length) && CheckPattern(grid, f, i+length-1, j, -1, 1)) {
                    count++
                }
            }
            if IsWithinBounds(grid, i, j, 1, -1, length) && CheckPattern(grid, f, i, j, 1, -1) {
                if (IsWithinBounds(grid, i, j-length+1, 1, 1, length) && CheckPattern(grid, f, i, j-length+1, 1, 1)) ||
                (IsWithinBounds(grid, i+length-1, j, -1, -1, length) && CheckPattern(grid, f, i+length-1, j, -1, -1)) {
                    count++
                }
            }
            if IsWithinBounds(grid, i, j, -1, 1, length) && CheckPattern(grid, f, i, j, -1, 1) {
                if (IsWithinBounds(grid, i, j+length-1, -1, -1, length) && CheckPattern(grid, f, i, j+length-1, -1, -1)) ||
                (IsWithinBounds(grid, i-length+1, j, 1, 1, length) && CheckPattern(grid, f, i-length+1, j, 1, 1)) {
                    count++
                }
            }
            if IsWithinBounds(grid, i, j, -1, -1, length) && CheckPattern(grid, f, i, j, -1, -1) {
                if (IsWithinBounds(grid, i, j-length+1, -1, 1, length) && CheckPattern(grid, f, i, j-length+1, -1, 1)) ||
                (IsWithinBounds(grid, i-length+1, j, 1, -1, length) && CheckPattern(grid, f, i-length+1, j, 1, -1)) {
                    count++
                }
            }
        }
    }
    return "X form of MAS (X-MAS) appears : ", count/2, "times" // it must be devided by 2 due to each founding was counted two while it just suppose to be one.
}

func CheckPattern(grid [][]rune, word string, x, y, dx, dy int) bool {
    length := len(word)
    for k := 0; k < length; k++ {
        nx, ny := x+k*dx, y+k*dy
        if grid[nx][ny] != rune(word[k]) {
            return false
        }
    }
    return true
}

func IsWithinBounds(grid [][]rune, x, y, dx, dy, length int) bool {
    for k := 0; k < length; k++ {
        nx, ny := x+k*dx, y+k*dy
        if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) {
            return false
        }
    }
    return true
}

/*
// usage: 
    input := `copy the input data here`
    fmt.Println(learn.FindXmas(input, "XMAS"))
    fmt.Println(learn.FindXMASPattern(input2, "MAS"))
*/