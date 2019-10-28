package main

import (
	"fmt"
)

/*
* 迪杰斯特拉算法(Dijkstra)
* 从一个顶点到其余各顶点的最短路径算法，解决的是有权图中最短路径问题。
	迪杰斯特拉算法主要特点是以起始点为中心向外层层扩展，直到扩展到终点为止。
* Dijkstra算法算是贪心思想实现的，首先把起点到所有点的距离存下来找个最短的，然后松弛一次再找出最短的，
	所谓的松弛操作就是，遍历一遍看通过刚刚找到的距离最短的点作为中转站会不会更近，
	如果更近了就更新距离，这样把所有的点找遍之后就存下了起点到其他所有点的最短距离。
*/

func main() {

}