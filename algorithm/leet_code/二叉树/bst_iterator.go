package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Root *TreeNode
	Node *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		Root: root,
		Node: nil,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node := travels(this.Root, this.Node)
	if node != nil {
		this.Node = node
		return node.Val
	} else {
		return this.Node.Val
	}
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return travels(this.Root, this.Node) != nil
}

func travels(root *TreeNode, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if node != nil {
		if root.Val > node.Val {
			lnode := travels(root.Left, node)
			if lnode != nil {
				return lnode
			}
		} else {
			return travels(root.Right, node)
		}
	} else {
		lnode := travels(root.Left, node)
		if lnode != nil {
			return lnode
		}
	}
	return root
}
