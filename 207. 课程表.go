package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)  // 上完这个课能解锁什么课
	inDegree := make([]int, numCourses) // 这个课需要几门前置课

	for _, p := range prerequisites {
		course, pre := p[0], p[1]
		graph[pre] = append(graph[pre], course)
		inDegree[course]++
	}

	q := make([]int, 0)
	for i := range numCourses {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	count := 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:] // 上完课就出队
		count++

		for _, nextCourse := range graph[curr] {
			inDegree[nextCourse]--         // 他解锁的课，前置课的数量减少
			if inDegree[nextCourse] == 0 { // 如果减少到0了，就入队准备上课
				q = append(q, nextCourse)
			}
		}
	}

	return count == numCourses
}
