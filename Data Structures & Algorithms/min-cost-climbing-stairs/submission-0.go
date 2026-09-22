func minCostClimbingStairs(cost []int) int {
    cost = append(cost, 0)
	res := make([]int, len(cost))
	res[0] = cost[0]
	res[1] = cost[1]
	for i:=2; i<len(cost);i++{
		res[i] = min(cost[i] + res[i-1], cost[i] + res[i-2])
	}
	fmt.Println(res)
	return res[len(res)-1]

}
