class Solution(object):
    def trap(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        result = 0
        left = 0
        right = len(height)-1
        max_left = height[left]
        max_right = height[right]
        while left<right:
            max_left = max(max_left,height[left])
            max_right = max(max_right,height[right])
            if max_left<max_right:
                result += max_left - height[left]
                left+=1
            else:
                result += max_right - height[right]
                right-=1
        return result
s = Solution()
print(s.trap([0,1,0,2,1,0,1,3,2,1,2,1]))