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

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return false
	}

	if s.Val == t.Val {
		if checkSubTree(s, t) {
			return true
		}
	}
	if isSubtree(s.Left, t) {
		return true
	}
	if isSubtree(s.Right, t) {
		return true
	}
	return false
}

func checkSubTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if (s != nil && t == nil) || (s == nil && t != nil) {
		return false
	}

	if s.Val != t.Val {
		return false
	}
	if !checkSubTree(s.Left, t.Left) {
		return false
	}
	if !checkSubTree(s.Right, t.Right) {
		return false
	}
	return true
}
