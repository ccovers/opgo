package main

import (
	"fmt"
)

type Node struct {
	Key   int
	Value int
	Next  *Node
}

type MyHashMap struct {
	Buckets []*Node
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		Buckets: make([]*Node, 10000),
	}
}

func hash(key int) int {
	return key % 10000
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	head := this.Buckets[hash(key)]
	if head != nil {
		for head != nil {
			if head.Key == key {
				head.Value = value
				break
			}

			if head.Next != nil {
				head = head.Next
			} else {
				head.Next = &Node{
					Key:   key,
					Value: value,
					Next:  nil,
				}
				break
			}
		}
	} else {
		this.Buckets[hash(key)] = &Node{
			Key:   key,
			Value: value,
			Next:  nil,
		}
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	head := this.Buckets[hash(key)]
	for head != nil {
		if head.Key == key {
			return head.Value
		}
		head = head.Next
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
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

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
