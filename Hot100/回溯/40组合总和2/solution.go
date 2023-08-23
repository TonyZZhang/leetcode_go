package leetcode

func combinationSumV2(candidates []int, target int) [][]int {
	re := [][]int{}
	p := []int{}
	databackup(candidates, target, p, &re)
	return re
}

func databackup(candidates []int, target int, path []int, re *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*re = append(*re, path)
	}

	for i := 0; i < len(candidates); i++ {
		if contains(path, candidates[i]) {
			continue
		}
		databackup(candidates[i:], target-candidates[i], append(path, candidates[i]), re)
	}
}

func contains(c []int, target int) bool {
	for i := 0; i < len(c); i++ {
		if c[i] == target {
			return true
		}
	}
	return false
}

