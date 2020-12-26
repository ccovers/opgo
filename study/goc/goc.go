package main

//https://studygolang.com/articles/3190
//https://blog.csdn.net/zdy0_2004/article/details/79124269?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromBaidu-1.not_use_machine_learn_pai&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromBaidu-1.not_use_machine_learn_pai
//https://blog.csdn.net/u014633283/article/details/52225274?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromBaidu-1.not_use_machine_learn_pai&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromBaidu-1.not_use_machine_learn_pai

/*
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <unistd.h>
struct stat sayHi(char *filename)
{
	struct stat buf;
	if(filename=="")
	{
		return buf;
	}
	stat(filename, &buf);
	return buf;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	filename := "/etc/hosts"
	cs := C.CString(filename)
	buf := C.sayHi(cs)
	C.free(unsafe.Pointer(cs))
	fmt.Printf("%s: %d, [%+v]\n", filename, buf.st_size, buf)
}
