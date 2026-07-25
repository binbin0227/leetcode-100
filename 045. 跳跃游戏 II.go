package main

func jump(nums []int) int {
	// 先把第一个位置放入队列；visited 表存放到达该索引的最短步数
	// 队头出队，它能到达的位置如果没填过 visited 表，就把它放入队并填表
	// 填表时的距离必定是最短距离

	n := len(nums)
	q := make([]int, 0, n)
	q = append(q, 0)
	visited := make([]int, n)

	for len(q) != 0 {
		curr := q[0]
		q = q[1:]
		for i := curr + 1; i <= curr+nums[curr]; i++ {
			if i >= n || visited[i] != 0 {
				continue
			}
			visited[i] = visited[curr] + 1
			q = append(q, i)
		}
		if visited[n-1] != 0 {
			return visited[n-1]
		}
	}
	return 0
}
