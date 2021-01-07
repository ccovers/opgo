package main

import (
	"testing"
)

/*
## 测试覆盖率
1.输出
go test -run . -coverprofile=c.out

2.web展示
go tool cover -html=c.out


## 基准测试
go test -bench=.

## 剖析
1.输出
go test -cpuprofile=cpu.out
go test -blockprofile=block.out
go test -memprofile=mem.out
go test -run=NONE -bench=BenchmarkOpgoTest -cpuprofile=cpu.out

2.分析
go tool pprof -text -nodecount=10 . cpu.out
*/

func TestOpgoTest(t *testing.T) {
	slice := make([]int, 0)
	for i := 0; i < 10000000; i++ {
		slice = append(slice, i)
	}
}

func BenchmarkOpgoTest(b *testing.B) {
	slice := make([]int, 0)
	for i := 0; i < 10000000; i++ {
		slice = append(slice, i)
	}
}
