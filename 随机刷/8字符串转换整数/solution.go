package leetcode

func myAtoi(s string) int {
	if len(s) == 0 {
		return -1
	}
	res := 0
	tenNum := 1
	for i := len(s)-1; i >= 0 ; i-- {
		res = res + tenNum* int((s[i]-'0'))
		tenNum = tenNum * 10
	}
	return res
}
