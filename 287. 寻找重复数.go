package main

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}

	slow = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}

	return slow
}
