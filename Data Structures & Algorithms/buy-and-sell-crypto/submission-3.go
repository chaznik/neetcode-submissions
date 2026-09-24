func maxProfit(prices []int) int {
	left := 0
	right := 1
	maxProfit := 0

	for right < len(prices) {
		if prices[left] >= prices[right] {
			left = right
		} else {
			if (prices[right] - prices[left]) > maxProfit {
				maxProfit = prices[right] - prices[left]
			}
		}
		right++
	}
	return maxProfit
}
