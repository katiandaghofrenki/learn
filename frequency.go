package learn

import (
    "strings"
)

type Node struct {
    Freq int
    Row  int
    Col  int
}

func Frequency(input string) int {
    lines := strings.Split(input, "\n")
    nodes := make(map[rune][]Node)

    for row := range lines {
        for col, char := range lines[row] {
            if char != '.' {
                nodes[char] = append(nodes[char], Node{Freq: int(char), Row: row, Col: col})
            }
        }
    }

    antinodes := make(map[[2]int]bool)
    for _, nodeList := range nodes {
        for _, node := range nodeList {
            for _, node2 := range nodeList {
                if node == node2 {
                    continue
                }
                dx := node2.Col - node.Col
                dy := node2.Row - node.Row
				destRow := node.Row - dy
				destCol := node.Col - dx
				if destRow >= 0 && destRow < len(lines) && destCol >= 0 && destCol < len(lines[0]) {
					antinodes[[2]int{destRow, destCol}] = true
				}
            }
        }
    }

    return len(antinodes)
}

func Frequency2(input string) int {
    lines := strings.Split(input, "\n")
    nodes := make(map[rune][]Node)

    for row := range lines {
        for col, char := range lines[row] {
            if char != '.' {
                nodes[char] = append(nodes[char], Node{Freq: int(char), Row: row, Col: col})
            }
        }
    }

    antinodes := make(map[[2]int]bool)
    for _, nodeList := range nodes {
        for _, node := range nodeList {
            for _, node2 := range nodeList {
                if node == node2 {
                    continue
                }
                dx := node2.Col - node.Col
                dy := node2.Row - node.Row
                onboard := true
                totaldx := dx
                totaldy := dy
                for onboard {
                    destRow := node.Row - totaldy
                    destCol := node.Col - totaldx
                    if destRow >= 0 && destRow < len(lines) && destCol >= 0 && destCol < len(lines[0]) {
                        antinodes[[2]int{destRow, destCol}] = true
                        totaldx += dx
                        totaldy += dy
                    } else {
                        onboard = false
                    }
                }
                antinodes[[2]int{node.Row, node.Col}] = true
            }
        }
    }

    return len(antinodes)
}