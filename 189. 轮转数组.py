class Solution(object):
    def rotate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        k %= n

        def reverse(left, right):
            """反转闭区间 l-r"""
            while left < right:
                nums[left], nums[right] = nums[right], nums[left]
                left += 1
                right -= 1

        # 假设 nums = [1,2,3,4,5] ,k = 2
        reverse(0, n - 1)  # [5,4,3,2,1]
        reverse(0, k - 1)  # [4,5,3,2,1]
        reverse(k, n - 1)  # [4,5,1,2,3]

        # 时间O(n)解法：nums[:] = nums[n - k :] + nums[: n - k]
