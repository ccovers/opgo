package main

import (
	"fmt"
)

// 归并排序
func main() {
	slice := []int{4, 3, 2, 5, 7, 1, 9, 8}

	merge_sort(slice, 0, len(slice)-1)
	fmt.Println(slice)
}

func merge_sort(slice []int, l int, r int) {
	if l >= r {
		return
	}

	mid := (l + r) / 2

	merge_sort(slice, l, mid)
	merge_sort(slice, mid+1, r)

	merge(slice, l, r, mid)
}

func merge(slice []int, l int, r int, mid int) {
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = slice[i]
	}

	ln := l
	rn := mid + 1
	for i := l; i <= r; i++ {
		if ln > mid {
			slice[i] = temp[rn-l]
			rn++
		} else if rn > r {
			slice[i] = temp[ln-l]
			ln++
		} else {
			if temp[ln-l] <= temp[rn-l] {
				slice[i] = temp[ln-l]
				ln++
			} else {
				slice[i] = temp[rn-l]
				rn++
			}
		}
	}
}
