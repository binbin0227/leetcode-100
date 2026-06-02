import collections


class Solution(object):
    def numIslands(self, grid):
        """
        :type grid: List[List[str]]
        :rtype: int
        """
        self.res = 0
        self.m = len(grid)
        self.n = len(grid[0])

        def dfs_clean(grid, i, j):
            if (
                i < 0 or i > self.m - 1 or j < 0 or j > self.n - 1 or grid[i][j] == "0"
            ):  # 必须先检查边界，再查值
                return
            grid[i][j] = "0"
            dfs_clean(grid, i + 1, j)
            dfs_clean(grid, i - 1, j)
            dfs_clean(grid, i, j + 1)
            dfs_clean(grid, i, j - 1)

        for i in range(self.m):
            for j in range(self.n):
                if grid[i][j] == "1":
                    self.res += 1
                    dfs_clean(grid, i, j)
        return self.res

    def numIslands2(self, grid):
        """
        :type grid: List[List[str]]
        :rtype: int
        """
        self.res = 0
        self.m = len(grid)
        self.n = len(grid[0])

        def bfs_clean(grid, i, j):
            q = collections.deque([(i, j)])
            grid[i][j] = "0"
            while q:
                curr_i, curr_j = q.popleft()

                # 向四个方向探查邻居
                directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]
                for x, y in directions:
                    temp_i, temp_j = curr_i + x, curr_j + y
                    if (
                        0 <= temp_i < self.m
                        and 0 <= temp_j < self.n
                        and grid[temp_i][temp_j] == "1"
                    ):
                        grid[temp_i][temp_j] = "0"
                        q.append((temp_i, temp_j))

        for i in range(self.m):
            for j in range(self.n):
                if grid[i][j] == "1":
                    self.res += 1
                    bfs_clean(grid, i, j)
        return self.res
