package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := make([][]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	flag := false

	for len(stack) != 0 {
		length := len(stack)
		temp := make([]int, 0)
		for i := 0; i < length; i++ {
			curr := stack[i]
			temp = append(temp, curr.Val)
			if curr.Left != nil {
				stack = append(stack, curr.Left)
			}
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
		}

		if flag {
			l, r := 0, len(temp)-1
			for l < r {
				temp[l], temp[r] = temp[r], temp[l]
				l++
				r--
			}
		}

		res = append(res, temp)
		stack = stack[length:]

		flag = !flag
	}

	return res
}