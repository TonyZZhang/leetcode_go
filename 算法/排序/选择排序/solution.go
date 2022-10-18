package leetcode

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
