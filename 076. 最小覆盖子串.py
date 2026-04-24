class Solution(object):
    def minWindow(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: str
        """
        m = len(s)
        n = len(t)

        # 1. 初始化目标清单：记录字符串 t 中每个字符需要的数量
        t_dict = {}
        for char in t:
            t_dict[char] = t_dict.get(char, 0) + 1

        # 2. 初始化窗口工具：
        window_dict = {}  # 记录当前滑动窗口中各字符的数量
        valid = 0  # 记录窗口内有多少种字符已经达到了 t 中要求的数量
        min_start = 0  # 用于记录最终最短子串的起始索引
        left = 0  # 滑动窗口的左边界
        min_length = float("inf")  # 初始化最小长度为正无穷，方便后续通过对比更新最小值

        # 3. 开始移动右边界（寻宝阶段）
        # 窗口像蠕动的毛毛虫
        for right in range(m):
            char = s[right]

            # 如果当前字符是目标字符 t 中的，则更新窗口统计
            if char in t_dict:
                window_dict[char] = window_dict.get(char, 0) + 1
                # 关键：当某种字符的数量“刚好”达到目标要求时，有效计数 valid 加 1
                if window_dict[char] == t_dict.get(char, 0):
                    valid += 1

            # 4. 判断收缩条件：如果窗口已包含 t 中所有字符（valid 达标）
            while valid == len(t_dict):
                # 【记录阶段】：既然当前窗口匹配，先看它是不是比之前的记录更短
                if right - left + 1 < min_length:
                    min_start = left  # 锁定当前最短的起始点
                    min_length = right - left + 1  # 更新最短长度记录

                # 【收缩阶段】：尝试右移左边界 left，来压榨窗口空间
                out_char = s[left]
                if out_char in t_dict:
                    # 如果要移出的字符是关键字符，且移除后数量不再达标，则 valid 减 1
                    if window_dict[out_char] == t_dict[out_char]:
                        valid -= 1
                    window_dict[out_char] -= 1

                # 移动左指针，继续循环尝试下一次收缩
                left += 1

        # 5. 返回结果：如果 min_length 还是初始的无穷大，说明没找到；否则切片输出
        return (
            "" if min_length == float("inf") else s[min_start : min_start + min_length]
        )
