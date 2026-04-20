from collections import defaultdict


class Solution(object):
    def groupAnagrams(self, strs):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """
        #错误写法：
        # d = {}
        # for word in strs:
        #     word_set = set(word)
        #     word_set_str = str(word_set)
        #     if word_set_str in d:
        #         d[word_set_str].append(word)
        #     else:
        #         d[word_set_str]=[word]
        # return d.values()
        
        #及格写法：
        # d = {}
        # for word in strs:
        #     key="".join(sorted(word))
        #     if key in d:
        #         d[key].append(word)
        #     else:
        #         d[key]=[word]
        # return d.values()
    
    #最优写法：
        d = defaultdict(list)
        for word in strs:
            key="".join(sorted(word))
            if key in d:
                d[key].append(word)
            else:
                d[key]=[word]
        return d.values()
    
s = Solution()
print(s.groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]))