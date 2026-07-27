package main

func productExceptSelf(nums []int) []int {
	lProduct := make([]int, len(nums))
	rProduct := make([]int, len(nums))
	lProduct[0] = 1
	rProduct[len(nums)-1] = 1

	lCurrProduct := 1
	for i := 1; i < len(nums); i++ {
		lCurrProduct *= nums[i-1]
		lProduct[i] = lCurrProduct
	}
	rCurrProduct := 1
	for i := len(nums) - 2; i >= 0; i-- {
		rCurrProduct *= nums[i+1]
		rProduct[i] = rCurrProduct
	}

	res := make([]int, len(nums))
	for i := range len(nums) {
		res[i] = lProduct[i] * rProduct[i]
	}

	return res
}
