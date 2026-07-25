package main

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	var res int
	for _, currPrice := range prices {
		if currPrice < minPrice {
			minPrice = currPrice
			continue
		}
		if currPrice-minPrice > res {
			res = currPrice - minPrice
		}
	}
	return res
}
