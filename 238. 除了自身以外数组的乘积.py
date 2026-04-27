class Solution(object):
    def productExceptSelf(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        n = len(nums)
        L = [1] * n
        result = [1] * n
        
        # 1. 填充 L 数组 (L[i] 是 nums[i] 左侧所有元素的乘积，不包括L[i])
        # 对于 i = 0，左侧没有元素，L[0] 就是默认的 1
        for i in range(1, n):
            L[i] = L[i - 1] * nums[i - 1]
        right_product = 1
        
        # 2. right_product表示右侧的数的乘积，也就是R[i]
        for i in range(n - 1, -1, -1):
            result[i]=L[i]*right_product
            right_product *=nums[i]
        return result
    
def productExceptSelf2(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        # 1. 初始化预分配数组，默认全为 1，完美避开 insert 导致的性能问题
        n = len(nums)
        L = [1] * n
        R = [1] * n
        result = [1] * n
        
        # 2. 填充 L R 数组 (L[i] 是 nums[i] 左侧所有元素的乘积，不包括L[i])
        # 对于 i = 0，左侧没有元素，L[0] 就是默认的 1
        for i in range(1, n):
            L[i] = L[i - 1] * nums[i - 1]
        for i in range(n - 2, -1, -1):
            R[i] = R[i + 1] * nums[i + 1]
            
        # 3. 合并结果
        for i in range(n):
            result[i] = L[i] * R[i]
        return result