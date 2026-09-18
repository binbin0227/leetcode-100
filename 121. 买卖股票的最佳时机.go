func maxProfit(prices []int) int {
	res := 0
	minPrice := prices[0]

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}

		if price-minPrice > res {
			res = price - minPrice
		}
	}

	return res
}