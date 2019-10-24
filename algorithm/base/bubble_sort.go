package main

import (
	"fmt"
)

/*
* 冒泡排序
* 算法原理：比较两个相邻的元素，将值大的元素右移。
* 算法思路：首先第一个元素和第二个元素比较，如果第一个大，则二者交换，否则不交换；
		然后第二个元素和第三个元素比较，如果第二个大，则二者交换，否则不交换……一直执行下去，
		最终最大的那个元素被交换到最后，一趟冒泡排序完成。最坏的情况是排序是逆序的。

* 特点：平均时间复杂度O(n^2)，最坏时间复杂度O(n^2)，额外空间O(1)，稳定排序
*/

func main() {
	nums := []int{0, 9, 2, 4, 6, 7, 1, 3, 8, 5}
	fmt.Println(nums)
	fmt.Println(bubbleSort(nums))
}

func bubbleSort(nums []int) []int {
	length := len(nums)

	for i := 0; i < length; i++ {
		flag := false
		for j := i + 1; j < length; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return nums
}
