package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 确保 nums1 是短数组
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	l, r := 0, m
	totalLeft := (m + n + 1) / 2

	for l <= r {
		// 确定 A B 切割点
		cutA := l + (r-l)/2
		cutB := totalLeft - cutA

		// 提取关键点，用无穷来处理越界
		L1 := math.MinInt64
		if cutA > 0 {
			L1 = nums1[cutA-1]
		}
		R1 := math.MaxInt64
		if cutA < m {
			R1 = nums1[cutA]
		}
		L2 := math.MinInt64
		if cutB > 0 {
			L2 = nums2[cutB-1]
		}
		R2 := math.MaxInt64
		if cutB < n {
			R2 = nums2[cutB]
		}

		if L1 > R2 {
			r = cutA - 1 // A 切太多了，往左缩
		} else if L2 > R1 {
			l = cutA + 1 // A 切太少了，往右扩
		} else {
			// 求出左侧最大值
			maxLeft := L1
			if L2 > maxLeft {
				maxLeft = L2
			}

			// 如果是奇数，中位数就是左侧最大值
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			// 如果是偶数，还要找右侧最小值
			minRight := R1
			if R2 < minRight {
				minRight = R2
			}
			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0.0
}
