package learn

import (
    "strings"
)

type Point1 struct {
    i, j int
}

func InGrid(i, j, n int) bool {
    return 0 <= i && i < n && 0 <= j && j < n
}

func CountFree(i, j int, grid [][]int, n int, dd [][2]int) int {
    ans := 0
    for _, d := range dd {
        ii, jj := i+d[0], j+d[1]
        if !InGrid(ii, jj, n) || grid[ii][jj] != grid[i][j] {
            ans++
        }
    }
    return ans
}

func Perimeter(plot [][2]int, grid [][]int, n int, dd [][2]int) int {
    ans := 0
    for _, p := range plot {
        i, j := p[0], p[1]
        ans += CountFree(i, j, grid, n, dd)
    }
    return ans
}

func Sides(plot [][2]int, grid [][]int, n int, dd [][2]int) int {
    up, down, left, right := make(map[Point1]bool), make(map[Point1]bool), make(map[Point1]bool), make(map[Point1]bool)
    plotSet := make(map[Point1]bool)

    // Convert plot slice to a set for quick lookup
    for _, p := range plot {
        plotSet[Point1{p[0], p[1]}] = true
    }

    // Iterate over all points in the plot
    for _, p := range plot {
        i, j := p[0], p[1]
        if !plotSet[Point1{i - 1, j}] {
            up[Point1{i, j}] = true
        }
        if !plotSet[Point1{i + 1, j}] {
            down[Point1{i, j}] = true
        }
        if !plotSet[Point1{i, j - 1}] {
            left[Point1{i, j}] = true
        }
        if !plotSet[Point1{i, j + 1}] {
            right[Point1{i, j}] = true
        }
    }

    count := 0

    // Iterate over the points on the top edge
    for p := range up {
        if left[p] {
            count++
        }
        if right[p] {
            count++
        }
        if right[Point1{p.i - 1, p.j - 1}] && !left[p] {
            count++
        }
        if left[Point1{p.i - 1, p.j + 1}] && !right[p] {
            count++
        }
    }

    // Iterate over the points on the bottom edge
    for p := range down {
        if left[p] {
            count++
        }
        if right[p] {
            count++
        }
        if right[Point1{p.i + 1, p.j - 1}] && !left[p] {
            count++
        }
        if left[Point1{p.i + 1, p.j + 1}] && !right[p] {
            count++
        }
    }

    return count
}

func CalculatePerimeter(inputString string) (string, int, string, int) {
    grid := [][]int{}
    for _, line := range strings.Split(inputString, "\n") {
        row := []int{}
        for _, char := range line {
            row = append(row, int(char)-'0')
        }
        grid = append(grid, row)
    }
    n := len(grid)
    dd := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

    seen := make(map[Point1]bool)
    plots := []struct {
        value int
        points [][2]int
    }{}

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if seen[Point1{i, j}] {
                continue
            }
            stack := []Point1{{i, j}}
            plot := struct {
                value  int
                points [][2]int
            }{grid[i][j], nil}

            for len(stack) > 0 {
                ci, cj := stack[len(stack)-1].i, stack[len(stack)-1].j
                stack = stack[:len(stack)-1]
                if seen[Point1{ci, cj}] || !InGrid(ci, cj, n) || grid[ci][cj] != grid[i][j] {
                    continue
                }
                seen[Point1{ci, cj}] = true
                plot.points = append(plot.points, [2]int{ci, cj})

                for _, d := range dd {
                    ii, jj := ci+d[0], cj+d[1]
                    stack = append(stack, Point1{ii, jj})
                }
            }
            plots = append(plots, plot)
        }
    }

    part1, part2 := 0, 0
    for _, plot := range plots {
        part1 += len(plot.points) * Perimeter(plot.points, grid, n, dd)
        part2 += len(plot.points) * Sides(plot.points, grid, n, dd)
    }

    return "result part1: ", part1, "result part2 : ", part2
}