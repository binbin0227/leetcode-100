package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) != 0 {
		length := len(q)
		for i := range length {
			node := q[i]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, q[length-1].Val)
		q = q[length:]
	}

	return res
}
