package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	totalLeft := (m + n + 1) / 2
	l, r := 0, m // cutA表示 nums1 的几个元素放到左边，l,r 是 cutA 的取值范围

	for l <= r {
		cutA := l + (r-l)/2
		cutB := totalLeft - cutA

		L1, L2 := math.MinInt, math.MinInt
		R1, R2 := math.MaxInt, math.MaxInt
		if cutA != 0 {
			L1 = nums1[cutA-1]
		}
		if cutA != m {
			R1 = nums1[cutA]
		}
		if cutB != 0 {
			L2 = nums2[cutB-1]
		}
		if cutB != n {
			R2 = nums2[cutB]
		}

		if L1 > R2 {
			r = cutA - 1
		} else if L2 > R1 {
			l = cutA + 1
		} else {
			leftMax := max(L1, L2)
			rightMin := min(R1, R2)
			if (m+n)%2 == 1 {
				return float64(leftMax)
			} else {
				return float64(leftMax+rightMin) / 2
			}
		}
	}

	return 0.0
}
