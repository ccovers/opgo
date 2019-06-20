package main

import (
	"bou.ke/monkey"
	"fmt"
	"work/opgo/example/xpp"
)

/*
猴子补丁
在测试用例中挺好用的：https://github.com/bouk/monkey。注意下载和引入是`bou.ke/monkey`，github地址只是放的代码，而不是引入用的地址。

参考：https://www.jianshu.com/p/2f675d5e334e?utm_campaign=maleskine&utm_content=note&utm_medium=seo_notes&utm_source=recommendation

但有时候可能不生效，比如有内敛优化的时候，此时可以加上-gcflags=-l禁用内敛优化来测试，但也未必会生效。
*/

func main() {
	p := monkey.Patch(xpp.GetName, func(id int) string {
		return "替换值"
	})
	defer p.Unpatch()

	test()
}

func test() {
	fmt.Println(xpp.GetName(1))
}
