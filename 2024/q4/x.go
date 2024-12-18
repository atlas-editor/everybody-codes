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
	data, _ := os.ReadFile(os.Args[1])
	input := strings.TrimSpace(string(data))

	switch os.Args[2] {
	case "1":
		fmt.Println(p(input))
	case "2":
		fmt.Println(p(input))
	case "3":
		fmt.Println(p3(input))
	}
}

func p(input string) int {
	nums := ints(input)

	min_ := slices.Min(nums)
	res := 0
	for _, n := range nums {
		res += n - min_
	}
	return res
}

func p3(input string) int {
	nums := ints(input)
	slices.Sort(nums)

	median := nums[len(nums)/2]
	res := 0
	for _, n := range nums {
		res += abs(n - median)
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ints(s string) (r []int) {
	p := regexp.MustCompile(`-?\d+`)
	for _, e := range p.FindAllString(s, -1) {
		a, _ := strconv.Atoi(e)
		r = append(r, a)
	}
	return
}
