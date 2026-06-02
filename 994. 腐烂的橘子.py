import collections


class Solution(object):
    def orangesRotting(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        m = len(grid)
        n = len(grid[0])
        q = collections.deque()
        fresh_count = 0
        minites = 0
        directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]

        for i in range(m):
            for j in range(n):
                if grid[i][j] == 2:
                    q.append((i, j))
                elif grid[i][j] == 1:
                    fresh_count += 1

        while q and fresh_count != 0:
            minites += 1
            for _ in range(len(q)):
                curr_i, curr_j = q.popleft()
                for x, y in directions:
                    temp_i, temp_j = curr_i + x, curr_j + y
                    if 0 <= temp_i < m and 0 <= temp_j < n and grid[temp_i][temp_j] == 1:
                        grid[temp_i][temp_j] = 2
                        fresh_count -= 1
                        q.append((temp_i, temp_j))
        return minites if fresh_count == 0 else -1
