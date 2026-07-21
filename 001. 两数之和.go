package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		another_idx, ok := m[target-num]
		if ok {
			return []int{idx, another_idx}
		} else {
			m[num] = idx
		}
	}
	return nil
}
