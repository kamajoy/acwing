package main

/*
在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留。

样例1
输入：1->2->3->3->4->4->5

输出：1->2->5
样例2
输入：-1->1->1->1->2->3

输出：2->3
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplication(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	newHead := &ListNode{}
	tmp := newHead

	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			head = head.Next
		} else {
			tmp.Next = head
			tmp = tmp.Next
			head = head.Next
		}
	}

	tmp.Next = head
	return newHead.Next
}

func deleteDuplication2(head *ListNode) *ListNode {
	dummy := ListNode{Val: -1, Next: head}
	p := &dummy

	for p.Next != nil {
		q := p.Next

		for q != nil && p.Next.Val == q.Val {
			q = q.Next
		}

		if p.Next != nil && p.Next.Next == q {
			p = p.Next
		} else {
			p.Next = q
		}
	}

	return dummy.Next
}

func main() {
	head := ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}}}
	println(deleteDuplication2(&head).Val)
}
