class Solution(object):
    def rotate(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: None Do not return anything, modify matrix in-place instead.
        """
        # 顺时针旋转90°: 先上下翻转，再沿主对角线翻转
        N = len(matrix)
        for m in range(N // 2):
            matrix[m], matrix[N - m - 1] = matrix[N - m - 1], matrix[m]
        # 或者 matrix.reverse()
        # 遍历左下角,和右上角换
        for m in range(N):
            for n in range(m):
                matrix[m][n], matrix[n][m] = matrix[n][m], matrix[m][n]

    def rotate2(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: None Do not return anything, modify matrix in-place instead.
        """
        N = len(matrix)
        l = N - 1  # 矩阵右边界

        # 如果 N = 5，行走2行，列走3列
        # A A A B B
        # A A A B B
        # D D · B B
        # D D C C C
        # D D C C C
        for m in range(N // 2):
            for n in range((N + 1) // 2):

                # 真正的 90 度旋转坐标 (行变列，列变行)
                (
                    matrix[n][l - m],  # 右上 (B) 接收 左上 (A)
                    matrix[l - m][l - n],  # 右下 (C) 接收 右上 (B)
                    matrix[l - n][m],  # 左下 (D) 接收 右下 (C)
                    matrix[m][n],  # 左上 (A) 接收 左下 (D)
                ) = (
                    matrix[m][n],  # 左上 (A) 给 右上 (B)
                    matrix[n][l - m],  # 右上 (B) 给 右下 (C)
                    matrix[l - m][l - n],  # 右下 (C) 给 左下 (D)
                    matrix[l - n][m],  # 左下 (D) 给 左上 (A)
                )
