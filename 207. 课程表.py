import collections


class Solution(object):
    def canFinish(self, numCourses, prerequisites):
        """
        :type numCourses: int
        :type prerequisites: List[List[int]]
        :rtype: bool
        """
        in_degree = [0] * numCourses  # 记录课程需要几节前置课
        adj = collections.defaultdict(list)  # 记录学完这门课程能解锁什么课
        for target, pre in prerequisites:
            adj[pre].append(target)
            in_degree[target] += 1

        # 无前置课程的课入队
        q = collections.deque()
        for i in range(numCourses):
            if in_degree[i] == 0:
                q.append(i)
        learned_count = 0

        # 开始上课，出队的课就是上完的
        while q:
            curr_course = q.popleft()
            learned_count += 1
            for next_course in adj[curr_course]:
                in_degree[next_course] -= 1  # 更新后续课程所需的前置课数量
                if in_degree[next_course] == 0:  # 如果前置课上完了，就入队
                    q.append(next_course)
        return learned_count == numCourses
