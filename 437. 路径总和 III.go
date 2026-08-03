package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	preSumMap := map[int]int{0: 1} //

	var dfs func(node *TreeNode, preSum int)
	dfs = func(node *TreeNode, preSum int) {
		if node == nil {
			return
		}

		currSum := preSum + node.Val
		if count, ok := preSumMap[currSum-targetSum]; ok {
			res += count
		}

		preSumMap[currSum] += 1
		dfs(node.Left, currSum)
		dfs(node.Right, currSum)
		preSumMap[currSum] -= 1
	}

	dfs(root, 0)
	return res
}
