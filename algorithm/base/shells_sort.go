package main

import (
	"fmt"
)

/*
* 希尔排序
* 算法思想：希尔算法又名缩小增量排序，本质是插入排序，只不过是将待排序的序列按某种规则分成几个子序列，
	分别对几个子序列进行直接插入排序。这个规则就是增量，增量选取很重要，增量一般选序列长度一半，
	然后逐半递减，直到最后一个增量为1，为1相当于直接插入排序。

* 特点：平均时间复杂度O(n*logn)，最坏时间复杂度O(n^s)(1<s<2)，额外空间O(1)，不稳定排序
*/
func main() {
	nums := []int{0, 9, 2, 4, 6, 7, 1, 3, 8, 5}
	fmt.Println(nums)
	fmt.Println(shellsSort(nums, len(nums)))
}

func shellsSort(nums []int, nlen int) []int {
	length := len(nums)

	for ; nlen > 0; nlen /= 2 {
		for i := 0; i < length; i++ {
			index := i
			for j := i - nlen; j > 0; j -= nlen {
				if nums[index] < nums[j] {
					nums[index], nums[j] = nums[j], nums[index]
					index = j
				}
			}
		}
	}
	return nums
}
