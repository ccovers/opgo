package main

import (
	"fmt"
)

/*
* 快速排序
* 算法原理：以一个数为中心将序列分成两个部分，一边全是比它小，另一边全是比它大。

*特点：平均时间复杂度O(n*logn)，最坏时间复杂度O(n^2)（序列基本有序时，退化为冒泡排序），额外空间O(logn)
	，不稳定排序，当n较大时较好（当也不能太大，用了递归就要考虑栈溢出）！
*/

var nums []int

func main() {
	nums = []int{0, 9, 2, 4, 6, 7, 1, 3, 8, 5}
	fmt.Println(nums)
	quick_sort(0, len(nums)-1)
	fmt.Println(nums)
}

func quick_sort(left, right int) {
	if left < right {
		mid := partion(left, right)
		quick_sort(left, mid-1)
		quick_sort(mid+1, right)
	}
}

func partion(left, right int) int {
	for left < right {
		for ; left < right; right-- {
			if nums[left] >= nums[right] {
				nums[left], nums[right] = nums[right], nums[left]
				break
			}
		}

		for ; left < right; left++ {
			if nums[left] >= nums[right] {
				nums[left], nums[right] = nums[right], nums[left]
				break
			}
		}
	}
	return left
}
