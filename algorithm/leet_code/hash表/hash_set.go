package main

import (
	"fmt"
)

type Node struct {
	Key  int
	Next *Node
}

type MyHashSet struct {
	Buckets []*Node
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		Buckets: make([]*Node, 10000),
	}
}

func hash(key int) int {
	return key % 10000
}

func (this *MyHashSet) Add(key int) {
	head := this.Buckets[hash(key)]
	if head != nil {
		for head != nil {
			if head.Key == key {
				break
			}

			if head.Next != nil {
				head = head.Next
			} else {
				head.Next = &Node{
					Key:  key,
					Next: nil,
				}
				break
			}
		}
	} else {
		this.Buckets[hash(key)] = &Node{
			Key:  key,
			Next: nil,
		}
	}

}

func (this *MyHashSet) Remove(key int) {
	head := this.Buckets[hash(key)]
	if head != nil && head.Key == key {
		this.Buckets[hash(key)] = head.Next
	} else {
		for head != nil {
			if head.Next != nil {
				if head.Next.Key == key {
					head.Next = head.Next.Next
					break
				}
			}
			head = head.Next
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	head := this.Buckets[hash(key)]
	for head != nil {
		if head.Key == key {
			return true
		}
		head = head.Next
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
