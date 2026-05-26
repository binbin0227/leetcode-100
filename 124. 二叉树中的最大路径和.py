# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def maxPathSum(self, root):
        """
        :type root: Optional[TreeNode]
        :rtype: int
        """
        self.res = float("-inf")

        def dfs(node):
            if not node:
                return 0

            # 拿到左右部门的最强单线（如果是负的直接按 0 算，原地开除）
            left_max = max(dfs(node.left), 0)
            right_max = max(dfs(node.right), 0)

            # 算一下倒V拱桥，尝试更新全局变量
            self.res = max(self.res, left_max + right_max + node.val)

            # 返回一条最强单线
            return node.val + max(left_max, right_max)

        dfs(root)
        return self.res
