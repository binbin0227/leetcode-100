class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        # 未遇到0时快慢指针指向同一个数，一起向右移动；
        # 遇到0时慢指针停下，快指针指向下一个数，交换，慢指针右移，二者再次相等
        slow = 0  
        for fast in range(len(nums)):
            if nums[fast]!=0:
                nums[slow],nums[fast]=nums[fast],nums[slow]
                slow+=1


nums = [0, 1, 0, 3, 12]
s = Solution()
s.moveZeroes(nums)
print(nums)
