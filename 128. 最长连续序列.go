package main

func longestConsecutive(nums []int) int {
	// 1. 把所有数字放进集合
	// 2. 遍历集合中的每个数字
	// 3. 如果 num-1 存在，说明 num 不是起点，跳过
	// 4. 如果 num-1 不存在，从 num 开始不断检查 num+1、num+2……
	// 5. 记录最长长度

	var res int

	s := make(map[int]struct{})
	for _, num := range nums {
		s[num] = struct{}{}
	}

	for num := range s {
		if _, ok := s[num-1]; ok {
			continue
		}

		currNum := num
		for {
			if _, ok := s[currNum+1]; ok {
				currNum++
				continue
			}
			break
		}
		length := currNum - num + 1
		if length > res {
			res = length
		}
	}
	return res
}
