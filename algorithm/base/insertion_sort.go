package main

import (
	"fmt"
)

/*
* 插入排序
* 算法思路：假定一个已排好序的序列和一个元素，只需将该元素从序列末尾向前比较，
	找到第一个小于它的序列元素，排在其之后即可。

* 特点：平均时间复杂度O(n^2)，最坏时间复杂度O(n^2)，额外空间O(1)，
	稳定排序（比较元素和序列时，找到序列中相等元素的话，排在其之后），序列大部分已排好序时
*/

func main() {
	nums := []int{0, 9, 2, 4, 6, 7, 1, 3, 8, 5}
	fmt.Println(nums)
	fmt.Println(insetionSort(nums))
}

func insetionSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		index := i
		for j := i - 1; j >= 0; j-- {
			if nums[index] < nums[j] {
				nums[index], nums[j] = nums[j], nums[index]
				index = j
			} else {
				break
			}
		}
	}
	return nums
}
