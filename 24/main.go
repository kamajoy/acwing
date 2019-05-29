package main

import "time"

/**
地上有一个 m 行和 n 列的方格，横纵坐标范围分别是 0∼m−1 和 0∼n−1。

一个机器人从坐标0,0的格子开始移动，每一次只能向左，右，上，下四个方向移动一格。

但是不能进入行坐标和列坐标的数位之和大于 k 的格子。

请问该机器人能够达到多少个格子？

样例1
输入：k=7, m=4, n=5

输出：20
样例2
输入：k=18, m=40, n=40

输出：1484

解释：当k为18时，机器人能够进入方格（35,37），因为3+5+3+7 = 18。
      但是，它不能进入方格（35,38），因为3+5+3+8 = 19。
注意:

0<=m<=50
0<=n<=50
0<=k<=100
*/
func getSum(x, y int) int {
	sum := 0
	for ; x != 0; {
		sum += x % 10
		x = x / 10
	}

	for ; y != 0; {
		sum += y % 10
		y = y / 10
	}

	return sum
}

type pair struct {
	x, y int
}

func movingCount(threshold, rows, cols int) int {
	if rows == 0 || cols == 0 {
		return 0
	}

	q := make(chan pair)
	st := make([][]bool, 0)

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	res := 0
	q <- pair{0, 0}
	for t := range q {
		if st[t.x][t.y] || getSum(t.x, t.y) > threshold {
			continue
		}
		res++
		st[t.x][t.y] = true

		for i := 0; i < 4; i++ {
			x := t.x + dx[i]
			y := t.y + dy[i]

			if x >= 0 && x < rows && y >= 0 && y < cols {
				q <- pair{x, y}
			}
		}
	}
	return res
}

func main() {
	go time.Sleep(1 * time.Second)
	println(movingCount(7, 4, 5))
}
