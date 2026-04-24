from collections import deque


class Solution(object):
    def maxSlidingWindow(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """
        result = []
        q = deque()
        for index, num in enumerate(nums):
            # 1. 劝退：保持单调递减
            while q and nums[q[-1]] <= num:
                q.pop()
            q.append(index)
            # 2. 过期：检查队首下标是否已经超出了窗口范围
            # 窗口范围是 [index-k+1, index]，所以下标必须 > index-k
            if q[0]==index - k:
                q.popleft()
            # 3. 记录：只要窗口形成了，队首就是当前窗口的最大值
            if index >= k - 1:
                result.append(nums[q[0]])
        return result
