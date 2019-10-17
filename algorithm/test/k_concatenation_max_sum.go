package main

import (
	"fmt"
	"math"
)

/*
给你一个整数数组 arr 和一个整数 k。
首先，我们要对该数组进行修改，即把原数组 arr 重复 k 次。
举个例子，如果 arr = [1, 2] 且 k = 3，那么修改后的数组就是 [1, 2, 1, 2, 1, 2]。
然后，请你返回修改后的数组中的最大的子数组之和。
注意，子数组长度可以是 0，在这种情况下它的总和也是 0。
由于 结果可能会很大，所以需要 模（mod） 10^9 + 7 后再返回。
*/

func main() {

}

func kConcatenationMaxSum(arr []int, k int) int {
	single := 0
	for _, v := range arr {
		single += v
	}
	tmp := 0
	start := 0
	end := 0
	for i := 0; i < len(arr); i++ {
		tmp += arr[i]
		if tmp > start {
			start = tmp
		}
	}
	tmp = 0
	for i := len(arr) - 1; i >= 0; i-- {
		tmp += arr[i]
		if tmp > end {
			end = tmp
		}
	}

	tempArr := make([]int, 0)
	max := 0
	if k >= 1 {
		tempArr = append(tempArr, arr...)
		max = getMax(tempArr)
		if k >= 2 {
			tempArr = append(tempArr, arr...)
			two := getMax(tempArr)
			if two > max {
				max = two
			}
			if k >= 3 {
				if (k-2)*single+start+end > max {
					max = (k-2)*single + start + end
				}
			}
		}
	}
	return max % (getNum(10, 9) + 7)
}

func getNum(n int, k int) int {
	num := 1
	for i := 0; i < k; i++ {
		num *= n
	}
	return num
}

func getMax(arr []int) int {
	oldMax := 0
	max := 0
	for _, v := range arr {
		if max+v > 0 {
			max += v
			if max > oldMax {
				oldMax = max
			}
		} else {
			max = 0
		}
	}
	return oldMax
}
