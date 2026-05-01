class Solution(object):
    def firstMissingPositive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        # 把数组本身当作哈希表
        n = len(nums)
        count = 0
        for i in range(n):
            # 只有数字1-n有位置坐；如果它在正确位置（数字3在索引2）就不用换位
            while 0 < nums[i] <= n and nums[nums[i] - 1] != nums[i]:
                correct_index = nums[i] - 1
                nums[i], nums[correct_index] = nums[correct_index], nums[i]
        # 换完座位后，再巡视一圈，如果谁没坐在该坐的位置上，那本该坐在这里的数字就是缺失的第一个正数
        for i in range(0, n):
            if nums[i] != i + 1:
                return i + 1
        # 如果大家极其完美地排排坐好了 (比如 [1, 2, 3])，那就缺下一个4
        return n + 1
