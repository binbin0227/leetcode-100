class Solution(object):
    def partition(self, s):
        """
        :type s: str
        :rtype: List[List[str]]
        """
        res = []

        def backtrack(path, start_idx):
            if start_idx == len(s):
                res.append(path[:])
                return

            for i in range(start_idx, len(s)):
                curr = s[start_idx : i + 1]
                if curr == curr[::-1]:
                    path.append(curr)
                    backtrack(path, i + 1)
                    path.pop()

        backtrack([], 0)
        return res
