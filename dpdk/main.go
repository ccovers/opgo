package main

/*
DPDK_LIB_DIR=$(RTE_SDK)/${RTE_TARGET}/lib
#网卡驱动
LDFLAGS += -lrt -pthread -ldl -Wl,--whole-archive \
		  $(DPDK_LIB_DIR)/librte_pmd_vmxnet3_uio.a \
		  $(DPDK_LIB_DIR)/librte_pmd_e1000.a \
		  $(DPDK_LIB_DIR)/librte_pmd_ixgbe.a \
		  $(DPDK_LIB_DIR)/librte_pmd_af_packet.a \
		  $(DPDK_LIB_DIR)/librte_pmd_bond.a \
		  $(DPDK_LIB_DIR)/librte_pmd_fm10k.a \
		  $(DPDK_LIB_DIR)/librte_pmd_enic.a \
		  $(DPDK_LIB_DIR)/librte_pmd_i40e.a \
		  $(DPDK_LIB_DIR)/librte_pmd_null.a \
		  -Wl,--no-whole-archive
#dpdk库
LDFLAGS += -Wl,--whole-archive \
		  $(DPDK_LIB_DIR)/librte_mbuf.a \
		  $(DPDK_LIB_DIR)/librte_eal.a \
		  $(DPDK_LIB_DIR)/librte_mempool.a \
		  $(DPDK_LIB_DIR)/librte_ring.a \
		  $(DPDK_LIB_DIR)/librte_ethdev.a \
		  $(DPDK_LIB_DIR)/librte_kvargs.a \
		  $(DPDK_LIB_DIR)/librte_hash.a \
		  $(DPDK_LIB_DIR)/librte_cmdline.a \
		  -Wl,--no-whole-archive
#cgo LDFLAGS += -L/root/work/pcap/dpdk_pcap/lib -ldpdk_pcap
#cgo CXFLAGS = -std=c++11 -Wall -g -O0 -I/root/work/pcap/dpdk_pcap/include

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

	fmt.Printf("gopacket: %v\n", gopacket)
}

func main() {
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

	C.dpdk_pcap_loop()
}
