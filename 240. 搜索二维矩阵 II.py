class Solution(object):
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        # 从矩阵的右上角开始搜索，比target小就往左，比target大就往下，越界则不存在target
        m = len(matrix) - 1
        n = len(matrix[0]) - 1
        i = 0
        j = n
        while 0 <= i <= m and 0 <= j <= n:
            if target > matrix[i][j]:
                i += 1
            elif target < matrix[i][j]:
                j -= 1
            else:
                return True
        return False
