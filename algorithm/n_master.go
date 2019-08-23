package main

import (
	"fmt"
)

// n 皇后问题

func main() {
	solveNQueens(4)
}

func solveNQueens(n int) {
	for i := 0; i < n; i++ {
		solveNQueen(0+1, n, map[int]int{0: i})
	}
}

func solveNQueen(n, h int, logMap map[int]int) {
	if n >= h {
		if len(logMap) == h {
			for i := 0; i < h; i++ {
				k, _ := logMap[i]
				for j := 0; j < h; j++ {
					if k == j {
						fmt.Printf("* ")
					} else {
						fmt.Printf("- ")
					}
				}
				fmt.Printf("\n")
			}
			fmt.Printf("=========================\n")
		}
		return
	}

	for i := 0; i < h; i++ {
		flag := 0
		for row, line := range logMap {
			if line == i || row-n == line-i || row-n == i-line {
				flag = 1
				break
			}
		}
		if flag != 0 {
			continue
		}

		logMap[n] = i
		solveNQueen(n+1, h, logMap)
		delete(logMap, n)
	}
}
