class Solution(object):
    def subarraySum(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        result = 0
        pre_sum = 0
        n = len(nums)
        pre_sum_dict = {0: 1}
        #为什么要加0:1
        # 假设 nums = [3, 4, 7], k = 3。如果没有 {0: 1}：
        # 遍历到 3：pre_sum = 3。寻找 target = 3 - 3 = 0。
        # 此时字典是空的，找不到 0，result 仍为 0。
        for num in nums:
            pre_sum += num
            if pre_sum - k in pre_sum_dict:
                #数组中有0或负数会导致pre_sum有多个相同
                result += pre_sum_dict[pre_sum - k]
            pre_sum_dict[pre_sum] = pre_sum_dict.get(pre_sum, 0) + 1
        return result


s = Solution()
print(s.subarraySum([1, 1, 1], 2))
