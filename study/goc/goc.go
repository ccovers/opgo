package main

//https://studygolang.com/articles/3190
//https://blog.csdn.net/zdy0_2004/article/details/79124269?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromBaidu-1.not_use_machine_learn_pai&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromBaidu-1.not_use_machine_learn_pai
//https://blog.csdn.net/u014633283/article/details/52225274?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromBaidu-1.not_use_machine_learn_pai&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromBaidu-1.not_use_machine_learn_pai

/*
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <unistd.h>
//#include "dpdk_demo.h"

struct stat sayHi(char *filename, int n) {
	struct stat buf;
	if(n <= 0 || filename == "") {
		return buf;
	}
	char *name = (char *)malloc(n+1);
	name[n] = '\0';
	memcpy(name, filename, n);

	stat(name, &buf);
	free(name);
	return buf;
}

struct PortCard {
	struct PortCard *next;
	char macbuf[32];
};
struct PortCard* dpdk_pcap_init(int mac_count, const char **mac_addr, int param_count, char **params)
{
	struct PortCard *head_card = NULL;
	head_card = (struct PortCard *)malloc(sizeof(struct PortCard));
	head_card->next = NULL;
	return head_card;
}
struct AAA {
	char* S;
	int N;
};
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	filename := "/etc/hosts"
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))

	buf := C.sayHi(cs, C.int(len(filename)))
	fmt.Printf("%s: %d\n", filename, buf.st_size)

	p := (*reflect.StringHeader)(unsafe.Pointer(&filename))
	buf = C.sayHi((*C.char)(unsafe.Pointer(p.Data)), C.int(len(filename)))
	fmt.Printf("%s: %d\n", filename, buf.st_size)

	s := C.CString("vvv")
	defer C.free(unsafe.Pointer(s))
	x := C.struct_AAA{
		S: s,
		N: 10,
	}
	fmt.Printf("%v, %v\n", C.GoString(x.S), x.N)

	var s0 string
	var s0Hdr = (*reflect.StringHeader)(unsafe.Pointer(&s0))
	s0Hdr.Data = uintptr(unsafe.Pointer(s))
	s0Hdr.Len = int(C.strlen(s))
	fmt.Println(s0)

	go_mac_addr := []string{"w0", "w1", "w2", "w3"}
	c_mac_addr := make([]*C.char, len(go_mac_addr))
	for i, _ := range go_mac_addr {
		c_mac_addr[i] = C.CString(go_mac_addr[i])
		defer C.free(unsafe.Pointer(c_mac_addr[i]))
	}
	go_params := []string{"-n 4", "--proc-type=auto", "-m 512"}
	c_params := make([]*C.char, len(go_params))
	for i, _ := range go_params {
		c_params[i] = C.CString(go_params[i])
		defer C.free(unsafe.Pointer(c_params[i]))
	}
	portCard := C.dpdk_pcap_init(C.int(len(go_mac_addr)), (**C.char)(unsafe.Pointer(&c_mac_addr[0])),
		C.int(len(go_params)), (**C.char)(unsafe.Pointer(&c_params[0])))
	fmt.Printf("%v\n", portCard)
}

func GoStrings(length int, argv **C.char) []string {
	if argv == nil {
		return nil
	}
	tmpslice := (*[1 << 30]*C.char)(unsafe.Pointer(argv))[:length:length]

	gostrings := make([]string, length)
	for i, s := range tmpslice {
		gostrings[i] = C.GoString(s)
	}
	return gostrings
}
