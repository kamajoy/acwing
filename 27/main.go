package main

/**
实现函数double Power(double base, int exponent)，求base的 exponent次方。

不得使用库函数，同时不需要考虑大数问题。

注意：

不会出现底数和指数同为0的情况
样例1
输入：10 ，2

输出：100
样例2
输入：10 ，-2

输出：0.01
*/

func Abs(value int) int {
	if value >= 0 {
		return value
	}
	return -value
}

func Power(base float64, exponent int) float64 {
	result := 1.0

	if base == 0 {
		return 0
	}

	for i := 0; i < Abs(exponent); i++ {
		result = result * float64(base)
	}

	if exponent < 0 {
		return 1 / result
	} else {
		return result
	}
}

func main() {
	println(Power(10, -2))
}
