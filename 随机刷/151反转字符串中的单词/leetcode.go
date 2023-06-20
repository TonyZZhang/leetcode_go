package leetcode

func reverseWords(s string) string {
	x := make([][]uint8, len(s))
	k := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' '{
			k++
			continue
		}
		x[k] = append(x[k], s[i])
	}

	var str string
	var start int
	for i := len(x)-1; i >= 0 ; i-- {
		if len(x[i]) != 0 {
			start = i
			break
		}
	}

	for i := start; i >= 0 ; i-- {
		if len(x[i]) == 0 {
			str += " "
		} else {
			str += string(x[i]) + " "
		}
	}
	return str
}

