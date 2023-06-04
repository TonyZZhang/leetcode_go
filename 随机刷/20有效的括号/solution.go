package leetcode
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//"(([]){})"
//({)[}]
func isValid(s string) bool {
	if len(s) % 2 != 0 || len(s) < 2{
		return false
	}

	if len(s) == 2 {
		return isCorrect(s[0], s[1])
	}

	mid := len(s)/2 - 1
	for mid >= 0 {
		left := s[mid]
		right := s[mid+1]
		if isCorrect(left, right) {
			newStr := s[0:mid]+s[mid+2:]
			return isValid(newStr)
		}
		return isValid(s[:mid+1]) && isValid(s[mid+1:])
	}
	return false
}

func isCorrect(left, right uint8) bool {
	if left == '{' {
		return right == '}'
	}
	if left == '(' {
		return right == ')'
	}
	if left == '[' {
		return right == ']'
	}
	return false
}


