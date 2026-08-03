package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	res := math.MinInt

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		maxLeft := dfs(node.Left)
		maxRight := dfs(node.Right)

		res = max(res, maxLeft+maxRight+node.Val)

		return max(max(maxLeft, maxRight)+node.Val, 0)
	}

	dfs(root)
	return res
}
