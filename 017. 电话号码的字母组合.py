class Solution(object):
    def letterCombinations(self, digits):
        """
        :type digits: str
        :rtype: List[str]
        """
        if not digits:
            return []

        phone_map = {
            "2": "abc",
            "3": "def",
            "4": "ghi",
            "5": "jkl",
            "6": "mno",
            "7": "pqrs",
            "8": "tuv",
            "9": "wxyz",
        }
        self.res = []

        def backtrack(path, index):
            if len(path) == len(digits):
                self.res.append("".join(path))
                return
            for char in phone_map.get(digits[index]):
                path.append(char)
                backtrack(path, index + 1)
                path.pop()

        backtrack([], 0)
        return self.res
