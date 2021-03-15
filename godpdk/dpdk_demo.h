#ifndef _DPDK_SLAVE_H_
#define _DPDK_SLAVE_H_

#include <inttypes.h>


#define NUM_MBUFS 8191
#define MBUF_CACHE_SIZE 250
#define RX_RING_SIZE 128
#define TX_RING_SIZE 512
#define BURST_SIZE 32

struct PortCard {
	struct PortCard *next;
	char macbuf[32];
	uint8_t port_id;
};

typedef void (*pcap_callback)(uint32_t cap_len, uint32_t pkt_len, uint8_t *packet, struct PortCard *port_card);


/*初始化dpdk*/
struct PortCard* dpdk_pcap_init(int mac_count, const char **mac_addr, int param_count, char **params, pcap_callback pcap_cb);

/*dpdk轮询抓包*/
void dpdk_pcap_loop(void);

#endif
