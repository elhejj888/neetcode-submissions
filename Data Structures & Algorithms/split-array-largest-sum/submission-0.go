import ("slices")
func splitArray(nums []int, k int) int {
	if k == len(nums) {
		return slices.Max(nums)
	}
	left := slices.Max(nums)
	right:= 0
	for _, value := range nums {
		right += value
	}
	for left < right {
		mid := left + (right - left) / 2



		if isMaxSplit(nums, k, mid) {
			right = mid
		} else {
			left = mid+1
		}
	 
	}
	return left
}

func isMaxSplit(nums []int, k int, max int) bool {
	count := 1
	currentSum := 0

	for i:= 0; i<len(nums); i++ {
		currentSum += nums[i]
		if currentSum > max {
			currentSum = nums[i]
			count++
		}
	}
	return count <= k
}
