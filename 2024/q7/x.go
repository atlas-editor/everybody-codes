package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])
	input := strings.TrimSpace(string(data))

	switch os.Args[2] {
	case "1":
		fmt.Println(p(input, 1))
	case "2":
		fmt.Println(p(input, 2))
	case "3":
		fmt.Println(p(input, 3))
	}
}

func p(input string, part int) string {
	lines := strings.Split(input, "\n")

	track := "="
	switch part {
	case 2:
		track = readTrack("S-=++=-==++=++=-=+=-=+=+=--=-=++=-==++=-+=-=+=-=+=+=++=-+==++=++=-=-=--\n-                                                                     -\n=                                                                     =\n+                                                                     +\n=                                                                     +\n+                                                                     =\n=                                                                     =\n-                                                                     -\n--==++++==+=+++-=+=-=+=-+-=+-=+-=+=-=+=--=+++=++=+++==++==--=+=++==+++-")
	case 3:
		track = readTrack("S+= +=-== +=++=     =+=+=--=    =-= ++=     +=-  =+=++=-+==+ =++=-=-=--\n- + +   + =   =     =      =   == = - -     - =  =         =-=        -\n= + + +-- =-= ==-==-= --++ +  == == = +     - =  =    ==++=    =++=-=++\n+ + + =     +         =  + + == == ++ =     = =  ==   =   = =++=       \n= = + + +== +==     =++ == =+=  =  +  +==-=++ =   =++ --= + =          \n+ ==- = + =   = =+= =   =       ++--          +     =   = = =--= ==++==\n=     ==- ==+-- = = = ++= +=--      ==+ ==--= +--+=-= ==- ==   =+=    =\n-               = = = =   +  +  ==+ = = +   =        ++    =          -\n-               = + + =   +  -  = + = = +   =        +     =          -\n--==++++==+=+++-= =-= =-+-=  =+-= =-= =--   +=++=+++==     -=+=++==+++-")
	}

	loops := 10
	if part == 3 {
		loops = 2024
	}

	plans := map[string]string{}
	var names []string
	for _, line := range lines {
		tmp := strings.Split(line, ":")
		names = append(names, tmp[0])
		plans[tmp[0]] = strings.Join(strings.Split(tmp[1], ","), "")
	}

	score := map[string]int{}
	for name, plan := range plans {
		score[name] = sim(track, loops, plan)
	}

	slices.SortFunc(names, func(a, b string) int {
		return score[b] - score[a]
	})

	if part < 3 {
		return strings.Join(names, "")
	} else {
		c := make(chan int)
		plans_ := generatePlans()
		for _, plan := range plans_ {
			go func() {
				c <- sim(track, loops, plan)
			}()
		}

		res := 0
		for range len(plans_) {
			if <-c > score["A"] {
				res++
			}
		}

		return strconv.Itoa(res)
	}

}

func sim(track string, loops int, plan string) int {
	s := 10
	l := len(track)
	rounds := l * loops
	score := 0
	for i := range rounds {
		switch track[i%l] {
		case '+':
			s++
		case '-':
			s--
		default:
			switch plan[i%len(plan)] {
			case '+':
				s++
			case '-':
				s--
			}
		}
		score += s
	}

	return score
}

type pt [2]int

func pop[T any](slice *[]T) T {
	n := len(*slice)
	if n == 0 {
		panic("empty slice")
	}
	back := (*slice)[n-1]
	*slice = (*slice)[:n-1]
	return back
}

func readTrack(track string) string {
	rows := strings.Split(track, "\n")
	R, C := len(rows), len(rows[0])
	t := ""
	start := pt{0, 1}
	q := []pt{start}
	seen := map[pt]bool{pt{0, 0}: true, start: true}

	for len(q) > 0 {
		curr := pop(&q)
		t += string(rows[curr[0]][curr[1]])

		for _, n := range neighbors(curr[0], curr[1], R, C) {
			if rows[n[0]][n[1]] == ' ' {
				continue
			}
			if v := seen[n]; !v {
				q = append(q, n)
				seen[n] = true
			}
		}
	}

	return t + "S"
}

func neighbors(r, c, R, C int) (n []pt) {
	for _, d := range []pt{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		rr, cc := r+d[0], c+d[1]
		if 0 <= rr && rr < R && 0 <= cc && cc < C {
			n = append(n, pt{rr, cc})
		}
	}
	return
}

func generatePlans() []string {
	var plans []string

	var f func(string, int, int, int)
	f = func(curr string, plus, minus, equals int) {
		if plus+minus+equals == 0 {
			plans = append(plans, curr)
		}

		if plus > 0 {
			f(curr+"+", plus-1, minus, equals)
		}

		if minus > 0 {
			f(curr+"-", plus, minus-1, equals)
		}

		if equals > 0 {
			f(curr+"=", plus, minus, equals-1)
		}
	}

	f("", 5, 3, 3)
	return plans
}
