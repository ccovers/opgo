package main

import (
	"fmt"
)

/*
给出一个字符串 s（仅含有小写英文字母和括号）。
请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。
注意，您的结果中 不应 包含任何括号。

示例 1：
输入：s = "(abcd)"
输出："dcba"

示例 2：
输入：s = "(u(love)i)"
输出："iloveu"

示例 3：
输入：s = "(ed(et(oc))el)"
输出："leetcode"

示例 4：
输入：s = "a(bcdefghijkl(mno)p)q"
输出："apmnolkjihgfedcbq"
*/
func main() {
	fmt.Println(reverseParentheses("(ed(et(oc))el)"))
}

func reverseParentheses(s string) string {
	_, res := reverse(0, 0, []byte(s))
	return string(res)
}

func reverse(i, cnt int, s []byte) (int, []byte) {
	res := make([]byte, 0)

	for ; i < len(s); i++ {
		if s[i] == '(' {
			last, xs := reverse(i+1, cnt+1, s)
			i = last
			res = append(res, xs...)
		} else if s[i] == ')' {
			return i, unArray(res)
		} else {
			res = append(res, s[i])
		}
	}
	return len(s), res
}

func unArray(bytes []byte) []byte {
	length := len(bytes)

	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[length-i-1] = bytes[length-i-1], bytes[i]
	}
	return bytes
}
