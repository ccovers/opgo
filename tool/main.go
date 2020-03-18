package main

import (
	"fmt"
)

func main() {
	fmt.Println(handle(12, map[int]int{}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {

}
