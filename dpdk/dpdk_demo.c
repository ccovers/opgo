#include "dpdk_demo.h"
#include <rte_eal.h>
#include <rte_ethdev.h>
#include <rte_cycles.h>
#include <rte_lcore.h>
#include <rte_mbuf.h>
#include <rte_mempool.h>
#include <rte_mbuf.h>
#include <rte_ether.h>
#include <rte_ip.h>
#include <rte_tcp.h>
#include <rte_ethdev.h>

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <string.h>
#include <net/if.h>
#include <sys/ioctl.h>


//指定的网卡列表
struct PortCard *g_port_card = NULL;
//数据包抓包回调
pcap_callback g_pcap_cb = NULL;
//缺省配置
static const struct rte_eth_conf dev_conf_default = {
	.rxmode = { .max_rx_pkt_len = ETHER_MAX_LEN, },
};



uint16_t rte_rx_callback(uint8_t port __rte_unused, uint16_t queue __rte_unused,
	struct rte_mbuf *pkts[], uint16_t nb_pkts, uint16_t max_pkts __rte_unused, void *user_param)
{
	if (g_pcap_cb && nb_pkts > 0) {
		uint16_t index;
		for (index = 0; index < nb_pkts; index++) {
			struct rte_mbuf *mbuf = pkts[index];
			g_pcap_cb(mbuf->data_len, mbuf->pkt_len, (unsigned char*)mbuf->buf_addr + mbuf->data_off, (struct PortCard*)user_param);
		}
	}
	return nb_pkts;
}

int port_init(struct PortCard *port_card, struct rte_mempool *mbuf_pool)
{
	struct rte_eth_conf dev_conf = dev_conf_default;
	const uint16_t nb_rx_q = 1;
	const uint16_t nb_tx_q = 0;
	uint8_t port_id = port_card->port_id;
	int retval;
	
	if (port_id >= rte_eth_dev_count()) {
		return -1;
	}
	retval = rte_eth_dev_configure(port_id, nb_rx_q, nb_tx_q, &dev_conf);
	if (retval != 0) {
		return -1;
	}
	uint16_t queue_id;
	for (queue_id = 0; queue_id < nb_rx_q; queue_id++) {
		retval = rte_eth_rx_queue_setup(port_id, queue_id, RX_RING_SIZE, 
			rte_eth_dev_socket_id(port_id), NULL, mbuf_pool);
		if (retval < 0) {
			return -1;
		}
	}
	/*for (queue_id = 0; queue_id < nb_tx_q; queue_id++) {
		retval = rte_eth_tx_queue_setup(port_id, queue_id, TX_RING_SIZE, 
			rte_eth_dev_socket_id(port_id), NULL);
		if (retval < 0) {
			return -1;
		}
	}*/
	retval = rte_eth_dev_start(port_id);
	if (retval < 0) {
		return -1;
	}

	/*struct rte_eth_dev_info dev_info;
	rte_eth_dev_info_get(port_id, &dev_info);
	struct ether_addr mac_addr;
	rte_eth_macaddr_get(port_id, &mac_addr);
	printf("Dev %u MAC: %02"PRIx8" %02"PRIx8" %02"PRIx8
		" %02"PRIx8" %02"PRIx8" %02"PRIx8"\n", port_id, 
		mac_addr.addr_bytes[0], mac_addr.addr_bytes[1], 
		mac_addr.addr_bytes[2], mac_addr.addr_bytes[3],
		mac_addr.addr_bytes[4], mac_addr.addr_bytes[5]);*/
	rte_eth_promiscuous_enable(port_id);
	//rx下行流量        tx上行流量
	rte_eth_add_rx_callback(port_id, 0, rte_rx_callback, (void*)port_card);
	//rte_eth_add_tx_callback(dev_id, 0, rte_tx_callback_fn fn, NULL);
	return 0;
}

/*匹配指定的网卡*/
struct PortCard* get_match_networkcard(int mac_count, const char **mac_addr)
{
	//uint32 nb_cores = rte_lcore_count(); //内核数
	uint8_t ports = rte_eth_dev_count(); //igb网卡数

	struct PortCard *head_card = NULL;
	struct PortCard *temp_card = NULL;
	char macbuf[32] = {0};
	uint8_t port_id;
	int index;
	
	for (port_id = 0; port_id < ports; port_id++) {
		struct ether_addr ether_addr;
		rte_eth_macaddr_get(port_id, &ether_addr);
		snprintf(macbuf, 32, "%02x:%02x:%02x:%02x:%02x:%02x", 
			ether_addr.addr_bytes[0], ether_addr.addr_bytes[1], 
			ether_addr.addr_bytes[2], ether_addr.addr_bytes[3], 
			ether_addr.addr_bytes[4], ether_addr.addr_bytes[5]);
		
		for (index = 0; index < mac_count; index++) {
			//匹配网卡
			if ((strlen(mac_addr[index]) == strlen(macbuf)) && (strcmp(mac_addr[index], macbuf) == 0)) {
				if (head_card != NULL) {
					temp_card->next = (struct PortCard *)malloc(sizeof(struct PortCard));
					temp_card = temp_card->next;
				} else {
					head_card = (struct PortCard *)malloc(sizeof(struct PortCard));
					temp_card = head_card;
				}
				temp_card->next = NULL;
				temp_card->port_id = port_id;
				strcpy(temp_card->macbuf, macbuf);
				break;
			}
		}
	}
	return head_card;
}

struct PortCard* dpdk_pcap_init(int mac_count, const char **mac_addr, int param_count, char **params, pcap_callback pcap_cb)
{
	//抓包回调
	g_pcap_cb = pcap_cb;
	
	/*初始化EAL*/
	int ret = rte_eal_init(param_count, params);
	if (ret < 0) {
		printf("EAL initialization error[%d]\n", ret);
		return NULL;
	}
	/*初始化大页内存*/
	struct rte_mempool *mbuf_pool = NULL;
	mbuf_pool = rte_pktmbuf_pool_create("MBUF_POOL", NUM_MBUFS * mac_count, 
		MBUF_CACHE_SIZE, 0, RTE_MBUF_DEFAULT_BUF_SIZE, rte_socket_id());
	if (NULL == mbuf_pool) {
		printf("Cannot create mbuf pool\n");
		return NULL;
	}
	/*初始化网卡*/
	struct PortCard *port_card = NULL;
	g_port_card = get_match_networkcard(mac_count, mac_addr);
	for (port_card = g_port_card; port_card != NULL; port_card = port_card->next) {
		if (port_init(port_card, mbuf_pool) != 0) {
			printf("Cannot init network card: %s\n", port_card->macbuf);
			return NULL;
		}
	}
	return g_port_card;
}

void dpdk_pcap_loop(void)
{
	uint16_t index;
	struct PortCard *port_card = NULL;
	struct rte_mbuf *bufs[BURST_SIZE];
	for (;;) {
		for (port_card = g_port_card; port_card != NULL; port_card = port_card->next) {
			const uint16_t nb_rx = rte_eth_rx_burst(g_port_card->port_id, 0, bufs, BURST_SIZE);
			if (unlikely(0 == nb_rx)) {
				continue;
			}
			for (index = 0; index < nb_rx; index++) {
				rte_pktmbuf_free(bufs[index]);
			}
		}
	}
}

