# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def isValidBST(self, root):
        """
        :type root: Optional[TreeNode]
        :rtype: bool
        """

        def check(node, lower_limit, upper_limit):
            if node is None:
                return True
            if node.val <= lower_limit or node.val >= upper_limit:
                return False
            is_left_valid = check(node.left, lower_limit, node.val)
            is_right_valid = check(node.right, node.val, upper_limit)
            return is_left_valid and is_right_valid

        return check(root, float("-inf"), float("inf"))
