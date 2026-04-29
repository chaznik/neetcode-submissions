func maxProfit(prices []int) int {
	minBuy := prices[0]
	bestProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minBuy {
			minBuy = prices[i]
		} else {
			profit := prices[i] - minBuy
			if profit > bestProfit {
				bestProfit = profit
			}			
		}
	}
	return bestProfit
}