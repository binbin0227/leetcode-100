class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        result = []
        nums.sort()
        for i in range(len(nums)):
            if nums[i] > 0:
                break
            if i>0 and nums[i] == nums[i-1]:
                continue
            left = i + 1
            right = len(nums) - 1
            while left < right:
                total = nums[i] + nums[left] + nums[right]
                if total < 0:
                    left += 1
                elif total > 0:
                    right -= 1
                else:
                    result.append([nums[i], nums[left], nums[right]])
                    # --- 核心去重逻辑 ---
                    # 只要左边下一个数和当前一样，就一直往右挪
                    while left < right and nums[left] == nums[left + 1]:
                        left += 1
                    # 只要右边下一个数和当前一样，就一直往左挪
                    while left < right and nums[right] == nums[right - 1]:
                        right -= 1
                    # ------------------
                    left += 1
                    right -= 1
        return result
s = Solution()
print(s.threeSum([-1,0,1,2,-1,-4]))
