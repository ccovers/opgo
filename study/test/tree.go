package main

import (
	"fmt"
	"time"
)

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}

	ch <- t.Value

	Walk(t.Left, ch)
	Walk(t.Right, ch)
}

func Same(t1, t2 *Tree) {
	ch1 := make(chan int, 0)
	ch2 := make(chan int, 0)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var num1 int
	var num2 int
	for true {
		num1 = -1
		num2 = -1

		select {
		case num1 = <-ch1:
		case <-time.After(time.Second * 1):
		}
		select {
		case num2 = <-ch2:
		case <-time.After(time.Second * 1):
		}

		if num1 < 0 || num2 < 0 {
			if num1 != num2 {
				fmt.Printf("error: num1=%d, num2=%d\n", num1, num2)
			}
			break
		}
		if num1 != num2 {
			fmt.Printf("not equal: num1=%d, num2=%d\n", num1, num2)
			break
		}
	}
}

func getTree(nums []int) *Tree {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2

	t := &Tree{
		Value: nums[mid],
		Left:  nil,
		Right: nil,
	}

	if mid > 0 {
		t.Left = getTree(nums[0:mid])
	}
	if len(nums) > mid {
		t.Right = getTree(nums[mid+1 : len(nums)])
	}
	return t
}

func main() {
	/*t1 := getTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	t2 := getTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	Same(t1, t2)
	*/
	fmt.Println("over")

}
