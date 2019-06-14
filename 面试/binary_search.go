package main

import (
	"fmt"
)

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
