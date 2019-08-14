package main

import (
	"fmt"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: nums[len(nums)/2],
	}

	if len(nums) > 1 {
		root.Left = sortedArrayToBST(nums[0 : len(nums)/2])
	}
	if len(nums) > len(nums)/2+1 {
		root.Right = sortedArrayToBST(nums[len(nums)/2+1:])
	}

	return root
}

func main() {
	tm, _ := time.Parse("2006-01-02", "2009-09-09")
	fmt.Printf("%s,%d\n", tm.String(), tm.YearDay())
}
