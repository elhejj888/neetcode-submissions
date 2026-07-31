func maxProfit(prices []int) int {
	cheapest := math.MaxInt64
	best := 0

	for _, value := range prices {
		cheapest = min(cheapest, value)
		best = max(best, value - cheapest)
	}
	fmt.Println(best)
	fmt.Println(cheapest)
	return best

}
