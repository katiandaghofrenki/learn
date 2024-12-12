package learn

import (
    "strings"
)

func Historian(s string) ([][]rune, int) {
    lines := strings.Split(s, "\n")
    
    grid := make([][]rune, len(lines))
    for i, line := range lines {
        parts := strings.TrimSpace(line)
        grid[i] = []rune(parts)
    }

    char := '^'
    x, y, found := Found(grid, char)

    for i := 0; i < len(grid)*len(grid[0]); i++ { // Ensure enough iterations
        // Initial movement loop
        for found {
            if x > 0 && grid[x-1][y] != '#' { // Ensure boundary check first
                grid[x][y] = 'X'
                x = x-1 // Move to the new position
                grid[x][y] = char
            } else if x > 0 && grid[x-1][y] == '#' {
                grid[x][y] = '>' // Mark the final position before hitting '#'
                char = '>'
                x, y, found = Found(grid, char) // Re-find position
                break
            } else if x <= 0 { // Check boundary to avoid out-of-bounds access
                grid[x][y] = 'X'
                found = false // Stop the loop
                break
            }
        }
    
        // Continue moving right
        for found {
            if y+1 < len(grid[0]) && grid[x][y+1] != '#' { // Ensure boundary check first
                grid[x][y] = 'X'
                y = y + 1 // Move to the new position
                grid[x][y] = char
            } else if y+1 < len(grid[0]) && grid[x][y+1] == '#' {
                grid[x][y] = 'v' // Mark the final position before hitting '#'
                char = 'v'
                x, y, found = Found(grid, char) // Re-find position
                break
            } else if y+1 >= len(grid[0]) { // Check boundary to avoid out-of-bounds access
                grid[x][y] = 'X'
                found = false // Stop the loop
                break
            }
        }
    
        // Continue moving down
        for found {
            if x+1 < len(grid) && grid[x+1][y] != '#' { // Ensure boundary check first
                grid[x][y] = 'X'
                x = x + 1 // Move to the new position
                grid[x][y] = char
            } else if x+1 < len(grid) && grid[x+1][y] == '#' {
                grid[x][y] = '<' // Mark the final position before hitting '#'
                char = '<'
                x, y, found = Found(grid, char) // Re-find position
                break
            } else if x+1 >= len(grid) { // Check boundary to avoid out-of-bounds access
                grid[x][y] = 'X'
                found = false // Stop the loop
                break
            }
        }
    
        // Continue moving left
        for found {
            if y > 0 && grid[x][y-1] != '#' { // Ensure boundary check first
                grid[x][y] = 'X'
                y = y - 1 // Move to the new position
                grid[x][y] = char
            } else if y > 0 && grid[x][y-1] == '#' {
                grid[x][y] = '^' // Mark the final position before hitting '#'
                char = '^'
                x, y, found = Found(grid, char) // Re-find position
                break
            } else if y <= 0 { // Check boundary to avoid out-of-bounds access
                grid[x][y] = 'X'
                found = false // Stop the loop
                break
            }
        }
    }

    foundX := FoundX(grid, 'X')
    return grid, foundX
}
/* // available on helps.go
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
*/
func GridRuneToString(grid [][]rune) string {
    result := ""
    for i, row := range grid {
        result += string(row)
        if i < len(grid)-1 {
            result += "\n"
        }
    }
    return result
}

func Historian2(inputStr string) int {
    // Convert input string into a 2D slice (grid) of characters
    lines := strings.Split(strings.TrimSpace(inputStr), "\n")
    grid := make([][]rune, len(lines))
    for i := range lines {
        grid[i] = []rune(lines[i])
    }

    n := len(grid)      // Number of rows in the grid
    m := len(grid[0])   // Number of columns in the grid

    // Find the initial position of the character '^'
    found := false
    var i, j int
    for i = 0; i < n; i++ {
        for j = 0; j < m; j++ {
            if grid[i][j] == '^' {
                found = true
                break
            }
        }
        if found {
            break
        }
    }

    ii, jj := i, j  // Initial position of '^'
    dd := [][]int{
        {-1, 0},  // up
        {0, 1},   // right
        {1, 0},   // down
        {0, -1},  // left
    }

    dir := 0
    ogSeen := make(map[[2]int]bool)
    for {
        ogSeen[[2]int{i, j}] = true

        // Calculate next position based on current direction
        nextI := i + dd[dir][0]
        nextJ := j + dd[dir][1]

        // Check if the next position is within grid boundaries
        if !(0 <= nextI && nextI < n && 0 <= nextJ && nextJ < m) {
            break
        }

        // Change direction if the next position is a wall '#'
        if grid[nextI][nextJ] == '#' {
            dir = (dir + 1) % 4  // Rotate direction clockwise
        } else {
            i, j = nextI, nextJ
        }
    }

    Will_Loop := func(oi, oj int) bool {
        if grid[oi][oj] == '#' {
            return false  // Cannot place obstacle on a wall '#'
        }

        grid[oi][oj] = '#'  // Temporarily place obstacle
        i, j = ii, jj

        dir = 0
        seen := make(map[[3]int]bool)
        for {
            if seen[[3]int{i, j, dir}] {
                grid[oi][oj] = '.'  // Reset the temporary obstacle
                return true  // Infinite loop detected
            }
            seen[[3]int{i, j, dir}] = true

            nextI := i + dd[dir][0]
            nextJ := j + dd[dir][1]

            if !(0 <= nextI && nextI < n && 0 <= nextJ && nextJ < m) {
                grid[oi][oj] = '.'  // Reset the temporary obstacle
                return false  // Not an infinite loop
            }

            if grid[nextI][nextJ] == '#' {
                dir = (dir + 1) % 4  // Rotate direction clockwise
            } else {
                i, j = nextI, nextJ
            }
        }
    }

    ans := 0
    for pos := range ogSeen {
        oi, oj := pos[0], pos[1]
        if oi == ii && oj == jj {
            continue
        }
        if Will_Loop(oi, oj) {
            ans++
        }
    }

    return ans
}