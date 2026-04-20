class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        # 如果你移动长板：宽度变小，而高度受限于原有的短板，面积一定会减小。
        # 如果你移动短板：宽度虽然变小，但有可能遇到更长的板子，从而提升高度，面积有可能增大。
        n = len(height)
        left = 0
        right = n - 1
        max_volume = (right - left) * min(height[left] , height[right])
        while right > left:
            if height[left] > height[right]:
                right -= 1
            else:
                left += 1
            volume = (right - left) * min(height[left] , height[right])
            max_volume=max(max_volume,volume)
        return max_volume
s = Solution()
print(s.maxArea([1,8,6,2,5,4,8,3,7]))