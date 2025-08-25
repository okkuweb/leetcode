package main

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insertionSortList(head *ListNode) *ListNode {
	finalList := &ListNode{}
	var finalArr []int

	var loop func(*ListNode)
	loop = func(node *ListNode) {
		finalArr = append(finalArr, node.Val)
		if node.Next != nil {
			loop(node.Next)
		}
	}
	loop(head)
	sort.Ints(finalArr)

	root := finalList
	for i, v := range finalArr {
		root.Val = v
		if i < len(finalArr)-1 {
			newNode := &ListNode{}
			root.Next = newNode
			root = newNode
		}
	}

	return finalList
}

// @leet end
