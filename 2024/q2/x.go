package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])
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

func p1(input string) int {
	parts := strings.Split(input, "\n\n")
	runic := strings.Split(strings.Split(parts[0], ":")[1], ",")
	words := strings.Fields(parts[1])

	res := 0
	for _, w := range words {
		for _, r := range runic {
			if strings.Contains(w, r) {
				res++
			}
		}
	}

	return res
}

func p2(input string) int {
	parts := strings.Split(input, "\n\n")
	runic := strings.Split(strings.Split(parts[0], ":")[1], ",")
	lines := strings.Split(parts[1], "\n")

	res := 0
	for _, line := range lines {
		for _, w := range strings.Fields(line) {
			indices := map[int]bool{}
			for _, r := range runic {
				re := regexp.MustCompile(r)

				for _, m := range re.FindAllStringIndex(w, -1) {
					for j := m[0]; j < m[1]; j++ {
						indices[j] = true
					}
				}

				for _, m := range re.FindAllStringIndex(reverse(w), -1) {
					for j := m[0]; j < m[1]; j++ {
						indices[len(w)-j-1] = true
					}
				}
			}

			res += len(indices)
		}

	}

	return res
}

func p3(input string) int {
	parts := strings.Split(input, "\n\n")
	runic := strings.Split(strings.Split(parts[0], ":")[1], ",")
	lines := strings.Split(parts[1], "\n")

	grid := make([][]bool, len(lines))
	for i := range lines {
		grid[i] = make([]bool, len(lines[0]))
	}

	for i, w := range lines {
		for _, r := range runic {
			for j := 0; j < len(w); j++ {
				m := match(r, w, j, 2)
				if m {
					for k := range len(r) {
						grid[i][(j+k)%len(w)] = true
					}
				}

				m = match(r, reverse(w), j, 2)
				if m {
					for k := range len(r) {
						grid[i][posMod(len(w)-j-1-k, len(w))] = true
					}
				}
			}
		}
	}

	for i, w := range vertLines(lines) {
		for _, r := range runic {
			for j := 0; j < len(w); j++ {
				m := match(r, w, j, 1)
				if m {
					for k := range len(r) {
						grid[(j+k)%len(w)][i] = true
					}
				}

				m = match(r, reverse(w), j, 1)
				if m {
					for k := range len(r) {
						grid[posMod(len(w)-j-1-k, len(w))][i] = true
					}
				}
			}
		}
	}

	res := 0
	for _, line := range grid {
		for _, e := range line {
			if e {
				res++
			}
		}
	}

	return res
}

func reverse(s string) (r string) {
	tmp := []byte(s)
	slices.Reverse(tmp)
	return string(tmp)
}

func match(test, sample string, idx int, r int) bool {
	return strings.HasPrefix(strings.Repeat(sample, r)[idx:], test)
}

func posMod(x, m int) int {
	x = x % m
	for x < 0 {
		x += m
	}
	return x
}

func vertLines(lines []string) []string {
	v := make([]string, len(lines[0]))

	for _, line := range lines {
		for i, r := range line {
			v[i] += string(r)
		}
	}
	return v
}
