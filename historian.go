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

/*
usage : on your func main()
input := `copy your input data here`
grid, count := learn.Historian(input)
	fmt.Println("Grid State:")
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Printf("Count of 'X': %d\n", count)
*/

// part two under construction

func Historian2(s string) ([][]rune, int) {
    lines := strings.Split(s, "\n")
    
    grid := make([][]rune, len(lines))
    for i, line := range lines {
        parts := strings.TrimSpace(line)
        grid[i] = []rune(parts)
    }

    count := 0

    char := '^'
    x, y, found := Found(grid, char)

    for i := 0; i < len(grid)*len(grid[0]); i++ {
        if found {
            if x > 0 && grid[x-1][y] != '#' {
                grid[x-1][y] = '|'
                x = x - 1
            } else if x > 0 && grid[x-1][y] == '#' && grid[x][y] != '^' {
                grid[x][y] = '+'
                // Fill '-' all the way to the '#'
                if grid[x][y] == '+' && grid[x-1][y] == '#' && y+1 < len(grid[0]) && grid[x][y+1] != '#' {
                    for y+1 < len(grid[0]) && grid[x][y+1] != '#' {
                        grid[x][y+1] = '-'
                        y = y + 1
                    }
                    // Change direction to down when hitting a '#'
                    if y+1 < len(grid[0]) && grid[x][y+1] == '#' {
                        grid[x][y] = '+'
                        if grid[x][y] == '+' && grid[x][y+1] == '#' && x+1 < len(grid) && grid[x+1][y] != '#' {
                            for x+1 < len(grid) && grid[x+1][y] != '#' {
                                grid[x+1][y] = '|'
                                x = x + 1
                            }
                            // Change direction to left when hitting a '#'
                            if x+1 < len(grid) && grid[x+1][y] == '#' {
                                grid[x][y] = '+'
                                if grid[x][y] == '+' && grid[x+1][y] == '#' && y > 0 && grid[x][y-1] != '#' {
                                    for y > 0 && grid[x][y-1] != '#' {
                                        if grid[x][y-1] != '|' || grid[x][y-1] != '+' || grid[x][y-1] != '^' {
                            
                                            if grid[x][y-1] == '^' {
                                                grid[x][y-1] = '^'
                                                grid[x][y-2] = '0'
                                                break
                                            } else {
                                                grid[x][y-1] = '-'
                                                y = y - 1
                                            }

                                        }
                                        
                                    }
                                    // Change direction to left when hitting a '#'
                                    if y > 0 && grid[x][y-1] == '#' {
                                        grid[x][y] = '+'
                                    }
                                }
                            }
                        }
                    }
                }
            }
        } else {
            continue
        }
    }
    

    
    return grid, count // not yet done
}