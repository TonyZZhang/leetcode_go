package leetcode

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var in []int
	in = append(in, 1)
	in = append(in, 5)
	in = append(in, 2)
	in = append(in, 4)
	in = append(in, 67)
	in = append(in, 3)
	fmt.Println(selectSort(in))
}

func selectSort(x []int) []int {
	for i := 0; i < len(x); i++ {
		min := x[i]
		needSwap := i
		for j := i; j < len(x); j++ {
			if x[j] < min {
				min = x[j]
				needSwap = j
			}
		}
		minX := x[needSwap]
		x[needSwap] = x[i]
		x[i] = minX

	}
	return x
}
