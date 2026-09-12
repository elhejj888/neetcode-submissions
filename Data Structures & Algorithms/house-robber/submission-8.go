func rob(nums []int) int {
	prev1, prev2 := 0,0


	for i:=0; i<len(nums); i++ {
		curr := max(nums[i] + prev2, prev1)
		prev2 = prev1
		prev1 = curr
	}
	return prev1
}
