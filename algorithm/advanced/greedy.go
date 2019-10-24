package main

import (
	"fmt"
	"sort"
)

/*
* 贪心算法
	对某些求解最优问题的更为简单的方法。贪心算法每次都考虑一个局部最优解，总是考虑当前状态下的最优的选择。所以贪心算法并不是对每个问题都有最优解的，但是某些问题，比如单源最短路径，最小生成树问题。（关键是贪心策略的选择，选择的贪心策略必须具备无后效性，即某个状态以前的过程不会影响以后的状态，只与当前状态有关）
* 算法思路：
	建立数学模型来描述问题；
	把求解的问题分成若干个子问题；
	对每一子问题求解，得到子问题的局部最优解；
	把子问题的解局部最优解合成原来解问题的一个解。
*/

/*

假设现在有一批宝物，价值和重量如表 2-3 所示，毛驴运载能力 m=30，那么怎么装入
最大价值的物品？
表 2-3  宝物清单
宝物 i  1  2  3  4  5  6  7  8  9  10
重量 w[i]  4  2  9  5  5  8  5  4  5  5
价值 v[i]  3  8  18  6  8  20  5  6  7  15
*/
type Info struct {
	Wight  float32 //重量
	Amount float32 // 价值
	Px     float32 // 性价比
}
type InfoArray []Info

func (p InfoArray) Less(i, j int) bool {
	return p[i].Px > p[j].Px
}
func (p InfoArray) Len() int {
	return len(p)
}
func (p InfoArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	infos := InfoArray{
		{4, 3, 0}, {2, 8, 0}, {9, 18, 0}, {5, 6, 0}, {5, 8, 0}, {8, 20, 0},
		{5, 5, 0}, {4, 6, 0}, {5, 7, 0}, {5, 15, 0}}
	for i, _ := range infos {
		infos[i].Px = infos[i].Amount / infos[i].Wight
	}
	sort.Sort(infos)
	fmt.Println(infos, "\n", getMaxAmount(30, infos))
}

func getMaxAmount(m float32, infos []Info) float32 {
	var amount float32 = 0
	for _, info := range infos {
		if m >= info.Wight {
			amount += info.Amount
			m -= info.Wight
		} else {
			amount += info.Px * m
			m = 0
			break
		}
	}
	return amount
}
