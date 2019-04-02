package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildTree(t *testing.T) {
	expect := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	actual := BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	assert.EqualValues(t, expect, actual)
}
