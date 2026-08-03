package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	res := make([]int, 0, k)
	stack := make([]*TreeNode, 0)
	curr := root

	for curr != nil || len(stack) != 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, curr.Val)
			curr = curr.Right
		}

		if len(res) == k {
			break
		}
	}

	return res[k-1]
}
