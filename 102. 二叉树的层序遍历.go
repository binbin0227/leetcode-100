package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) != 0 {
		length := len(q)
		currRes := make([]int, 0, length)

		for i := range length {
			currRes = append(currRes, q[i].Val)

			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}

		q = q[length:]
		res = append(res, currRes)
	}

	return res
}
