package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	两个单向链表l1、l2，分别保存的一串单个数值，将l1与l2相加
	超过一位数值的移到上一节点
	除了只有一只节点的，其余的第一节点不应为0
*/

func main() {
	a := make([]int, 0)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	fmt.Printf("%+v", len(a[1:]))
	return

	l1 := &ListNode{
		Val:  7,
		Next: nil,
	}
	l1.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	l1.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}
	l1.Next.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}

	l2 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l2.Next = &ListNode{
		Val:  6,
		Next: nil,
	}
	l2.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}

	l := addTwoNumbers(l1, l2)
	for t := l; t != nil; t = t.Next {
		fmt.Printf("%d ", t.Val)
	}
	fmt.Println("")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	deep1 := 0
	for t := l1; t != nil; t = t.Next {
		deep1 += 1
	}
	deep2 := 0
	for t := l2; t != nil; t = t.Next {
		deep2 += 1
	}

	var deep int
	if deep1 > deep2 {
		deep = deep1
	} else {
		deep = deep2
	}
	fmt.Println(deep, deep1, deep2)

	var remain int
	var node *ListNode
	for i := 0; i < deep; i++ {
		var v1 int
		var v2 int
		if deep1 > 0 {
			v1 = getDeepValue(l1, deep1)
			deep1 -= 1
		}
		if deep2 > 0 {
			v2 = getDeepValue(l2, deep2)
			deep2 -= 1
		}
		remain += v1 + v2
		fmt.Println("remain=", remain)

		t := &ListNode{
			Val:  remain % 10,
			Next: node,
		}
		remain = remain / 10
		node = t
	}

	if node == nil {
		node = &ListNode{
			Val:  remain,
			Next: nil,
		}
	} else {
		if remain > 0 {
			t := &ListNode{
				Val:  remain,
				Next: node,
			}
			node = t
		}
	}
	return node
}

func getDeepValue(l *ListNode, deep int) int {
	if l == nil {
		return 0
	}

	for i := 1; i < deep; i++ {
		l = l.Next
		if l == nil {
			return 0
		}
	}
	return l.Val
}
