func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sums := make(map[int]int)

	sums[0] = nums[0]
	sums[1] = max(nums[0], nums[1])
	maxN := sums[1]
	for i:=2; i<len(nums); i++ {
		sums[i] = max(nums[i] + sums[i-2] , sums[i-1])
		maxN = sums[i]
	}
	//fmt.Println(sums)
	return maxN
}
