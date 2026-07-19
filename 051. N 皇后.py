class Solution(object):
    def solveNQueens(self, n):
        """
        :type n: int
        :rtype: List[List[str]]
        """
        res = []

        def check(path, row, col):
            for i in range(row):
                if path[i][col] == "Q":  # 检查上方
                    return False
            i, j, k = row - 1, col - 1, col + 1
            while i >= 0:
                if j >= 0 and path[i][j] == "Q":  # 检查左上
                    return False
                if k < n and path[i][k] == "Q":  # 检查右上
                    return False
                i -= 1
                j -= 1
                k += 1
            return True

        def backtrack(path, n, row):
            if row == n:
                res.append(path[:])
                return
            for col in range(n):
                if check(path, row, col):
                    path.append("." * col + "Q" + "." * (n - col - 1))
                    backtrack(path, n, row + 1)
                    path.pop()

        backtrack([], n, 0)
        return res
