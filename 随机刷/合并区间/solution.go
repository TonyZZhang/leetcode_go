package leetcode

func merge(intervals [][]int) [][]int {
	r := make([][]int, 0)
	max := 0
	for _, v := range intervals {
		if v[1] > max {
			max = v[1]
		}
	}
	res := make([]int, max+1)

	for _, v := range intervals {
		for i := v[0]; i < v[1]; i++ {
			res[i] = 1
		}
	}

	var start, end  int

	if res[0] == 1 {
		start = 0
	}

	for k, _ := range res {
		if k-1 >= 0 {
			if res[k-1] == 0 && res[k] == 1 {
				start = k
			}
		}


		if k+1 < len(res) {
			if res[k] == 1 && res[k+1] == 0 {
				end = k+1
				if start != -1 {
					temp := []int{start, end}
					r = append(r, temp)
				}
				start = -1
			}
		}
	}

	if res[len(res)-1] == 1 && start != -1 {
		end = len(res) - 1
		temp := []int{start, end}
		r = append(r, temp)
	}

	return r
}
