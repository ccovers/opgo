package main

import (
	"fmt"
)

/*
力扣数据中心有 n 台服务器，分别按从 0 到 n-1 的方式进行了编号。
它们之间以「服务器到服务器」点对点的形式相互连接组成了一个内部集群，其中连接 connections 是无向的。
从形式上讲，connections[i] = [a, b] 表示服务器 a 和 b 之间形成连接。任何服务器都可以直接或者间接地通过网络到达任何其他服务器。
「关键连接」是在该集群中的重要连接，也就是说，假如我们将它移除，便会导致某些服务器无法访问其他服务器。
请你以任意顺序返回该集群内的所有 「关键连接」。

输入：n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
输出：[[1,3]]
解释：[[3,1]] 也是正确的。


提示：
1 <= n <= 10^5
n-1 <= connections.length <= 10^5
connections[i][0] != connections[i][1]
不存在重复的连接
*/

func main() {
	fmt.Println(criticalConnections(10, [][]int{[]int{1, 2}}))
}

func criticalConnections(n int, connections [][]int) [][]int {
	orgins := make([][]int, n)
	for _, conns := range connections {
		orgins[conns[0]] = append(orgins[conns[0]], conns[1])
		orgins[conns[1]] = append(orgins[conns[1]], conns[0])
	}

	nmap := make(map[int64]bool)
	for dst, arr := range orgins {
		for _, src := range arr {
			if !criticalConnection(orgins, dst, src, dst, map[int]bool{}) {
				var key int64
				if src < dst {
					key = int64(src) + int64(dst)<<32
				} else {
					key = int64(dst) + int64(src)<<32
				}
				_, ok := nmap[key]
				if !ok {
					nmap[key] = true
				}

			}
		}
	}

	dsts := make([][]int, 0)
	for key, _ := range nmap {
		dsts = append(dsts, []int{int(key & 0xFFFFFFFF), int(key >> 32)})
	}
	return dsts
}

func criticalConnection(origins [][]int, pre, src, dst int, path map[int]bool) bool {
	path[pre] = true
	arr := origins[src]

	for _, v := range arr {
		if v != pre && v == dst {
			return true
		}

		_, ok := path[v]
		if !ok {
			if criticalConnection(origins, src, v, dst, path) {
				return true
			}
		}
	}
	return false
}
