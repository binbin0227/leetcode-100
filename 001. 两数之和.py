class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """

        d = {}  # 存该数字和它的索引
        for index, num in enumerate(nums):
            diff = target - num
            if diff in d:
                return [index,d[diff]]
            d[num]=index


s = Solution()
print(s.twoSum([2, 7, 11, 15], 9))
