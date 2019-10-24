package main

import (
	"fmt"
)

/*
* 归并排序
* 算法思想：分治的思想，就是用递归先将序列分解成只剩一个元素的子序列，然后逐渐向上进行合并，
	每次合并过程就是将两个内部已排序的子序列进行合并排序

* 特点：平均时间复杂度O(n*logn)，最坏时间复杂度O(n*logn)，额外空间O(n)
	（另外需要一个数组），稳定排序（用了递归就要考虑栈溢出）
*/

func main() {
	nums := []int{0, 9, 2, 4, 6, 7, 1, 3, 8, 5}
	fmt.Println(nums)
	sort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func sort(nums []int, left, right int) {
	tmps := make([]int, right-left+1)
	mergeSort(nums, left, right, tmps)
}

func mergeSort(nums []int, left, right int, tmps []int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(nums, left, mid, tmps)
		mergeSort(nums, mid+1, right, tmps)
		merge(nums, left, mid, right, tmps)
	}
}

func merge(nums []int, left, mid, right int, tmps []int) {
	index := left
	i := left
	j := mid + 1
	for ; i <= mid && j <= right; index++ {
		if nums[i] <= nums[j] {
			tmps[index] = nums[i]
			i += 1
		} else {
			tmps[index] = nums[j]
			j += 1
		}
	}
	for ; i <= mid; index++ {
		tmps[index] = nums[i]
		i += 1
	}
	for ; j <= right; index++ {
		tmps[index] = nums[j]
		j += 1
	}
	for k := left; k <= right; k++ {
		nums[k] = tmps[k]
	}
}
