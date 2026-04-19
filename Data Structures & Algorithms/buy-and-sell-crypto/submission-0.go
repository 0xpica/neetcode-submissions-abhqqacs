func maxProfit(prices []int) int {
	tMax := 0 
	for i := 0 ; i < len(prices) - 1 ; i++ {
		for j := i + 1 ; j < len(prices); j++ {
			tMax = max(tMax, (prices[j] - prices[i]))
		}
	}
	return tMax
}
