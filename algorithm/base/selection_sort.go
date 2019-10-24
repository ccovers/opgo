package main

import (
	"fmt"
)

/*
* 选择排序
* 基本思想：选择排序，从头至尾扫描序列，找出最小的一个元素，和第一个元素交换。
			接着从剩下的元素中继续这种选择和交换方式，最终得到一个有序序列。

* 特点：平均时间复杂度O(n^2)，最坏时间复杂度O(n^2)，额外空间O(1)，不稳定排序.
*/

func main() {
	nums := []int{0, 9, 2, 4, 6, 7, 1, 3, 8, 5}
	fmt.Println(nums)
	fmt.Println(selectionSort(nums))
}

func selectionSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		index := i
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[index] {
				index = j
			}
		}

		if index != i {
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	return nums
}
