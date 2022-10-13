package leetcode

import (
	"fmt"
	"testing"
)

type Node struct {
}

func TestName(t *testing.T) {
	var nodeQueue []*Node
	nodeQueue = append(nodeQueue, nil)
	fmt.Println(len(nodeQueue))
}
