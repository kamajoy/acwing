package main

import "fmt"

/**
在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。

请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

输入数组：

[
  [1,2,8,9]，
  [2,4,9,12]，
  [4,7,10,13]，
  [6,8,11,15]
]

如果输入查找数值为7，则返回true，

如果输入查找数值为5，则返回false。
 */
func main() {
	array := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}

	b := searchArray(array, 7)
	fmt.Println(b)

	b = searchArray(array, 5)
	fmt.Println(b)
}

func searchArray(array [][]int, target int) bool {
	if array == nil || len(array) == 0 {
		return false
	}

	i, j := 0, len(array[0])-1
	for i < len(array) && j >= 0 {
		if target == array[i][j] {
			return true
		}

		if array[i][j] > target {
			j--
		} else {
			i++
		}
	}

	return false
}
