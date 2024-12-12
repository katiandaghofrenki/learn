package learn

func DiscMap(s string) (string, int, string, int) {
    transform := Transforms(s)
    reposition := Reposition(transform)
    total := Calculates(reposition)
	
	transform2 := Transforms(s)
	reposition2 := Reposition2(transform2)
	total2 := Calculates(reposition2)
    return "part1 result: ", total, "part2 result: ", total2
}

func Transforms(s string) []string {
    var result []string

    for index, digit := range s {
        count := int(digit - '0')
        if index%2 == 0 {
            // Even index
            for i := 0; i < count; i++ {
                result = append(result, IntToString(index/2)) //func IntToString is on other file
            }
        } else {
            // Odd index
            for i := 0; i < count; i++ {
                result = append(result, ".")
            }
        }
    }

    return result
}

func Calculates(s []string) int {
    total := 0
    for i := 0; i < len(s); i++ {
        if s[i] != "." {
            num := ToInt(s[i])*i
            total += num
        }
    }
    return total
}

func Reposition(s []string) []string {
    i, j := 0, len(s)-1

    for i < j {
        for i < len(s) && s[i] != "." {
            i++
        }
        for j >= 0 && s[j] == "." {
            j--
        }
        if i < j {
            s[i], s[j] = s[j], s[i]
            i++
            j--
        }
    }

    return s
}

func Reposition2(s []string) []string {
    n := len(s)
    for fileID := n; fileID >= 0; fileID-- {
        fileStr := IntToString(fileID)
        fileSize := 0
        fileStart := -1

        // Find the file size and start position
        for i := 0; i < n; i++ {
            if s[i] == fileStr {
                if fileStart == -1 {
                    fileStart = i
                }
                fileSize++
            } else if fileStart != -1 {
                break
            }
        }

        if fileSize > 0 {
            // Find the leftmost span of free space large enough for the file
            freeSize := 0
            freeStart := -1
            for i := 0; i < fileStart; i++ {
                if s[i] == "." {
                    if freeStart == -1 {
                        freeStart = i
                    }
                    freeSize++
                    if freeSize >= fileSize {
                        break
                    }
                } else {
                    freeSize = 0
                    freeStart = -1
                }
            }

            // Move the file if there's enough space
            if freeSize >= fileSize {
                for i := 0; i < fileSize; i++ {
                    s[freeStart+i] = fileStr
                    s[fileStart+i] = "."
                }
            }
        }
    }

    return s
}