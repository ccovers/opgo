package main

import (
	"fmt"
	"math/rand"
	"time"
)

var externalInArray []int
var externalOutArray []int
var max int = 10

func main() {
	externalInArray = make([]int, max*max)
	externalOutArray = make([]int, max*max)
	rand.Seed(time.Now().Unix())

	for i := 0; i < max*max; i++ {
		externalInArray[i] = rand.Intn(100)
	}
	fmt.Println(externalInArray)
	externalSort(max)
	fmt.Println(externalInArray)
}

func externalSort(n int) {
	for i := 0; i < len(externalInArray); i += n {
		end := i + n
		if end > len(externalInArray) {
			end = len(externalInArray)
		}
		fmt.Printf("%+v [%d, %d]\n", externalInArray[i:end], 0, end-i-1)
		quick_sort(externalInArray[i:end], 0, end-i-1)
		fmt.Println("+++++")
	}
}

func quick_sort(nums []int, left, right int) {
	fmt.Println(left, right)
	if left < right {

		mid := partion(nums, left, right)
		quick_sort(nums, left, mid-1)
		quick_sort(nums, mid+1, right)
	}
}

func partion(nums []int, left, right int) int {
	fmt.Println(left, right)
	for left < right {
		for ; left < right; right-- {
			if nums[left] >= nums[right] {
				nums[left], nums[right] = nums[right], nums[left]
				break
			}
		}

		for ; left < right; left++ {
			if nums[left] >= nums[right] {
				nums[left], nums[right] = nums[right], nums[left]
				break
			}
		}
	}
	fmt.Println("==========")
	return left
}
