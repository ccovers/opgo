package main

import (
	"fmt"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	nums := make([]int, 0)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	return generateTreesEx(nums)
}

func generateTreesEx(nums []int) []*TreeNode {
	if len(nums) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, 0)

	for _, n := range nums {
		leftNums, rightNums := getNums(n, nums)
		leftNodes := generateTreesEx(leftNums)
		rightNodes := generateTreesEx(rightNums)

		if len(leftNodes) > 0 && len(rightNodes) > 0 {
			for _, leftNode := range leftNodes {
				for _, rightNode := range rightNodes {
					nodes = append(nodes, &TreeNode{
						Val:   n,
						Left:  leftNode,
						Right: rightNode,
					})
				}
			}
		} else if len(leftNodes) > 0 {
			for _, leftNode := range leftNodes {
				nodes = append(nodes, &TreeNode{
					Val:  n,
					Left: leftNode,
				})
			}
		} else if len(rightNodes) > 0 {
			for _, rightNode := range rightNodes {
				nodes = append(nodes, &TreeNode{
					Val:   n,
					Right: rightNode,
				})
			}
		} else {
			nodes = append(nodes, &TreeNode{
				Val: n,
			})
		}
	}
	return nodes
}

func getNums(n int, nums []int) ([]int, []int) {
	leftNums := make([]int, 0)
	rightNums := make([]int, 0)
	for _, num := range nums {
		if num < n {
			leftNums = append(leftNums, num)
		} else if num > n {
			rightNums = append(rightNums, num)
		}
	}
	return leftNums, rightNums
}
