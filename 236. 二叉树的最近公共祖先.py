# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def lowestCommonAncestor(self, root, p, q):
        """
        :type root: TreeNode
        :type p: TreeNode
        :type q: TreeNode
        :rtype: TreeNode
        """
        if not root:
            return None
        if root == p or root == q:
            return root

        left_report = self.lowestCommonAncestor(root.left, p, q)
        right_report = self.lowestCommonAncestor(root.right, p, q)

        if left_report and right_report:
            return root

        if not left_report:
            return right_report
        if not right_report:
            return left_report
