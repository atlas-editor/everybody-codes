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
	lines := strings.Split(input, "\n")

	rules := map[string][]string{}
	for _, line := range lines {
		tmp := strings.Split(line, ":")
		from := tmp[0]
		to := strings.Split(tmp[1], ",")

		rules[from] = to
	}

	switch os.Args[2] {
	case "1":
		fmt.Println(simulate(map[string]int{"A": 1}, rules, 4))
	case "2":
		fmt.Println(simulate(map[string]int{"Z": 1}, rules, 10))
	case "3":
		fmt.Println(p3(rules))
	}
}

func p3(rules map[string][]string) int {
	c := make(chan int)
	for t := range rules {
		go func() {
			c <- simulate(map[string]int{t: 1}, rules, 20)
		}()
	}

	max_, min_ := 0, 1<<62
	for range len(rules) {
		n := <-c
		max_ = max(max_, n)
		min_ = min(min_, n)
	}

	return max_ - min_
}

func simulate(termites map[string]int, rules map[string][]string, days int) int {
	for range days {
		tmp := map[string]int{}
		for k, v := range termites {
			for _, vv := range rules[k] {
				tmp[vv] += v
			}
		}
		termites = tmp
	}

	res := 0
	for _, v := range termites {
		res += v
	}

	return res
}
