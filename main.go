package main

/*
#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
//#include "dpdk_demo.h"

struct PortCard {
	struct PortCard *next;
	char macbuf[32];
	uint8_t port_id;
};

typedef void(*pcap_callback)(uint32_t cap_len, uint32_t pkt_len, uint8_t *packet, struct PortCard *port_card);

static struct PortCard* dpdk_pcap_init(int mac_count, const char **mac_addr, int param_count, char **params, pcap_callback pcap_cb)
{
	struct PortCard *head_card = NULL;
	head_card = (struct PortCard *)malloc(sizeof(struct PortCard));
	head_card->next = NULL;

	uint8_t *packet = (uint8_t*)malloc(sizeof(uint8_t) * 5);
	packet[0] = 100;
	pcap_cb(5, 5, packet, head_card);
	free(packet);
	return head_card;
}

extern void pcapcallback(uint32_t cap_len, uint32_t pkt_len, uint8_t *packet, struct PortCard *port_card);
*/
import "C"
import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

//export pcapcallback
func pcapcallback(cap_len C.uint32_t, pkt_len C.uint32_t, packet *C.uint8_t, port_card *C.struct_PortCard) {
	var gopacket []uint8
	var pPacket = (*reflect.SliceHeader)(unsafe.Pointer(&gopacket))
	pPacket.Data = uintptr(unsafe.Pointer(packet))
	pPacket.Len = int(pkt_len)
	pPacket.Cap = int(cap_len)

	fmt.Printf("gopacket: %v\n", gopacket)
}

func main() {
	t := time.Now()
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
		C.int(len(go_params)), (**C.char)(unsafe.Pointer(&c_params[0])), C.pcap_callback(C.pcapcallback))
	fmt.Printf("portCard: %v\n", portCard)
	fmt.Println(time.Since(t))
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

/*
#include <stdio.h>

__attribute__((constructor)) void before_main() {
   printf("Before main\n");
}
__attribute__((destructor)) void after_main() {
   printf("After main\n");
}

int Print(void* temp) {
    printf("xxxx\n");
}

int main(int argc, char **argv) {
    int ret;
    int lcore_id;

    ret = rte_eal_init(argc, argv);
    if (ret < 0)
		rte_panic("Cannot init EAL\n");

    RTE_LCORE_FOREACH_SLAVE(lcore_id) {
        rte_eal_remote_launch(Print, NULL, lcore_id);
    }

    rte_eal_mp_wait_lcore();
   return 0;
}
*/
