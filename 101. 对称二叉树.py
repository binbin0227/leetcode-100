# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def isSymmetric(self, root):
        """
        :type root: Optional[TreeNode]
        :rtype: bool
        """

        def check(L, R):
            """检查LR子树是否对称"""
            if L is None and R is None:
                return True
            if L is None or R is None:
                return False
            if L.val != R.val:
                return False
            # 此层对称则继续检查下一层
            return check(L.left, R.right) and check(L.right, R.left)

        return check(root.left, root.right)
