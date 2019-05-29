package main

/*
输入一个32位整数，输出该数二进制表示中1的个数。

注意：

负数在计算机中用其绝对值的补码来表示。
样例1
输入：9
输出：2
解释：9的二进制表示是1001，一共有2个1。
样例2
输入：-2
输出：31
解释：-2在计算机里会被表示成11111111111111111111111111111110，
      一共有31个1。
*/

// NumberOf1 如果n为负数， 会陷入死循环
func NumberOf1(n int) int {
	var res = 0
	for ; n != 0; n >>= 1 {
		if n&1 == 1 {
			res++
		}
	}
	return res
}

func NumberOf2(n int) int {
	var res = 0
	var flag = 1
	for ; flag != 0; flag <<= 1 {
		if n&flag != 0 {
			res++
		}
	}
	return res
}

func NumberOf3(n int) int {
	var res = 0
	for ; n != 0; {
		res++
		n = (n - 1) & n
	}
	return res
}

func main() {
	println(NumberOf2(9))
}
