package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randNumbers(min, max, n int) []int {
	numbers := []int{}
	rand.Seed(time.Now().Unix())

	for i := 0; i < n; i++ {
		numbers = append(numbers, min+rand.Intn(max-min))
	}
	return numbers
}

const (
	NumN = 100
)

func main() {
	numbers := randNumbers(0, NumN, NumN)

	numArray := make([]int, NumN)

	for _, v := range numbers {
		numArray[v] = numArray[v] + 1
	}

	for i, v := range numArray {
		if v > 1 {
			fmt.Printf("%d: %d\n", i, v)
		}
	}

}
