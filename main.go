package main

import (
	"fmt"
)

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getNode(pos int, nums []int) *TreeNode {
	if pos < len(nums) && nums[pos] != 0xFFFFFFFF {
		return &TreeNode{
			Val:   nums[pos],
			Left:  getNode(2*pos+1, nums),
			Right: getNode(2*pos+2, nums),
		}
	}
	return nil
}

func main() {
	node := getNode(0, []int{1, 2, 3, 4, 0xFFFFFFFF, 0xFFFFFFFF, 5})
	x := maxDepth(node)
	fmt.Println("===", x)
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	x := maxDepth(root.Left)
	y := maxDepth(root.Right)
	if x > y {
		return x + 1
	}
	return y + 1
}
