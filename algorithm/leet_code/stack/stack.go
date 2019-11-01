package main

import (
	"fmt"
)

// 栈(): 后入先出的数据结构
type MinStack struct {
	Tail  int
	Stack []byte
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Tail:  -1,
		Stack: make([]byte, 0),
	}
}

func (this *MinStack) Push(x byte) {
	this.Tail += 1
	if this.Tail < len(this.Stack) {
		this.Stack[this.Tail] = x
	} else {
		this.Stack = append(this.Stack, x)
	}
}

func (this *MinStack) Pop() {
	if this.Tail >= 0 {
		this.Tail -= 1
	}
}

func (this *MinStack) Top() byte {
	if this.Tail >= 0 {
		return this.Stack[this.Tail]
	}
	return 0
}

func (this *MinStack) GetMin() byte {
	var min byte
	if this.Tail >= 0 {
		min = this.Stack[this.Tail]
	}
	for i := 0; i < this.Tail; i++ {
		if this.Stack[i] < min {
			min = this.Stack[i]
		}
	}
	return min
}

func (this *MinStack) IsEmpty() bool {
	return this.Tail == -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	obj := Constructor()
	obj.Push(10)
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
*/
func isValid(s /*[()][)]*/ string) bool {
	if len(s) == 1 {
		return false
	}

	obj := Constructor()
	for i := len(s) - 1; i >= 0; i-- {
		obj.Push(s[i])
	}

	xs := make([]byte, 0)
	for {
		if obj.IsEmpty() {
			break
		}

		v := obj.Top()
		obj.Pop()

		if v == '{' || v == '(' || v == '[' {
			xs = append(xs, v)
		} else if v == '}' || v == ')' || v == ']' {
			if len(xs) == 0 {
				return false
			}
			a := xs[len(xs)-1]
			if v == '}' && a != '{' ||
				v == ')' && a != '(' ||
				v == ']' && a != '[' {
				return false
			}
			xs = xs[:len(xs)-1]
		} else {
			return false
		}
	}
	if len(xs) > 0 {
		return false
	}
	return true
}

// 给定一个二叉树，返回它的中序 遍历。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	array := make([]int, 0)

	ls := inorderTraversal(root.Left)
	array = append(array, ls...)

	array = append(array, root.Val)

	rs := inorderTraversal(root.Right)
	array = append(array, rs...)
	return array
}
