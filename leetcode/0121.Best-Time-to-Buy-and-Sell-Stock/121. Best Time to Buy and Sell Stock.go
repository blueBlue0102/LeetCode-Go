package leetcode

func BestTimetoBuyandSellStock(prices []int) int {
	lowest := 1<<31 - 1
	maxPrice := 0
	for _, price := range prices {
		lowest = min(lowest, price)
		maxPrice = max(maxPrice, price-lowest)
	}
	return maxPrice
}
