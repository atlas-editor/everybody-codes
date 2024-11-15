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
	nums := ints(input)

	switch os.Args[2] {
	case "1":
		fmt.Println(p1(nums, []int{1, 3, 5, 10}))
	case "2":
		fmt.Println(p2(nums, []int{1, 3, 5, 10, 15, 16, 20, 24, 25, 30}))
	case "3":
		fmt.Println(p3(nums, []int{1, 3, 5, 10, 15, 16, 20, 24, 25, 30, 37, 38, 49, 50, 74, 75, 100, 101}))
	}
}

func p1(nums []int, stamps []int) int {
	slices.Sort(stamps)
	slices.Reverse(stamps)

	res := 0
	for _, n := range nums {
		for _, s := range stamps {
			for n >= s {
				n -= s
				res++
			}
		}
	}

	return res
}

func p2(nums []int, stamps []int) int {
	res := 0
	for _, n := range nums {
		res += beetles(n, stamps)
	}

	return res
}

func p3(nums []int, stamps []int) int {
	res := 0
	for _, n := range nums {
		min_ := 1 << 32
		for _, s := range splits(n) {
			min_ = min(min_, beetles(s[0], stamps)+beetles(s[1], stamps))
		}
		res += min_
	}

	return res
}

func ints(s string) (r []int) {
	p := regexp.MustCompile(`-?\d+`)
	for _, e := range p.FindAllString(s, -1) {
		a, _ := strconv.Atoi(e)
		r = append(r, a)
	}
	return
}

var cache = map[int]int{}

func beetles(num int, stamps []int) int {
	var dp func(int) int
	dp = func(i int) int {
		if i == 0 {
			return 0
		}

		if v, ok := cache[i]; ok {
			return v
		}

		min_ := 1 << 32
		for _, s := range stamps {
			if i >= s {
				min_ = min(min_, dp(i-s)+1)
			}
		}

		cache[i] = min_
		return min_
	}

	return dp(num)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func splits(num int) [][]int {
	var s [][]int
	for i := -50; i <= 50; i++ {
		a := num/2 + i
		b := num - a
		if abs(b-a) <= 100 {
			s = append(s, []int{a, b})
		}
	}
	return s
}
