class Solution(object):
    def productExceptSelf(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        n = len(nums)
        left_list = [nums[0]]
        for i in range(1, n):
            left_list.append(left_list[i - 1] * nums[i])
        right_list = [nums[-1]]
        for i in range(n - 1, 0,-1):
            right_list.insert(0,right_list[i + 1] * nums[i])
        result = []
        result.append(right_list[1])
        for i in range(1, n - 1):
            result.append(left_list[i - 1] * right_list[i + 1])
        return result
