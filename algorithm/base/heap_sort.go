package main

import (
	"fmt"
)

/*
* 堆排序

* 算法思路：堆是一种数据结构，可以把堆看成一棵完全二叉树，这棵完全二叉树满足：
	任何一个非叶子结点的值都不大于（或小于）其左右孩子结点的值。
	如父亲大孩子小于大顶堆，若父亲小孩子大于小顶堆。算法描述如下步骤：
1）从无序序列所确定的完全二叉树的第一个非叶子结点开始，从右至左，从下至上，对每个结点进行调整，得到大顶堆
2）将大顶堆变化得到的无序元素的第一个元素与最后一个元素交换位置。再次将剩余元素堆排序
3）重复上述步骤，直到树中仅含一个结点，结束。

* 特点：平均时间复杂度O(n*logn)，最坏时间复杂度O(n*logn)，额外空间O(1)，
	不稳定排序（涉及根节点与最后节点的交换，可能会破坏两相等元素的相对位置！），当n较大时较好（海量数据）！
*/
var nums []int

func main() {
	nums = []int{0, 9, 2, 4, 6, 7, 1, 3, 8, 5}
	fmt.Println(nums)
	sort()
	fmt.Println(nums)
}

func sort() {
	length := len(nums)
	// 从无序序列所确定的完全二叉树的第一个非叶子结点开始，从右至左，从下至上，对每个结点进行调整，得到大顶堆
	for i := length - 1; i >= 0; i-- {
		heapSort(i, length)
	}
	for i := length - 1; i >= 0; i-- {
		// 将大顶堆变化得到的无序元素的第一个元素与最后一个元素交换位置。再次将剩余元素堆排序
		nums[i], nums[0] = nums[0], nums[i]
		heapSort(0, i)
	}
}

func heapSort(index, length int) {
	left := index*2 + 1
	right := index*2 + 2

	tmp := -1
	if left < length && nums[left] > nums[index] {
		tmp = left
	}
	if right < length && nums[right] > nums[index] {
		if tmp >= 0 {
			if nums[right] > nums[tmp] {
				tmp = right
			}
		} else {
			tmp = right
		}
	}
	if tmp >= 0 {
		nums[tmp], nums[index] = nums[index], nums[tmp]
		heapSort(tmp, length)
	}
}
