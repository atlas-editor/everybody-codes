package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := os.Args[1]
	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	switch os.Args[2] {
	case "1":
		fmt.Println(p1(input))
	case "2":
		fmt.Println(p2(input))
	case "3":
		fmt.Println(p3(input))
	}
}

func p1(input string) string {
	var grid [8][8]byte
	for r, line := range strings.Split(input, "\n") {
		for c := range len(line) {
			grid[r][c] = line[c]
		}
	}

	res := ""
	i := 1
	for r := 2; r <= 5; r++ {
		for c := 2; c <= 5; c++ {
			ch, _ := rcEl(grid, r, c)
			res += string(ch)
			i++
		}
	}

	return res
}

func p2(input string) int {
	parts := strings.Split(input, "\n\n")
	perRow := len(strings.Fields(strings.Split(parts[0], "\n")[0]))

	res := 0
	for _, p := range parts {
		grids := make([][8][8]byte, perRow)
		for r, line := range strings.Split(p, "\n") {
			for i, gridRow := range strings.Fields(line) {
				for c := range len(gridRow) {
					grids[i][r][c] = gridRow[c]
				}
			}
		}

		for _, g := range grids {
			i := 1
			for r := 2; r <= 5; r++ {
				for c := 2; c <= 5; c++ {
					ch, _ := rcEl(g, r, c)
					res += i * int(ch-64)
					i++
				}
			}
		}
	}

	return res
}

func p3(input string) int {
	lines := strings.Split(input, "\n")

	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	res := 0
	for {
		tmp := 0
		for i := range 10 {
			for j := range 20 {

				// small grid
				var small [8][8]byte
				for r := range 8 {
					for c := range 8 {
						small[r][c] = grid[6*i+r][6*j+c]
					}
				}

				solved := solve(small)

				for r := range 8 {
					for c := range 8 {
						grid[6*i+r][6*j+c] = solved[r][c]
					}
				}

				k := 1
				for r := 2; r <= 5; r++ {
					for c := 2; c <= 5; c++ {
						if solved[r][c] != '.' {
							tmp += k * int(solved[r][c]-64)
							k++
						}
					}
				}
			}
		}

		if tmp == res {
			return res
		}
		res = tmp
	}
}

// find the correct element to fill in at point (r, c) in `grid`, the second value is false if we are not able to fill
// it in, either because there is no duplicate elements or there is a ? in that row/column
func rcEl(grid [8][8]byte, r, c int) (byte, bool) {
	counter := map[byte]int{}

	for _, idx := range []pt{{0, c}, {1, c}, {6, c}, {7, c}, {r, 0}, {r, 1}, {r, 6}, {r, 7}} {
		counter[grid[idx[0]][idx[1]]]++
	}

	sol := byte(0)
	cc := 0
	for k, v := range counter {
		if v == 2 && k != '?' {
			sol = k
			cc++
		}
	}

	if cc == 1 {
		return sol, true
	}

	return 0, false
}

type pt [2]int

func solve(small [8][8]byte) [8][8]byte {
	// in unsolvable cases we just return tje original small grid
	orig := small

	seen := map[byte]bool{}

	// fill in elements where there are no ? in that row/column
	for r := 2; r <= 5; r++ {
		for c := 2; c <= 5; c++ {
			if ch, ok := rcEl(orig, r, c); ok {
				// this means there is a conflict, we want to fill in a char that was previously used
				if seen[ch] {
					return orig
				}
				seen[ch] = true
				small[r][c] = ch
			}
		}
	}

	// here we try to infer elements with ? in their respective rows/columns
	for r := 2; r <= 5; r++ {
		for c := 2; c <= 5; c++ {
			if small[r][c] == '.' {

				possible := map[byte]bool{}
				qMark := pt{-11, 0}
				for _, idx := range []pt{{0, c}, {1, c}, {6, c}, {7, c}, {r, 0}, {r, 1}, {r, 6}, {r, 7}} {
					possible[small[idx[0]][idx[1]]] = true
					if small[idx[0]][idx[1]] == '?' {
						qMark = idx
					}
				}

				// this means we have a row and column with no ? but also no duplicate element
				if qMark[0] == -11 {
					return orig
				}

				for k, v := range possible {
					if v && !seen[k] && k != '?' {
						small[r][c] = k
						small[qMark[0]][qMark[1]] = k
						seen[k] = true
					}
				}

				// there is a ? in this row or column, but it is not solvable as no possible char can be filled in due
				//to it being used previously
				if small[r][c] == '.' {
					return orig
				}
			}

		}
	}

	return small
}
