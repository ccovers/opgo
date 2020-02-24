package main

import (
	"fmt"
)

//侦探调查了罪案的四位证人。从证人的话侦探得出的结论是：如果男管家说的是真话，那么厨师说的也是真话；厨师和园丁说的不可能都是真话；园丁和杂役不可能都在说谎；如果杂役说真话，那么厨师在说谎。侦探能判定这四位证人分别是在说谎还是在说真话吗？解释你的推理

func main() {
	getBool := func(v int) bool {
		if v == 0 {
			return false
		}
		return true
	}

	for a := 0; a <= 1; a++ {
		for b := 0; b <= 1; b++ {
			for c := 0; c <= 1; c++ {
				for d := 0; d <= 1; d++ {
					err := logic(getBool(a), getBool(b), getBool(c), getBool(d))
					if err == nil {
						fmt.Printf("管家:%t, 厨师:%t, 园丁:%t, 杂役:%t\n", getBool(a), getBool(b), getBool(c), getBool(d))
					}
				}
			}
		}
	}
}

func logic(a /*管家*/, b /*厨师*/, c /*园丁*/, d /*杂役*/ bool) error {
	if a && !b {
		return fmt.Errorf("如果男管家[%t]说的是真话，那么厨师[%t]说的也是真话", a, b)
	}

	if b && c {
		return fmt.Errorf("厨师[%t]和园丁[%t]说的不可能都是真话", b, c)
	}

	if !c && !d {
		return fmt.Errorf("园丁[%t]和杂役[%t]不可能都在说谎", c, d)
	}

	if d && b {
		return fmt.Errorf("如果杂役[%t]说真话，那么厨师[%t]在说谎", d, b)
	}
	return nil
}
