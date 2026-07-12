class Solution(object):
    def exist(self, board, word):
        """
        :type board: List[List[str]]
        :type word: str
        :rtype: bool
        """
        m = len(board)
        n = len(board[0])

        def dfs(idx, i, j):
            if not 0 <= i < m or not 0 <= j < n or board[i][j] != word[idx]:
                return False
            # 运行到这说明第 idx 个字母也找到了

            if idx == len(word) - 1:  # 找到最后一个字母
                return True

            temp = board[i][j]
            board[i][j] = "#"  # 防止再次找走过的格子

            # 分别继续找周围四格
            res = (
                dfs(idx + 1, i + 1, j)
                or dfs(idx + 1, i - 1, j)
                or dfs(idx + 1, i, j + 1)
                or dfs(idx + 1, i, j - 1)
            )

            board[i][j] = temp  # 把字母变回去

            return res

        for i in range(m):
            for j in range(n):
                if board[i][j] == word[0]:
                    if dfs(0, i, j):
                        return True

        return False
