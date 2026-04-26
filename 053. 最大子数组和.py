class Solution(object):
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        #只要前面的累加和是负数就扔掉、从当前数字重新开始算，否则就带着一起走，一路上随时记下出现过的最大值。
        result = pre = nums[0]
        for i in range(1,len(nums)):
            pre = max(nums[i],pre+nums[i])
            result=max(result,pre)
        return result
            