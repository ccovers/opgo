package main

import (
	"fmt"
)

// 队列(queue): 先入先出的数据结构
type MyCircularQueue struct {
	Head     int
	Tail     int
	Size     int
	Capacity int
	Queue    []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Head:     -1,
		Tail:     -1,
		Size:     0,
		Capacity: k,
		Queue:    make([]int, k),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.Head = 0
	}
	this.Tail = (this.Tail + 1) % this.Capacity

	this.Queue[this.Tail] = value
	this.Size += 1
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	if this.Head == this.Tail {
		this.Head = -1
		this.Tail = -1
	} else {
		this.Head = (this.Head + 1) % this.Capacity
	}
	this.Size -= 1
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Queue[this.Head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Queue[this.Tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.Head == -1
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return (this.Tail+1)%this.Capacity == this.Head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
func main() {
	obj := Constructor(3)
	param_1 := obj.EnQueue(1)
	fmt.Println(param_1)
	param_2 := obj.DeQueue()
	fmt.Println(param_2)
	param_3 := obj.Front()
	fmt.Println(param_3)
	param_4 := obj.Rear()
	fmt.Println(param_4)
	param_5 := obj.IsEmpty()
	fmt.Println(param_5)
	param_6 := obj.IsFull()
	fmt.Println(param_6)
}

// 完全平方数
// 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
func numSquares(n int) int {
	slice := make([]int, 0)
	for i := n; i > 0; i-- {
		slice = append(slice, i*i)
	}

	step := 0
	nodeMap := make(map[int]bool)
	queue := Constructor(n)
	queue.EnQueue(n)
	nodeMap[n] = true
	for !queue.IsEmpty() {
		step += 1
		size := queue.Size
		for ; size > 0; size-- {
			tmp := queue.Front()
			queue.DeQueue()
			for _, k := range slice {
				if k == tmp {
					return step
				} else if k < tmp {
					_, ok := nodeMap[tmp-k]
					if !ok {
						queue.EnQueue(tmp - k)
						nodeMap[n] = true
					}
				}
			}
		}
	}
	return step
}
