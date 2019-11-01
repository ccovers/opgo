package main

import (
	"fmt"
	"math/rand"
)

/*
* AVL树是带有平衡条件的二叉查找树。它要求在AVL树中任何节点的两个子树的高度(高度是指节点到一片树叶的最长路径的长) 最大差别为1。
* AVL树很好的解决了二叉搜索树退化成链表的问题，把插入，查找，删除的时间复杂度的最好情况和最坏情况都维持在O(logn)。
* 当我们执行插入一个节点时，很可能会破坏AVL树的平衡特性，所以我们需要调整AVL树的结构，使其重新平衡，而调整的方式称之为旋转。
* 针对父节点的位置分为左-左，左-右，右-右，右-左这4类情况分析，而对左-左，右-右情况进行单旋转，也就是一次旋转，对左-右，右-左情况进行双旋转，两次旋转，最终恢复其平衡特性。
 */
type TreeNode struct {
	Val   int
	High  int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	slice := make([]int, 10)
	for i := 0; i < 10; i++ {
		slice[i] = rand.Intn(10)
	}
	fmt.Println("初始:", slice)

	// 生成二叉树
	var root *TreeNode
	for _, val := range slice {
		root = newBalanceTree(val, root)
	}

	// 判断是否平衡二叉树
	fmt.Println("是否平衡二叉树:", isBalanced(root), root.Val, root.High)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val, abs(deep(root.Left)-deep(root.Right)))

	// 删除平衡树的指定节点
	root = deleteTreeNode(root, slice[1])
	fmt.Println("是否平衡二叉树:", isBalanced(root), root.Val, root.High)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val, abs(deep(root.Left)-deep(root.Right)))

	root = deleteTreeNode(root, slice[8])
	fmt.Println("是否平衡二叉树:", isBalanced(root), root.Val, root.High)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val, abs(deep(root.Left)-deep(root.Right)))

	root = deleteTreeNode(root, slice[6])
	fmt.Println("是否平衡二叉树:", isBalanced(root), root.Val, root.High)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val, abs(deep(root.Left)-deep(root.Right)))

	root = deleteTreeNode(root, slice[0])
	fmt.Println("是否平衡二叉树:", isBalanced(root), root.Val, root.High)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val, abs(deep(root.Left)-deep(root.Right)))

	root = deleteTreeNode(root, slice[4])
	fmt.Println("是否平衡二叉树:", isBalanced(root), root.Val, root.High)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val, abs(deep(root.Left)-deep(root.Right)))

	root = deleteTreeNode(root, slice[7])
	fmt.Println("是否平衡二叉树:", isBalanced(root), root.Val, root.High)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val, abs(deep(root.Left)-deep(root.Right)))

	root = deleteTreeNode(root, slice[9])
	fmt.Println("是否平衡二叉树:", isBalanced(root), root.Val, root.High)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val, abs(deep(root.Left)-deep(root.Right)))

	root = deleteTreeNode(root, slice[2])
	fmt.Println("是否平衡二叉树:", isBalanced(root), root.Val, root.High)
	//fmt.Println(root.Val, root.Left.Val, root.Right.Val, abs(deep(root.Left)-deep(root.Right)))
}

// 生成平衡二叉树
func newBalanceTree(val int, root *TreeNode) *TreeNode {
	if root == nil {
		root = &TreeNode{
			Val:  val,
			High: 0,
		}
	} else {
		if val < root.Val {
			// 元素比根小，加到左子树
			root.Left = newBalanceTree(val, root.Left)
			// 判断是否要旋转
			if abs(deep(root.Left)-deep(root.Right)) > 1 {
				if val >= root.Left.Val {
					// 比左子根大，加到左子树的右边，转两次，即 LR
					root = LR(root)
				} else {
					// 比左子根小，加到左子树的左边，转一次，即 LL
					root = LL(root)
				}
			}
		} else {
			// 元素比根大，加到右子树
			root.Right = newBalanceTree(val, root.Right)
			// 判断是否要旋转
			if abs(deep(root.Right)-deep(root.Left)) > 1 {
				if val >= root.Right.Val {
					// 比右子根大，加到右子树的右边，转一次，即 RR
					root = RR(root)
				} else {
					// 比右子根小，加到右子树的左边，转两次次，即 RL
					root = RL(root)
				}
			}
		}
	}
	root.High = max(deep(root.Left), deep(root.Right)) + 1
	return root
}

// 左旋转
func LL(root *TreeNode) *TreeNode {
	node := root.Left
	root.Left = node.Right
	node.Right = root
	root.High = max(deep(root.Left), deep(root.Right)) + 1
	return node
}

// 右旋转
func RR(root *TreeNode) *TreeNode {
	node := root.Right
	root.Right = node.Left
	node.Left = root
	root.High = max(deep(root.Left), deep(root.Right)) + 1
	return node
}

// 右边子节点左旋转再右旋转
func RL(root *TreeNode) *TreeNode {
	root.Right = LL(root.Right)
	root = RR(root)
	return root
}

// 左边子节点右旋转再左旋转
func LR(root *TreeNode) *TreeNode {
	root.Left = RR(root.Left)
	root = LL(root)
	return root
}

// 最大值
func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

// 树深度
func deep(root *TreeNode) int {
	if root == nil {
		return -1
	} else {
		return root.High
	}
}

// 返回绝对值
func abs(val int) int {
	if val >= 0 {
		return val
	} else {
		return -val
	}
}

// 删除平衡树节点
func deleteTreeNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val == root.Val {
		if root.Left != nil && root.Right != nil {
			root.Val = findMax(root.Left)
			root.Left = deleteTreeNode(root.Left, root.Val)
		} else if root.Left != nil {
			root = root.Left
		} else {
			root = root.Right
		}
	} else if val < root.Val {
		root.Left = deleteTreeNode(root.Left, val)
	} else {
		root.Right = deleteTreeNode(root.Right, val)
	}
	if root == nil {
		return nil
	}
	root.High = max(deep(root.Left), deep(root.Right)) + 1

	if getBalance(root) > 1 {
		if getBalance(root.Left) >= 0 {
			root = LL(root)
		} else {
			root = LR(root)
		}
	} else if getBalance(root) < -1 {
		if getBalance(root.Right) >= 0 {
			root = RL(root)
		} else {
			root = RR(root)
		}
	}
	return root
}

// 找到树节点的最大值
func findMax(root *TreeNode) int {
	if root.Right == nil {
		return root.Val
	} else {
		return findMax(root.Right)
	}
}

// 获取平衡树的平衡因子（左右两边深度差值）
func getBalance(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return deep(root.Left) - deep(root.Right)
}

// 判断是否平衡树
func isBalanced(root *TreeNode) bool {
	_, flag := childRootCnt(root)
	return flag

}

func childRootCnt(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	lcnt, lflag := childRootCnt(root.Left)
	if !lflag {
		return lcnt + 1, lflag
	}
	rcnt, rflag := childRootCnt(root.Right)
	if !rflag {
		return rcnt + 1, rflag
	}

	flag := true
	if lcnt > rcnt {
		if lcnt > rcnt+1 {
			flag = false
		}
		return lcnt + 1, flag
	} else {
		if rcnt > lcnt+1 {
			flag = false
		}
		return rcnt + 1, flag
	}
}
