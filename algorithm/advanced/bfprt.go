package main

import (
	"fmt"
)

/*
* BFPRT(线性查找算法) --- 未理解
* 在一大堆数中求其前k大或前k小的问题
* BFPRT算法步骤如下：
	选取主元；
	1.1. 将n个元素按顺序分为$⌊\frac n5⌋$个组，每组5个元素，若有剩余，舍去；
	1.2. 对于这$⌊\frac n5⌋$个组中的每一组使用插入排序找到它们各自的中位数；
	1.3. 对于 1.2 中找到的所有中位数，调用BFPRT算法求出它们的中位数，作为主元；
	以 1.3 选取的主元为分界点，把小于主元的放在左边，大于主元的放在右边；
	判断主元的位置与k的大小，有选择的对左边或右边递归。
*/

func main() {
	nums := []int{0, 9, 2, 4, 6, 7, 1, 3, 8, 5}
	fmt.Println(nums)
	BFPRT(nums, 0, 4, 3)
	fmt.Println(nums)
}

/**
 * 返回数组 nums[left, right] 的第 k 小数的下标
 */
func BFPRT(nums []int, left, right, k int) int {
	// 得到中位数的中位数下标（即主元下标）
	pivotIndex := getPivotIndex(nums, left, right)

	// 进行划分，返回划分边界
	partitionIndex := partition(nums, left, right, pivotIndex)
	num := partitionIndex - left + 1

	if num == k {
		return partitionIndex
	} else if num > k {
		return BFPRT(nums, left, partitionIndex-1, k)
	} else {
		return BFPRT(nums, partitionIndex+1, right, k-num)
	}
}

/**
 * 数组 nums[left, right] 每五个元素作为一组，并计算每组的中位数，
 * 最后返回这些中位数的中位数下标（即主元下标）。
 *
 * @attention 末尾返回语句最后一个参数多加一个 1 的作用其实就是向上取整的意思，
 * 这样可以始终保持 k 大于 0。
 */
func getPivotIndex(nums []int, left, right int) int {
	if right-left < 5 {
		return insertSort(nums, left, right)
	}

	subRight := left - 1
	for i := left; i+4 <= right; i += 5 {
		index := insertSort(nums, i, i+4)
		subRight += 1
		nums[subRight], nums[index] = nums[index], nums[subRight]
	}
	return BFPRT(nums, left, subRight, ((subRight-left+1)>>1)+1)
}

// 对数组nums[left, right]进行插入排序，并返回[left, right]的中位数
func insertSort(nums []int, left, right int) int {
	for i := left; i <= right; i++ {
		index := i
		for j := i - 1; j >= left; j-- {
			if nums[j] > nums[index] {
				nums[j], nums[index] = nums[index], nums[j]
				index = j
			} else {
				break
			}
		}
	}
	return (right-left)>>1 + left
}

/**
 * 利用主元下标 pivot_index 进行对数组 nums[left, right] 划分，并返回
 * 划分后的分界线下标。
 */
func partition(nums []int, left, right, pivotIndex int) int {
	// 把主元放置于末尾
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]

	// 跟踪划分的分界线
	partitionIndex := left
	for i := left; i < right; i++ {
		if nums[i] < nums[right] {
			// 比主元小的都放在左侧
			nums[partitionIndex], nums[i] = nums[i], nums[partitionIndex]
			partitionIndex += 1
		}
	}
	// 最后把主元换回来
	nums[pivotIndex], nums[right] = nums[pivotIndex], nums[right]
	return partitionIndex
}
