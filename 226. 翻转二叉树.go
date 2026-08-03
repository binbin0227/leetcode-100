/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	reverse(root)
	return root
}

func reverse(node *TreeNode) {
	if node == nil {
		return
	}

	node.Left, node.Right = node.Right, node.Left
	reverse(node.Left)
	reverse(node.Right)
}