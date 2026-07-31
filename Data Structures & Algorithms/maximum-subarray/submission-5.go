func maxSubArray(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	max := 0 
	topMax := math.MinInt64

	for i := 0 ; i < len(nums); i++ {

		max += nums[i]
		if max > topMax{
			topMax = max
		}
		if max < 0{
		
		max = 0
		}

		
	}
	return topMax
}
