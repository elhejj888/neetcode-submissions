func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	nJumps := 1
	jumps := nums[0]

	for i := 1; i <= jumps; i++ {

		if jumps >= len(nums) - 1 || i == len(nums) - 1 {
			return nJumps
		}
		
		if i + nums[i] > jumps {
			nJumps++
			jumps += nums[i]
		}
	}

	return nJumps

    
}
