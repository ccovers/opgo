package main

/*
给你一个字符串 text，你需要使用 text 中的字母来拼凑尽可能多的单词 "balloon"（气球）。
字符串 text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词 "balloon"。
*/

func maxNumberOfBalloons(text string) int {
	b := 0
	a := 0
	l := 0
	o := 0
	n := 0
	for i, _ := range text {
		switch text[i] {
		case 'b':
			b += 1
		case 'a':
			a += 1
		case 'l':
			l += 1
		case 'o':
			o += 1
		case 'n':
			n += 1
		}
	}

	min := b
	if min > a {
		min = a
	}
	if min > l/2 {
		min = l / 2
	}
	if min > o/2 {
		min = o / 2
	}
	if min > n {
		min = n
	}
	return min
}
