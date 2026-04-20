func maxProfit(prices []int) int {
	l , r := 0 , 1 
	res := 0 
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l] 
			if profit > res {
				res = profit
			}
		}else {
			l = r 
		}
		r++
	}
	return res
}
