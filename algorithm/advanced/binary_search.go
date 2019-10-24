package main

import (
	"fmt"
)

/*
* 二分查找算法
* 二分查找操作的数据集是一个有序的数据集
* 先找出有序集合中间的那个元素。如果此元素比要查找的元素大，就接着在较小的一个半区进行查找；反之，如果此元素比要找的元素小，就在较大的一个半区进行查找。在每个更小的数据集中重复这个查找过程，直到找到要查找的元素或者数据集不能再分割
* 元素必须存储在连续的空间中。因此，当待搜索的集合是相对静态的数据集时，此时使用二分查找是最好的选择
 */

func main() {
	array := []int{1, 10, 50, 80, 90, 100, 101, 200, 1000}
	fmt.Println(binarySearch(array, 101, 0, len(array)))
}

func binarySearch(array []int, num int, offset, length int) int {
	if len(array) == 0 || len(array) < offset+length || length <= 0 {
		return -1
	}

	if array[offset] == num {
		return offset
	}
	if array[offset+length-1] == num {
		return offset + length - 1
	}

	mid := length / 2
	if array[offset+mid] == num {
		return offset + mid
	}
	if mid <= 1 {
		return -1
	}

	if num < array[offset] {
		return binarySearch(array, num, offset+1, mid-1)
	} else {
		return binarySearch(array, num, offset+mid+1, length-mid-2)
	}
}
