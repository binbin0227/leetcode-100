package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 反着来，把大元素装到后面
	p1, p2, curr := m-1, n-1, m+n-1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[curr] = nums1[p1]
			p1--
		} else {
			nums1[curr] = nums2[p2]
			p2--
		}
		curr--
	}

	for p1 >= 0 {
		nums1[curr] = nums1[p1]
		p1--
		curr--
	}
	for p2 >= 0 {
		nums1[curr] = nums2[p2]
		p2--
		curr--
	}
}
