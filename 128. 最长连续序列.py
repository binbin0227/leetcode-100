class Solution(object):
    def longestConsecutive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        nums_set = set(nums)
        length_list = [0]
        for num in nums_set:
            if num - 1 not in nums_set:  # 如果num不是序列第一位，就从它开始往后找
                length = 1
                while num + 1 in nums_set:
                    length += 1
                    num += 1
                length_list.append(length)
        return max(length_list)


s = Solution()
print(s.longestConsecutive([100, 4, 200, 1, 3, 2]))
