package main

func canJump(nums []int) bool {
	// 贪心：遍历每一格并维护目前可到达的最远的位置
	// 如果当前位置不可到达，返回false；维护最远可到达位置；再看可到达位置是否超过了目标位置
	n := len(nums)
	var maxReach int
	for i := range n {
		if maxReach < i {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= n-1 {
			return true
		}
	}
	return false
}

func canJump2(nums []int) bool {
	// BFS
	q := make([]int, 0, len(nums))
	q = append(q, 0)

	visited := make([]bool, len(nums))
	visited[0] = true

	for len(q) != 0 {
		curr := q[0]

		if curr >= len(nums)-1 {
			return true
		}

		for i := curr + 1; i <= curr+nums[curr]; i++ {
			if i < len(nums) && visited[i] == false { // 防越界
				visited[i] = true
				q = append(q, i)
			}
		}

		q = q[1:]
	}
	return false
}
