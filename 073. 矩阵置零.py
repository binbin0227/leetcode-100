class Solution(object):
    def setZeroes(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: None Do not return anything, modify matrix in-place instead.
        """
        m = len(matrix)
        n = len(matrix[0])
        # 1. 先查底细：看看第一行和第一列原本有没有 0
        row0_has_zero = any(matrix[0][j] == 0 for j in range(n))
        col0_has_zero = any(matrix[i][0] == 0 for i in range(m))
        for i in range(1,m):
            for j in range(1,n):
                if matrix[i][j] == 0:
                    matrix[i][0] = matrix[0][j] = 0
        # 3. 读记事本：根据第一行和第一列的标记，把【剩下的部分】变 0
        for i in range(1, m):
            if matrix[i][0] == 0:
                for j in range(1, n):
                    matrix[i][j] = 0
        for j in range(1, n):
            if matrix[0][j] == 0:
                for i in range(1, m):
                    matrix[i][j] = 0
        # 4. 最后清算：根据最开始查的底细，决定要不要清空第一行和第一列
        if row0_has_zero:
            for j in range(n):
                matrix[0][j] = 0
        if col0_has_zero:
            for i in range(m):
                matrix[i][0] = 0

    def setZeroes2(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: None Do not return anything, modify matrix in-place instead.
        """
        # 记录要变成0的行和列
        m = len(matrix)
        n = len(matrix[0])
        m_zero = []
        n_zero = []
        for i in range(m):
            for j, num in enumerate(matrix[i]):
                if num == 0:
                    m_zero.append(i)
                    n_zero.append(j)
        for i in m_zero:
            for j in range(n):
                matrix[i][j] = 0
        for j in n_zero:
            for i in range(m):
                matrix[i][j] = 0
