package local

import (
	"github.com/huandu/goroutine"
)

func Goid() uint64 {
	return uint64(goroutine.GoroutineId())
}
