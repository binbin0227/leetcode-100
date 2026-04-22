class Solution(object):
    def findAnagrams(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: List[int]
        """
        result = []
        p_dict = {}  # p的哈希表
        for char in p:
            if char not in p_dict:
                p_dict[char] = 1
            else:
                p_dict[char] += 1
        left = 0
        right = left + len(p) - 1
        s_window_dict = {}
        #初始化窗口哈希表并比较
        for char in s[left: right + 1]:
            if char not in s_window_dict:
                s_window_dict[char] = 1
            else:
                s_window_dict[char] += 1
        if s_window_dict==p_dict:
            result.append(left)
        #开始滑动：循环为每右滑一次就对比一次，因此刚才初始化后就要立刻对比，避免初始化窗口漏比较
        while right<len(s)-1:
            if s_window_dict[s[left]]==1:
                s_window_dict.pop(s[left])
            else:
                s_window_dict[s[left]]-=1
            left+=1
            right+=1
            if s[right] not in s_window_dict:
                s_window_dict[s[right]]=1
            else:
                s_window_dict[s[right]]+=1
            if s_window_dict==p_dict:
                result.append(left)
        return result

    # 暴力解法
    def findAnagrams_1(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: List[int]
        """
        result = []
        p_dict = {}
        for char in p:
            if char not in p_dict:
                p_dict[char] = 1
            else:
                p_dict[char] += 1
        left = 0
        right = left + len(p)
        s_window_dict = {}
        while right <= len(s):
            s_window = s[left:right]
            s_window_dict.clear()
            for char in s_window:
                if char not in s_window_dict:
                    s_window_dict[char] = 1
                else:
                    s_window_dict[char] += 1
            if s_window_dict == p_dict:
                result.append(left)
            left += 1
            right += 1
        return result
