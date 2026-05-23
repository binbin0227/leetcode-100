# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def diameterOfBinaryTree(self, root):
        """
        :type root: Optional[TreeNode]
        :rtype: int
        """
        self.res = 0

        def get_depth(node):
            if node is None:
                return 0
            left_depth = get_depth(node.left)
            right_depth = get_depth(node.right)

            # 任务一：算自己的“展臂长度”，挑战全公司纪录！
            dia_temp = left_depth + right_depth
            self.res = max(self.res, dia_temp)

            # 任务二：算自己的“单边最长腿”，汇报给上级大老板！
            return max(left_depth, right_depth) + 1

        get_depth(root)
        return self.res
