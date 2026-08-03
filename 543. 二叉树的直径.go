package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var res int

	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftDepth := getDepth(node.Left)
		rightDepth := getDepth(node.Right)

		// 边数 = 左深度 + 右深度，不需要 +1
		length := leftDepth + rightDepth
		res = max(res, length)
		return max(leftDepth, rightDepth) + 1
	}

	getDepth(root)

	return res
}
