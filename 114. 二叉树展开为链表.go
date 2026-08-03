package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	var pre *TreeNode
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if pre != nil { // 保护循环开始时不崩溃
			pre.Right = node
			pre.Left = nil
		}

		pre = node
	}

	return
}
