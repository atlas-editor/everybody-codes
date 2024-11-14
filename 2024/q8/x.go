package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile(os.Args[1])
	input := strings.TrimSpace(string(data))
	num := atoi(input)

	switch os.Args[2] {
	case "1":
		fmt.Println(p1(num))
	case "2":
		fmt.Println(p2(num))
	case "3":
		fmt.Println(p3(num))
	}
}

func p1(num int) int {
	i := int(math.Ceil(math.Sqrt(float64(num))))
	return (i*i - num) * (2*i - 1)

}

func p2(num int) int {
	size, thickness := 1, 1
	i := 3
	for {
		thickness = (thickness * num) % 1111
		size += i * thickness

		if size > 20240000 {
			return i * (size - 20240000)
		}
		i += 2
	}
}

func p3(num int) int {
	i, j := 0, 200000
	for i < j {
		mid := (i + j) >> 1
		if blocks(mid, num, 10) < 202400000000 {
			i = mid + 1
		} else {
			j = mid
		}
	}

	return blocks(i, num, 10) - 202400000000
}

func blocks(layers int, highPriests int, acolytes int) int {
	thickness := 1
	thicknessSeq := []int{1}
	width := 1
	H := 1
	for i := 2; i < layers; i++ {
		thickness = (thickness*highPriests)%acolytes + acolytes
		thicknessSeq = append(thicknessSeq, thickness)
		H += thickness
		width = 2*i - 1
	}

	columnHeights := make([]int, width)
	columnHeights[width/2] = H
	columnMins := make([]int, width)
	columnMins[width/2] = 2
	blocksNeeded := H

	for i := 1; i < width/2+1; i++ {
		j0 := width/2 + i
		j1 := width/2 - i

		last := columnHeights[j0-1]
		columnHeights[j0] = last - thicknessSeq[i-1]
		columnHeights[j1] = last - thicknessSeq[i-1]

		columnMins[j0] = thicknessSeq[i] + 1
		columnMins[j1] = thicknessSeq[i] + 1

		blocksNeeded += 2 * (last - thicknessSeq[i-1])
	}

	for i := range width {
		toRemoveI := (highPriests * width * columnHeights[i]) % acolytes

		canRemove := columnHeights[i] - columnMins[i]
		if canRemove < toRemoveI && canRemove > 0 {
			fmt.Println(canRemove, toRemoveI)
		}
		removing := max(0, min(toRemoveI, canRemove))
		blocksNeeded -= removing
	}

	return blocksNeeded
}

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
