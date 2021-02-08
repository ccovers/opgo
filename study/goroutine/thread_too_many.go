package main

import (
	"fmt"
	"net"
	"runtime"
	"runtime/pprof"
	"sync"
	//"runtime/debug"
)

var threadProfile = pprof.Lookup("threadcreate")

func main() {
	fmt.Printf("cpu num: %d\n", runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
	//debug.SetMaxThreads(10)

	fmt.Printf("threads in starting: %d\n", threadProfile.Count())

	var wg sync.WaitGroup
	wg.Add(300)
	for i := 0; i < 300; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 300; j++ {
				// syscall 系统调用阻塞，导致创建太多线程来运行协程
				net.LookupHost("www.google.com")
			}
			if threadProfile.Count() > 10 {
				runtime.LockOSThread()
				//runtime.UnlockOSThread()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("threads after lookuphost: %d\n", threadProfile.Count())
}
