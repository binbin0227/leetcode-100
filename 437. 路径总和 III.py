# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def pathSum(self, root, targetSum):
        """
        :type root: Optional[TreeNode]
        :type targetSum: int
        :rtype: int
        """
        # 类似 560 题
        self.pre_sum_dict = {0: 1}
        self.res = 0

        def dfs(node, pre_sum):
            if not node:
                return
            pre_sum += node.val

            # 查
            if pre_sum - targetSum in self.pre_sum_dict:
                self.res += self.pre_sum_dict[pre_sum - targetSum]

            # 记
            self.pre_sum_dict[pre_sum] = self.pre_sum_dict.get(pre_sum, 0) + 1

            # 继续往下查（不用判断 node.left存不存在，因为开头有 if not node 兜底防守）
            dfs(node.left, pre_sum)
            dfs(node.right, pre_sum)

            # 回溯
            self.pre_sum_dict[pre_sum] -= 1

        dfs(root, 0)
        return self.res
