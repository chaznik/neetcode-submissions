func maxProfit(prices []int) int {
	left := 0
	right := 1
	maxProfit := 0

	for right < len(prices) {
		if prices[left] >= prices[right] {
			left = right
		} else {
			profit := prices[right] - prices[left]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
		right++
	}
	return maxProfit
}
