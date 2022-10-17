package leetcode

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var in []int
	in = append(in, 1)
	in = append(in, 3)
	in = append(in, 2)
	fmt.Println(sort(in))
}
