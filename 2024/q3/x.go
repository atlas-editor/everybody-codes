package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])
	input := strings.TrimSpace(string(data))

	switch os.Args[2] {
	case "1":
		fmt.Println(p(input, neighbors, false))
	case "2":
		fmt.Println(p(input, neighbors, false))
	case "3":
		fmt.Println(p(input, neighbors8, true))
	}
}

func p(input string, nbrFunc func(int, int, int, int) []pt, pad bool) int {
	lines := strings.Split(input, "\n")

	R, C := len(lines), len(lines[0])

	grid := make([][]byte, R)
	res := 0

	for r, line := range lines {
		grid[r] = make([]byte, C)
		for c, ch := range line {
			if ch == '.' {
				grid[r][c] = 0
			} else {
				grid[r][c] = 1
				res++
			}
		}
	}

	if pad {
		padded := make([][]byte, R+2)
		for r := range R + 2 {
			padded[r] = make([]byte, C+2)
			for c := range C + 2 {
				padded[r][c] = 0
			}
		}

		for r := range R {
			for c := range C {
				padded[r+1][c+1] = grid[r][c]
			}
		}

		grid = padded
		R = R + 2
		C = C + 2
	}

	removed := true

	for removed {
		removed = false
		for r := range R {
		outer:
			for c := range C {
				v := grid[r][c]
				if v == 0 {
					continue
				}
				for _, n := range nbrFunc(r, c, R, C) {
					if grid[n[0]][n[1]] < v {
						continue outer
					}
				}
				grid[r][c]++
				removed = true
				res++
			}
		}
	}

	return res
}

type pt [2]int

func neighbors(r, c, R, C int) (n []pt) {
	for _, d := range []pt{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		rr, cc := r+d[0], c+d[1]
		if 0 <= rr && rr < R && 0 <= cc && cc < C {
			n = append(n, pt{rr, cc})
		}
	}
	return
}

func neighbors8(r, c, R, C int) (n []pt) {
	for _, d := range []pt{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}} {
		rr, cc := r+d[0], c+d[1]
		if 0 <= rr && rr < R && 0 <= cc && cc < C {
			n = append(n, pt{rr, cc})
		}
	}
	return
}
