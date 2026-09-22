func minCostClimbingStairs(cost []int) int {
	curr,prev := cost[1],cost[0]
	for i:=2; i<len(cost);i++{
		curr,prev = min(curr + cost[i], prev + cost[i]), curr

	}
	return min(curr,prev)

}
