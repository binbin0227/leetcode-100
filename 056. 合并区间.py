import copy


class Solution(object):
    def merge(self, intervals):
        """
        :type intervals: List[List[int]]
        :rtype: List[List[int]]
        """
        #先排序，然后建一个新列表
        # 无重叠->放入
        # 有重叠->修改上一个进入的元素的右边界
        intervals.sort()
        result = [intervals[0]]
        for i in range(1,len(intervals)):
            curr = intervals[i]
            last_join = result[-1]
            if curr[0]<=last_join[1]:
                last_join[1]=max(last_join[1],curr[1])
            else:
                result.append(curr)
        return result