package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	curr := preorder[0]
	var midIdx int
	for midIdx = range len(inorder) {
		if inorder[midIdx] == curr {
			break
		}
	}
	return &TreeNode{
		Val:   curr,
		Left:  buildTree(preorder[1:midIdx+1], inorder[:midIdx]),
		Right: buildTree(preorder[midIdx+1:], inorder[midIdx+1:]),
	}
}
