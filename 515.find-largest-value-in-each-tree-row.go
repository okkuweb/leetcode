package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
	var final = []int{}
	if root == nil {
		return final
	}
	var loop func(*TreeNode, int)
	loop = func(node *TreeNode, depth int) {
		if len(final) <= depth {
			final = append(final, node.Val)
		} else if final[depth] < node.Val {
			final[depth] = node.Val
		}
		if node.Left != nil {
			loop(node.Left, depth+1)
		}
		if node.Right != nil {
			loop(node.Right, depth+1)
		}
	}
	loop(root, 0)
	return final
}

// @leet end
