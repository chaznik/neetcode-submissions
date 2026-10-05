func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	currentMax := -9999999
	maxProfit := 0
	leftPtr := 0
	rightPtr := 1

	for rightPtr < len(prices) {
		if prices[leftPtr] > prices[rightPtr] {
			leftPtr = rightPtr
		} else {
			currentMax = prices[rightPtr] - prices[leftPtr]
			if currentMax > maxProfit {
				maxProfit = currentMax
			}
		}
		rightPtr++
	}
	return maxProfit
}
