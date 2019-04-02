package main

import "fmt"

/**
输入一棵二叉树前序遍历和中序遍历的结果，请重建该二叉树。

注意:

二叉树中每个节点的值都互不相同；
输入的前序遍历和中序遍历一定合法；
样例
给定：
前序遍历是：[3, 9, 20, 15, 7]
中序遍历是：[9, 3, 15, 20, 7]

返回：[3, 9, 20, null, null, 15, 7, null, null, null, null]
返回的二叉树如下所示：
    3
   / \
  9  20
    /  \
   15   7
*/
func main() {
	tree := BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(tree)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pos = make(map[int]int, 0)

func BuildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	for k, v := range inorder {
		pos[v] = k
	}
	return construct(preorder, inorder, 0, n-1, 0, n-1)
}

func construct(preorder []int, inorder []int, pl, pr, il, ir int) *TreeNode {
	if pl > pr {
		return nil
	}

	k := pos[preorder[pl]] - il

	result := &TreeNode{
		Val: preorder[pl],
	}

	result.Left = construct(preorder, inorder, pl+1, pl+k, il, il+k-1)
	result.Right = construct(preorder, inorder, pl+k+1, pr, il+k+1, ir)
	return result
}
