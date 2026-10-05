func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	maxProfit := 0
	leftPtr := 0
	rightPtr := 1

	for rightPtr < len(prices) {
		currentProfit := prices[rightPtr] - prices[leftPtr]
		if currentProfit <= 0 {
			leftPtr = rightPtr
		} else {
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
		rightPtr++
	}
	return maxProfit
}
