package leetcode

func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) < 2 {
		return []int{}
	}
	right := len(temperatures) - 1
	result := make([]int, 0)
	result = append(result, 0)
	stackIndex := make([]int, 0)
	stackIndex  = append(stackIndex, right)
	mystack := make([]int, 0)
	mystack = append(mystack, temperatures[right])
	right = right - 1
	for right >= 0 {
		if temperatures[right] > mystack[len(mystack)-1] {
			mystack[len(mystack)-1] = temperatures[right]
			stackIndex[len(mystack)-1] = right
			result = append(result, 0)
		} else {
			mystack = append(mystack, temperatures[right])
			result = append(result, stackIndex[len(stackIndex)-1]-right)
			stackIndex = append(stackIndex, right)
		}
		right--
	}
	realRes := make([]int, 0)

	for k := len(result)-1; k > 0; k--  {
		realRes = append(realRes, result[k])
	}

	return realRes
}
