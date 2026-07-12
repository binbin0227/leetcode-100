class Solution(object):
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        self.res = []

        def backtrack(path, start_index, curr_sum):
            if curr_sum > target:
                return
            if curr_sum == target:
                self.res.append(path[:])
                return

            for i in range(start_index, len(candidates)):
                num = candidates[i]
                path.append(num)
                backtrack(path, i, curr_sum + num)
                path.pop()

        backtrack([], 0, 0)
        return self.res
