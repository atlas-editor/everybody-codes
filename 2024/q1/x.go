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

	fmt.Println(p1(input))
	fmt.Println(p2(input))
	fmt.Println(p3(input))
}

func p1(input string) int {
	notes := map[byte]int{}
	for i := 0; i < len(input); i++ {
		notes[input[i]]++
	}

	return notes['B'] + notes['C']*3
}

func p2(input string) int {
	res := 0
	for i := 0; i < len(input)-1; i += 2 {
		a, b := input[i], input[i+1]
		res += cost[a] + cost[b]
		if a != 'x' && b != 'x' {
			res += 2
		}
	}

	return res
}

func p3(input string) int {
	res := 0
	for i := 0; i < len(input)-2; i += 3 {
		a, b, c := input[i], input[i+1], input[i+2]
		res += cost[a] + cost[b] + cost[c]
		switch count(input[i:i+3], 'x') {
		case 1:
			res += 2
		case 0:
			res += 6
		}
	}

	return res
}

var cost = map[byte]int{'A': 0, 'B': 1, 'C': 3, 'D': 5}

func count(s string, b byte) (c int) {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			c++
		}
	}
	return
}
