class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        left = 0
        max_str = 0
        window = set()  # window表示以s[right]结尾的最长无重复字符的串
        for right in range(len(s)):
            while s[right] in window:
                window.remove(s[left])
                left += 1
            window.add(s[right])
            max_str = max(max_str, len(window))
        return max_str


s = Solution()
print(s.lengthOfLongestSubstring("abcabcbb"))
