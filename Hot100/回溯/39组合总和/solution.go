package leetcode

func combinationSumV2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	backtrackV2(candidates, target, []int{}, &result)
	return result
}

func backtrackV2(candidates []int, target int, path []int, re *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*re = append(*re, temp)
	}
	for i:=0;i<len(candidates);i++{
		backtrackV2(candidates[i:], target-candidates[i], append(path, candidates[i]), re)
	}
}
