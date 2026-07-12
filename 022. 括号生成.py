class Solution(object):
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        # 当前已经放进去的左括号数量，必须大于右括号数量
        self.res = []

        def backtrack(path, left, right):
            if len(path) == 2 * n:
                self.res.append(path)
                return
            if left < n:
                backtrack(path + "(", left + 1, right)
            if right < left:
                backtrack(path + ")", left, right + 1)

        backtrack("", 0, 0)
        return self.res
