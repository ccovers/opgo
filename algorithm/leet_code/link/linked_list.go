package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Pre  *Node
	Next *Node
}

type MyLinkedList struct {
	Head *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		Head: nil,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if this.Head == nil || index < 0 {
		return -1
	}

	tmp := this.Head
	for i := 0; i < index; i++ {
		if tmp.Next == nil {
			return -1
		}
		tmp = tmp.Next
	}
	return tmp.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		Val:  val,
		Next: this.Head,
		Pre:  nil,
	}
	if this.Head != nil {
		this.Head.Pre = node
	}
	this.Head = node
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{
		Val:  val,
		Next: nil,
		Pre:  nil,
	}

	if this.Head == nil {
		this.Head = node
	} else {
		tmp := this.Head
		for {
			if tmp.Next != nil {
				tmp = tmp.Next
			} else {
				node.Pre = tmp
				tmp.Next = node
				break
			}
		}
	}

}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	node := &Node{
		Val:  val,
		Next: nil,
		Pre:  nil,
	}

	if index <= 0 {
		node.Next = this.Head
		if this.Head != nil {
			this.Head.Pre = node
		}
		this.Head = node
	} else if this.Head != nil {
		tmp := this.Head
		for i := 1; i < index; i++ {
			if tmp.Next != nil {
				tmp = tmp.Next
			} else {
				return
			}
		}
		node.Pre = tmp
		node.Next = tmp.Next
		if node.Next != nil {
			tmp.Next.Pre = node
		}
		tmp.Next = node
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.Head == nil || index < 0 {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
		if this.Head != nil {
			this.Head.Pre = nil
		}
	} else {
		tmp := this.Head
		for i := 1; i < index; i++ {
			if tmp.Next != nil {
				tmp = tmp.Next
			} else {
				return
			}
		}
		if tmp.Next != nil {
			tmp.Next = tmp.Next.Next
			if tmp.Next != nil {
				tmp.Next.Pre = tmp
			}
		}
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func main() {

}
