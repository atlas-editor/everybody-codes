package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	path := os.Args[1]
	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	fmt.Println(p(input, 1))
	fmt.Println(p(input, 2))
	fmt.Println(p(input, 3))
}

func p(input string, part int) int {
	nums := ints(input)

	var columns [4][]int

	for i, n := range nums {
		columns[i%4] = append(columns[i%4], n)
	}

	//0 1 2 ... N-1 N N-1 N-2 ... 1 0 1 2 ...

	shouts := map[int]int{}
	max_ := 0
	i := 0
	for {
		if i == 10 && part == 1 {
			return shouted(columns)
		}
		curr := i % 4
		next := (i + 1) % 4
		clapper := popFront(&columns[curr])

		columns[next] = slices.Insert(columns[next], insertionIdx(len(columns[next]), clapper), clapper)

		s := shouted(columns)
		if s > max_ {
			max_ = s
		}
		shouts[s]++
		if shouts[s] == 2024 && part == 2 {
			return (i + 1) * s
		}
		i++
		if i > 2000 && part == 3 {
			return max_
		}
	}
}

func ints(s string) (r []int) {
	p := regexp.MustCompile(`-?\d+`)
	for _, e := range p.FindAllString(s, -1) {
		a, _ := strconv.Atoi(e)
		r = append(r, a)
	}
	return
}

func popFront[T any](slice *[]T) T {
	n := len(*slice)
	if n == 0 {
		panic("empty slice")
	}
	front := (*slice)[0]
	*slice = (*slice)[1:]
	return front
}

func insertionIdx(N, c int) int {
	m := (c - 1) % (2 * N)

	if m <= N {
		return m
	} else {
		return 2*N - m
	}
}

func shouted(cols [4][]int) int {
	n, _ := strconv.Atoi(fmt.Sprintf("%v%v%v%v", cols[0][0], cols[1][0], cols[2][0], cols[3][0]))
	return n
}
