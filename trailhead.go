package learn

import (
    "strconv"
    "strings"
)

var dirs = [][]int{
    {0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func Trailhead(input string) int {
    lines := strings.Split(input, "\n")
    trailheads := [][]int{}
    for row := 0; row < len(lines); row++ {
        for col := 0; col < len(lines[0]); col++ {
            if lines[row][col] == '0' {
                trailheads = append(trailheads, []int{row, col})
            }
        }
    }

    total := 0
    for _, th := range trailheads {
        v := [][]int{}
        summits := make(map[[2]int]bool)
        q := [][]int{th}
        for len(q) > 0 {
            current := q[0]
            q = q[1:]
            currentrow, currentcol := current[0], current[1]
            v = append(v, []int{currentrow, currentcol})

            if lines[currentrow][currentcol] == '9' {
                summits[[2]int{currentrow, currentcol}] = true
            } else {
                currentheight, _ := strconv.Atoi(string(lines[currentrow][currentcol]))
                for _, dir := range dirs {
                    testrow := currentrow + dir[0]
                    testcol := currentcol + dir[1]
                    if testrow < 0 || testrow >= len(lines) || testcol < 0 || testcol >= len(lines[0]) {
                        continue
                    }
                    testheight, _ := strconv.Atoi(string(lines[testrow][testcol]))
                    if testheight == currentheight+1 {
                        if !contains(v, []int{testrow, testcol}) && !contains(q, []int{testrow, testcol}) {
                            q = append(q, []int{testrow, testcol})
                        }
                    }
                }
            }
        }
        total += len(summits)
    }

    return total
}

func contains(slice [][]int, item []int) bool {
    for _, v := range slice {
        if v[0] == item[0] && v[1] == item[1] {
            return true
        }
    }
    return false
}

type Point struct {
    Row, Col int
}

type Nodes struct {
    Pos  Point
    Path []Point
}

func Trailhead2(input string) int {
    lines := strings.Split(input, "\n")
    dirs := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    var trailheads []Point

    for row := 0; row < len(lines); row++ {
        for col := 0; col < len(lines[0]); col++ {
            if lines[row][col] == '0' {
                trailheads = append(trailheads, Point{row, col})
            }
        }
    }

    totalRoutes := 0
    for _, th := range trailheads {
        q := []Nodes{{Pos: th, Path: []Point{}}}

        for len(q) > 0 {
            node := q[0]
            q = q[1:]
            currentrow, currentcol := node.Pos.Row, node.Pos.Col

            if lines[currentrow][currentcol] == '9' {
                totalRoutes++
            } else {
                currentheight, _ := strconv.Atoi(string(lines[currentrow][currentcol]))

                for _, dir := range dirs {
                    testrow := currentrow + dir.Row
                    testcol := currentcol + dir.Col

                    if !(0 <= testrow && testrow < len(lines) && 0 <= testcol && testcol < len(lines[0])) {
                        continue
                    }

                    testheight, _ := strconv.Atoi(string(lines[testrow][testcol]))

                    if testheight == currentheight+1 {
                        newPath := append([]Point{}, node.Path...)
                        newPath = append(newPath, Point{currentrow, currentcol})
                        q = append(q, Nodes{Pos: Point{testrow, testcol}, Path: newPath})
                    }
                }
            }
        }
    }

    return totalRoutes
}