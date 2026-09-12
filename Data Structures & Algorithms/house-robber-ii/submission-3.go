func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
    prev1, prev2 := 0,0
	l1 := 0
	for i:=0; i<len(nums) - 1; i++ {
		prev2, prev1 = prev1, max(prev2+nums[i], prev1)
	}
	l1 = prev1
	prev1,prev2 = 0,0

	for i:=1; i<len(nums); i++ {
		prev2, prev1 = prev1, max(prev2+nums[i], prev1)
	}
	
	return max(prev1,l1)
}
