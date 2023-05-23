package leetcode

///  十进制    二进制
//0        0
//1        1
//2        10
//3        11

//规律：奇数比其小于1（十进制）的偶数多一个1（二进制）
//     偶数跟其除以2的1的个数相同
func countBits(n int) []int {
	res := make([]int, n+1)
	res[0] = 0
	for i := 1; i < n+1; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i-1] + 1
		}
	}
	return res
}
