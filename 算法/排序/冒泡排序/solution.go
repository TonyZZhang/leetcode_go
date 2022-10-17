package leetcode

func sort(input []int) []int {
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input)-1; j++ {
			if input[j] < input[j+1] {
				big := input[j+1]
				input[j+1] = input[j]
				input[j] = big
			}
		}
	}
	return nil
}
