package learn

import (
	"strings"
)

func Teleport(s string) (int, int, int, int, int) {
	lines := strings.Split(s, "\n")

	var ilines [][]int
	var filines [][]int

	width := 101
	height := 103

	
	total := 0
	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	for _, line := range lines {
		
		if len(line) < 1 {
			continue
		}
		var iline []int
		var filine []int
		parts := strings.Split(line, " ")
		lhs := strings.Split(parts[0], "p=")
		rhs := strings.Split(parts[1], "v=")
		lh := strings.Split(lhs[1], ",")
		rh := strings.Split(rhs[1], ",")
		px := ToInt(lh[0]) 
		py := ToInt(lh[1]) 
		vx := ToInt(rh[0]) 
		vy := ToInt(rh[1]) 

		iline = append(iline, px, py, vx, vy)
		ilines = append(ilines, iline)

		vxf := (vx*100) %width
		// if vxf < 0 {
		// 	vxf += width
		// }
		vyf := (vy*100) %height
		// if vyf < 0 {
		// 	vyf += height
		// }
		xf := (px + vxf) % width
		if xf < 0 {
			xf += width
		}
		yf := (py + vyf) % height
		if yf < 0 {
			yf += height
		}
		

		midw := (width-1) / 2
		midh := (height-1) / 2

		if xf < midw {
			if yf < midh {
				q1 += 1
			} else if yf > midh {
				q3 += 1
			}
		} else if xf > midw {
			if yf < midh {
				q2 += 1
			} else if yf > midh {
				q4 += 1
			}
		}
		total = q1*q2*q3*q4

		filine = append(filine, xf, yf, vxf, vyf)
		filines = append(filines, filine)
		
	}
	
	return total, q1, q2, q3, q4
}