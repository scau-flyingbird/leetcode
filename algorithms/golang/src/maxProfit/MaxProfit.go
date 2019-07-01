package maxProfit

func maxProfit(prices []int) int {
	sum := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			sum += prices[i+1] - prices[i]
		}
	}
	return sum
}
