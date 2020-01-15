package main

import (
	"fmt"
)

/*
给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
示例 1:

输入: word1 = "horse", word2 = "ros"
输出: 3
解释:
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2:

输入: word1 = "intention", word2 = "execution"
输出: 5
解释:
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
*/
type Node struct {
	Val  int
	Next *Node
}

func main() {
	fmt.Println(minDistance("asd", "aasd"))

}

func minDistance(word1 string, word2 string) int {
	if len(word2) == 0 {
		return len(word1)
	}
	l1 := len(word1)
	l2 := len(word2)

	opCnt := 0
	index1 := 0
	index2 := 0

	if index1+1 == l1 {
		return opCnt + l2 - (index2 + 1)
	} else if index2+1 == l2 {
		return opCnt + l1 - (index1 + 1)
	} else {

	}
}
