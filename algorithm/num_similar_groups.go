package main

import (
	"fmt"
)

/*
如果我们交换字符串 X 中的两个不同位置的字母，使得它和字符串 Y 相等，那么称 X 和 Y 两个字符串相似。

例如，"tars" 和 "rats" 是相似的 (交换 0 与 2 的位置)； "rats" 和 "arts" 也是相似的，但是 "star" 不与 "tars"，"rats"，或 "arts" 相似。

总之，它们通过相似性形成了两个关联组：{"tars", "rats", "arts"} 和 {"star"}。注意，"tars" 和 "arts" 是在同一组中，即使它们并不相似。形式上，对每个组而言，要确定一个单词在组中，只需要这个词和该组中至少一个单词相似。

我们给出了一个不包含重复的字符串列表 A。列表中的每个字符串都是 A 中其它所有字符串的一个字母异位词。请问 A 中有多少个相似字符串组？

示例：

输入：["tars","rats","arts","star"]
输出：2
提示：

A.length <= 2000
A[i].length <= 1000
A.length * A[i].length <= 20000
A 中的所有单词都只包含小写字母。
A 中的所有单词都具有相同的长度，且是彼此的字母异位词。
此问题的判断限制时间已经延长。
备注：

      字母异位词[anagram]，一种把某个字符串的字母的位置（顺序）加以改换所形成的新词。
*/

func main() {
	fmt.Println(numSimilarGroups([]string{"blw", "bwl", "wlb"}))

	fmt.Println(numSimilarGroups([]string{"tars", "rats", "arts", "star"}))
}

func numSimilarGroups(A []string) int {
	okMap := make(map[string]bool)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if i == j {
				continue
			}
			_, ok := okMap[fmt.Sprintf("%s-%s", A[i], A[j])]
			if ok {
				continue
			}
			_, ok = okMap[fmt.Sprintf("%s-%s", A[j], A[i])]
			if ok {
				continue
			}

			iok := judge(A[i], A[j])
			if iok {
				okMap[fmt.Sprintf("%s-%s", A[i], A[j])] = true
			}
		}
	}
	return len(okMap)
}

func judge(as, bs string) bool {
	for i := 0; i < len(as); i++ {
		for j := 0; j < len(bs); j++ {
			if i == j {
				continue
			}
			tmp := make([]byte, 0)
			tmp = append(tmp, []byte(as)...)
			tmp[i] = as[j]
			tmp[j] = as[i]
			if string(tmp) == bs {
				fmt.Println(as, bs, string(tmp), i, j)
				return true
			}
		}
	}
	return false
}
