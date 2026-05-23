# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def sortedArrayToBST(self, nums):
        """
        :type nums: List[int]
        :rtype: Optional[TreeNode]
        """
        if not nums:
            return None
        mid_index = len(nums) // 2
        root = TreeNode(nums[mid_index])  # type: ignore
        root.left = self.sortedArrayToBST(nums[:mid_index])
        root.right = self.sortedArrayToBST(nums[mid_index + 1 :])
        return root
