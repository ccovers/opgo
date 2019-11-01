package main

import (
	"fmt"
	"math/rand"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	slice := make([]int, 5)
	for i := 0; i < 5; i++ {
		slice[i] = rand.Intn(5)
	}
	fmt.Println("初始:", slice)

	// 生成二叉树
	root := constructor(0, slice)

	// 前序遍历
	fmt.Println("前序遍历:", preorderTraversal(root))

	// 中序遍历
	fmt.Println("中序遍历:", inorderTraversal(root))

	// 后序遍历
	fmt.Println("后序遍历:", postorderTraversal(root))

	// 层序遍历
	fmt.Println("层次遍历:", levelOrder(root))
}

// 生成二叉树
func constructor(index int, slice []int) *TreeNode {
	if index >= len(slice) {
		return nil
	}

	return &TreeNode{
		Val:   slice[index],
		Left:  constructor(index*2+1, slice),
		Right: constructor(index*2+2, slice),
	}
}

// 前序遍历（根节点排最先，然后同级先左后右）
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	slice := make([]int, 0)

	slice = append(slice, root.Val)
	if root.Left != nil {
		slice = append(slice, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		slice = append(slice, preorderTraversal(root.Right)...)
	}

	return slice
}

// 中序遍历（先左后根最后右）
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	slice := make([]int, 0)

	if root.Left != nil {
		slice = append(slice, inorderTraversal(root.Left)...)
	}
	slice = append(slice, root.Val)
	if root.Right != nil {
		slice = append(slice, inorderTraversal(root.Right)...)
	}

	return slice
}

// 后序遍历（先左后右最后根）
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	slice := make([]int, 0)

	if root.Left != nil {
		slice = append(slice, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		slice = append(slice, postorderTraversal(root.Right)...)
	}
	slice = append(slice, root.Val)

	return slice
}

// 广度优先搜索时，按照层序遍历顺序
func levelOrder(root *TreeNode) [][]int {
	slice := make([][]int, 0)
	return levelOrderEx(root, 0, slice)
}

func levelOrderEx(root *TreeNode, index int, slice [][]int) [][]int {
	if root == nil {
		return slice
	}
	if len(slice) <= index {
		slice = append(slice, make([]int, 0))
	}
	slice[index] = append(slice[index], root.Val)

	slice = levelOrderEx(root.Left, index+1, slice)
	slice = levelOrderEx(root.Right, index+1, slice)
	return slice
}
