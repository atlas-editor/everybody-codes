package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	path := os.Args[1]
	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	fmt.Println(p(input, true))
	fmt.Println(p(input, false))
	fmt.Println(p(input, false))
}

func p(input string, p1 bool) string {
	edges := map[string][]string{}
	for _, line := range strings.Split(input, "\n") {
		tmp := strings.Split(line, ":")
		from := tmp[0]
		to := strings.Split(tmp[1], ",")
		edges[from] = to
	}

	q := []string{"RR"}
	dist := map[string]int{"RR": 0}
	branchLen := map[string]int{}
	lens := map[int]int{}
	parent := map[string]string{"RR": ""}

	for len(q) > 0 {
		curr := pop(&q)

		for _, nbr := range edges[curr] {
			d := dist[curr]
			switch nbr {
			case "@":
				branchLen[curr] = d + 1
				lens[d+1]++
			case "BUG", "ANT":
			default:
				dist[nbr] = d + 1
				parent[nbr] = curr
				q = append(q, nbr)
			}
		}
	}

	l := -1
	for k, v := range lens {
		if v == 1 {
			l = k
			break
		}
	}

	b := ""
	for k, v := range branchLen {
		if v == l {
			b = k
			break
		}
	}

	s := getPath(parent, b)

	if p1 {
		return strings.Join(s, "")
	} else {
		var a []string

		for _, p := range s {
			a = append(a, string(p[0]))
		}

		return strings.Join(a, "")
	}
}

func pop[T any](slice *[]T) T {
	n := len(*slice)
	if n == 0 {
		panic("empty slice")
	}
	back := (*slice)[n-1]
	*slice = (*slice)[:n-1]
	return back
}

func getPath(parent map[string]string, start string) []string {
	s := []string{"@"}
	for start != "" {
		s = append(s, start)
		start = parent[start]
	}
	slices.Reverse(s)
	return s
}
