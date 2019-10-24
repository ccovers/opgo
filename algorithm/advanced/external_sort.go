package main

import (
	"fmt"
	"math/rand"
	//"time"
)

/*
* 外部排序-有限内存排序
* 大量数据需要排序，但是内存有限制的情况
* 有一个很大的文件需要对内容进行排序，如何在有限的内存下进行排序，内存很小。
* 分析：
	1.文件很大我们需要分而治之，分为若干文件
	2.内存小，划分小文件的时候要注意，文件内容应该可以足够放入内存
	3.拆分小文件的时候，对改文件内容进行排序（ps:非本文章重点故省略）
	4.对有序的文件进行归并排序
*/
type ExternalArray []int

func (array ExternalArray) data(start, end, length int) ([]int, int) {
	if start+length > end {
		length = end - start
	}
	return array[start : start+length], length
}
func (array ExternalArray) len() int {
	return len(array)
}
func (array ExternalArray) write(data []int, offset int) int {
	length := len(data)
	cnt := 0
	for i := 0; i < length; i++ {
		if i+offset >= len(array) {
			break
		}
		array[i+offset] = data[i]
		cnt += 1
	}
	return cnt
}

var externalInArray ExternalArray
var externalOutArray ExternalArray
var max int = 1024
var min int = 10

func init() {
	externalInArray = make([]int, max*max)
	externalOutArray = make([]int, max*max)
	//rand.Seed(time.Now().Unix())

	for i := 0; i < max*max; i++ {
		externalInArray[i] = rand.Intn(10)
		externalOutArray[i] = -1
	}
}

func main() {
	fmt.Println(externalInArray)
	externalSort()
	fmt.Println(externalOutArray)
}

type Data struct {
	Data   []int
	Offset int
}

func externalSort() {
	cnt := 0
	length := externalInArray.len()
	for i := 0; i < length; i += max {
		nums, numlen := externalInArray.data(i, i+max, max)
		quick_sort(nums, 0, numlen-1)
		cnt += 1
	}
	fmt.Println(externalInArray)

	merge_sort(cnt, min, max)
}

// 归并
func merge_sort(bufCnt, bufLen, maxLen int) {
	datas := make([]*Data, bufCnt)

	offset := 0
	for {
		tmps := make([]int, bufLen)
		tmpLen := 0
		for k := 0; k < bufLen; k++ {
			index := -1
			for i := 0; i < bufCnt; i++ {
				if datas[i] == nil {
					datas[i] = &Data{
						Offset: i * maxLen,
					}
				}
				if len(datas[i].Data) == 0 {
					datas[i].Data, _ =
						externalInArray.data(datas[i].Offset, i*maxLen+maxLen, bufLen)
					datas[i].Offset += len(datas[i].Data)
					if len(datas[i].Data) == 0 {
						continue
					}
				}
				if index < 0 {
					index = i
				} else {
					if datas[i].Data[0] < datas[index].Data[0] {
						index = i
					}
				}
			}
			if index < 0 {
				break
			}

			tmps[k] = datas[index].Data[0]
			datas[index].Data = datas[index].Data[1:]
			tmpLen += 1
		}

		if tmpLen == 0 {
			break
		}

		offsetEx := externalOutArray.write(tmps[0:tmpLen], offset)
		if offsetEx == 0 {
			break
		}
		offset += offsetEx
	}
}

// 快排
func quick_sort(nums []int, left, right int) {
	if left < right {
		mid := partion(nums, left, right)
		quick_sort(nums, left, mid-1)
		quick_sort(nums, mid+1, right)
	}
}

func partion(nums []int, left, right int) int {
	for left < right {
		for ; left < right; right-- {
			if nums[left] > nums[right] {
				nums[left], nums[right] = nums[right], nums[left]
				break
			}
		}

		for ; left < right; left++ {
			if nums[left] > nums[right] {
				nums[left], nums[right] = nums[right], nums[left]
				break
			}
		}
	}
	return left
}
