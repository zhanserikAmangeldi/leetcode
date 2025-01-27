package main

func maxProfit(prices []int) int {
    maxProfit := 0
    cheapest := prices[0]

	for _, value := range prices {
		if value < cheapest {
			cheapest = value
		}

		maxProfit = max(value-cheapest, maxProfit)
	}

	return maxProfit
}