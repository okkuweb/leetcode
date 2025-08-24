package main

import "fmt"

// @leet start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	final := make(map[int]int)
	var loop func(*TreeNode, int)
	loop = func(node *TreeNode, depth int) {
		if _, ok := final[depth]; ok {
			if final[depth] < node.Val {
				final[depth] = node.Val
			}
		} else {
			final[depth] = node.Val
		}
		if node.Left != nil {
			loop(node.Left, depth+1)
		}
		if node.Right != nil {
			loop(node.Right, depth+1)
		}
	}
	if (root != nil) {
		loop(root, 0)
	}
	var keys = []int{}
	for k, _ := range final {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var finalSlice = []int{}
	for _, v := range keys {
		finalSlice = append(finalSlice, final[v])
	}
	return finalSlice
}

// @leet end
