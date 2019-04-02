package main

import "fmt"

/**
输入一个链表的头结点，按照 从尾到头 的顺序返回节点的值。

返回的结果用数组存储。

样例
输入：[2, 3, 5]
返回：[5, 3, 2]
 */
func main() {
	ints := printListReversingly(&ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}})
	for _, v := range ints {
		fmt.Println(v)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListReversingly(head *ListNode) []int {
	recursion(head)
	return result
}

var result []int


func recursion(head *ListNode) {
	if head != nil {
		if head.Next != nil {
			recursion(head.Next)
		}
		result = append(result, head.Val)
	}
}
