package main

/*
#cgo CFLAGS: -g -O0 -march=native -I/usr/local/dpdk/x86_64-native-linuxapp-gcc/include -I.
#cgo LDFLAGS: -lrt -pthread -ldl -lnuma -L/usr/local/dpdk/x86_64-native-linuxapp-gcc/lib -Wl,--whole-archive -lrte_pmd_virtio -lrte_pmd_vmxnet3_uio -lrte_pmd_e1000 -lrte_pmd_ixgbe -lrte_pmd_af_packet -lrte_pmd_bond -lrte_pmd_fm10k -lrte_pmd_enic -lrte_pmd_i40e -lrte_pmd_null -lrte_net -lrte_bus_pci -lrte_bus_vdev -lrte_pci -lrte_mbuf -lrte_eal -lrte_mempool -lrte_ring -lrte_ethdev -lrte_kvargs -lrte_hash -lrte_cmdline -lrte_mempool_ring -Wl,--no-whole-archive

#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include "dpdk_demo.h"

extern void pcapcallback(uint32_t cap_len, uint32_t pkt_len, uint8_t *packet, struct PortCard *port_card);
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

//export pcapcallback
func pcapcallback(cap_len C.uint32_t, pkt_len C.uint32_t, packet *C.uint8_t, port_card *C.struct_PortCard) {
	var gopacket []uint8
	var pPacket = (*reflect.SliceHeader)(unsafe.Pointer(&gopacket))
	pPacket.Data = uintptr(unsafe.Pointer(packet))
	pPacket.Len = int(pkt_len)
	pPacket.Cap = int(cap_len)

	fmt.Printf("cap_len: %v, pkt_len: %v, \ndata: %v\n", cap_len, pkt_len, gopacket)
	fmt.Printf("data: %v\n", pPacket.Data)
}

// ldconfig
// CGO_LDFLAGS_ALLOW='-Wl,--.*' go build
func main() {
	go_mac_addr := []string{"fa:16:3e:51:0a:93"}
	c_mac_addr := make([]*C.char, len(go_mac_addr))
	for i, _ := range go_mac_addr {
		c_mac_addr[i] = C.CString(go_mac_addr[i])
		defer C.free(unsafe.Pointer(c_mac_addr[i]))
	}
	go_params := []string{
		"-c", "0x03",
		"-n", "2", // number of memory channels per processor socket
		//"--proc-type=auto", // The type of process instance.
		"-m", "512", // Hugepages
	}
	c_params := make([]*C.char, len(go_params))
	for i, _ := range go_params {
		c_params[i] = C.CString(go_params[i])
		defer C.free(unsafe.Pointer(c_params[i]))
	}
	portCard := C.dpdk_pcap_init(C.int(len(go_mac_addr)), (**C.char)(unsafe.Pointer(&c_mac_addr[0])),
		C.int(len(go_params)), (**C.char)(unsafe.Pointer(&c_params[0])), C.pcap_callback(C.pcapcallback))
	if nil == portCard {
		fmt.Printf("init error\n")
		return
	}
	fmt.Printf("portCard: %v\n", portCard)

	C.dpdk_pcap_loop()
}
