package learn

import (
    "strings"
)

func InGrid(i, j, n int) bool {
    return (0 <= i && i < n) && (0 <= j && j < n)
}

func CountFree(i, j int, plot [][]int, grid []string, n int, dd [][2]int) int {
    ans := 0
    for _, d := range dd {
        ii, jj := i+d[0], j+d[1]
        if !InGrid(ii, jj, n) {
            ans++
        } else if grid[ii][jj] != grid[i][j] {
            ans++
        }
    }
    return ans
}

func Perimeter(plot [][]int, grid []string, n int, dd [][2]int) int {
    ans := 0
    for _, p := range plot {
        i, j := p[0], p[1]
        ans += CountFree(i, j, plot, grid, n, dd)
    }
    return ans
}

func CalculatePerimeter(inputString string) int {
    grid := strings.Split(strings.TrimSpace(inputString), "\n")
    n := len(grid)
    dd := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

    seen := make(map[[2]int]bool)
    var plots [][][]int

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if seen[[2]int{i, j}] {
                continue
            }

            stack := [][2]int{{i, j}}
            plots = append(plots, [][]int{{i, j}})
            seen[[2]int{i, j}] = true

            for len(stack) > 0 {
                ci, cj := stack[len(stack)-1][0], stack[len(stack)-1][1]
                stack = stack[:len(stack)-1]

                for _, d := range dd {
                    ii, jj := ci+d[0], cj+d[1]
                    if !InGrid(ii, jj, n) {
                        continue
                    }
                    if grid[ii][jj] != grid[i][j] {
                        continue
                    }
                    if seen[[2]int{ii, jj}] {
                        continue
                    }
                    seen[[2]int{ii, jj}] = true
                    stack = append(stack, [2]int{ii, jj})
                    plots[len(plots)-1] = append(plots[len(plots)-1], []int{ii, jj})
                }
            }
        }
    }

    ans := 0
    for _, plot := range plots {
        ans += Perimeter(plot, grid, n, dd) * len(plot)
    }
    return ans
}