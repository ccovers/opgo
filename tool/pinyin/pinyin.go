package pinyin

import (
	"strings"

	// github.com/lovego/pinyin
	"github.com/mozillazg/go-pinyin"
)

var firstLetters = getArgs()

func getArgs() pinyin.Args {
	a := pinyin.NewArgs()
	a.Style = pinyin.FirstLetter
	return a
}

// 获取所有汉字的首字母组合， 比如"长城ABC"，返回"CCABC"
func Initials(str string) string {
	var runes = []rune(strings.TrimSpace(str))
	var result = make([]rune, 0, len(runes))
	for _, r := range runes {
		letters := pinyin.SinglePinyin(r, firstLetters)
		if len(letters) > 0 && len(letters[0]) > 0 {
			r = rune(letters[0][0])
		}

		if r >= 'a' && r <= 'z' {
			r -= ('a' - 'A')
		}

		result = append(result, r)
	}
	return string(result)
}
