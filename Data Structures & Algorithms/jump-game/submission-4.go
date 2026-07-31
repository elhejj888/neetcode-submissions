func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
    jumps := nums[0]
	for i := 1 ; i <= jumps ; i++  {
		if jumps >= len(nums) || i == len(nums) - 1 {
			return true
		}
		if i + nums[i] > jumps {
		jumps += nums[i]
		}
	}
	return false


}
