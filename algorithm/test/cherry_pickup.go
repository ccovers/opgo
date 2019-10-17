package main

/*
一个N x N的网格(grid) 代表了一块樱桃地，每个格子由以下三种数字的一种来表示：

0 表示这个格子是空的，所以你可以穿过它。
1 表示这个格子里装着一个樱桃，你可以摘到樱桃然后穿过它。
-1 表示这个格子里有荆棘，挡着你的路。
你的任务是在遵守下列规则的情况下，尽可能的摘到最多樱桃：

从位置 (0, 0) 出发，最后到达 (N-1, N-1) ，只能向下或向右走，并且只能穿越有效的格子（即只可以穿过值为0或者1的格子）；
当到达 (N-1, N-1) 后，你要继续走，直到返回到 (0, 0) ，只能向上或向左走，并且只能穿越有效的格子；
当你经过一个格子且这个格子包含一个樱桃时，你将摘到樱桃并且这个格子会变成空的（值变为0）；
如果在 (0, 0) 和 (N-1, N-1) 之间不存在一条可经过的路径，则没有任何一个樱桃能被摘到。
示例 1:

输入: grid =
[[0, 1, -1],
 [1, 0, -1],
 [1, 1,  1]]
输出: 5
解释：
玩家从（0,0）点出发，经过了向下走，向下走，向右走，向右走，到达了点(2, 2)。
在这趟单程中，总共摘到了4颗樱桃，矩阵变成了[[0,1,-1],[0,0,-1],[0,0,0]]。
接着，这名玩家向左走，向上走，向上走，向左走，返回了起始点，又摘到了1颗樱桃。
在旅程中，总共摘到了5颗樱桃，这是可以摘到的最大值了。
说明:

grid 是一个 N * N 的二维数组，N的取值范围是1 <= N <= 50。
每一个 grid[i][j] 都是集合 {-1, 0, 1}其中的一个数。
可以保证起点 grid[0][0] 和终点 grid[N-1][N-1] 的值都不会是 -1。
*/

import (
	"fmt"
)

const (
	N = 3
)

type Point struct {
	X int
	Y int
}

func main() {
	grid := [N][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}}
	cnt := cherryPickup(grid[0:])
	fmt.Printf("max: %d\n", cnt)
}

func cherryPickup(grid [][]int) int {
	p := Point{X: 0, Y: 0}
	return getToPath(grid, &p, 0, []Point{p})
}

func getToPath(grid [][]int, o *Point, cnt int, points []Point) int {
	cgrid := [N][]int{{}, {}, {}}
	for i, line := range grid {
		for _, row := range line {
			cgrid[i] = append(cgrid[i], row)
		}
	}

	v := cgrid[o.X][o.Y]
	if v == -1 {
		return 0
	} else if v == 1 {
		cnt += 1
		cgrid[o.X][o.Y] = 0
	}

	if o.X == N-1 && o.Y == N-1 {
		cnt = getHomePath(cgrid[0:], &Point{X: o.X, Y: o.Y}, cnt, points)
	} else {
		var cntR int
		var cntD int
		if o.X < N-1 {
			p := Point{X: o.X + 1, Y: o.Y}
			cntR = getToPath(cgrid[0:], &p, cnt, append(points, p))
		}
		if o.Y < N-1 {
			p := Point{X: o.X, Y: o.Y + 1}
			cntD = getToPath(cgrid[0:], &p, cnt, append(points, p))
		}

		if cntR > cntD {
			cnt = cntR
		} else {
			cnt = cntD
		}
	}
	return cnt
}

func getHomePath(grid [][]int, o *Point, cnt int, points []Point) int {
	cgrid := [N][]int{{}, {}, {}}
	for i, line := range grid {
		for _, row := range line {
			cgrid[i] = append(cgrid[i], row)
		}
	}

	v := cgrid[o.X][o.Y]
	if v == -1 {
		return 0
	} else if v == 1 {
		cnt += 1
		cgrid[o.X][o.Y] = 0
	}

	if o.X == 0 && o.Y == 0 {
		// fmt.Println(points)
		// fmt.Println(cgrid)
		// fmt.Println(cnt)
	} else {
		var cntR int
		var cntD int
		if o.X > 0 {
			p := Point{X: o.X - 1, Y: o.Y}
			cntR = getHomePath(cgrid[0:], &p, cnt, append(points, p))
		}
		if o.Y > 0 {
			p := Point{X: o.X, Y: o.Y - 1}
			cntD = getHomePath(cgrid[0:], &p, cnt, append(points, p))
		}

		if cntR > cntD {
			cnt = cntR
		} else {
			cnt = cntD
		}
	}
	return cnt
}
