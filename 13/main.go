package main

import "fmt"

/**
给定一个长度为 n 的整数数组 nums，数组中所有的数字都在 0∼n−1 的范围内。

数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。

请找出数组中任意一个重复的数字。

注意：如果某些数字不在 0∼n−1 的范围内，或数组中不包含重复数字，则返回 -1；

样例
给定 nums = [2, 3, 5, 4, 3, 2, 6, 7]。

返回 2 或 3。

https://www.acwing.com/problem/content/14/
*/
func main() {
	array := duplicateInArray([]int{2, 3, 5, 4, 3, 2, 6, 7})
	fmt.Println(array)
}

func duplicateInArray(nums []int) int {
	length := len(nums)

	for _, v := range nums {
		if v < 0 || v > length-1 {
			return -1
		}
	}

	for i := 0; i < length; i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	return -1
}
