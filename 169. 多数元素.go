package main

func majorityElement(nums []int) int {
	res, count := 0, 0
	for _, num := range nums {
		if num == res {
			count++
		} else {
			if count == 0 {
				count++
				res = num
			} else {
				count--
			}
		}
	}
	return res
}
