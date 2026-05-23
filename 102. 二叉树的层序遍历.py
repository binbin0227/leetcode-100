# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
import collections


class Solution(object):
    def levelOrder(self, root):
        """
        :type root: Optional[TreeNode]
        :rtype: List[List[int]]
        """
        if root is None:
            return []
        res = []
        q = collections.deque([root])
        while q:
            # 题目要求每一层的数据各自用列表装
            level_size = len(q)
            level_res = []
            for _ in range(level_size):
                curr = q.popleft()
                level_res.append(curr.val)
                # 如果有左右节点，就放到队尾排队
                if curr.left:
                    q.append(curr.left)
                if curr.right:
                    q.append(curr.right)
            res.append(level_res)
        return res
