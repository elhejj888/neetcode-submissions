func search(nums []int, target int) int {

	right := len(nums)
	left := 0
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target{
			return mid
		}

	}
	return -1

	

}
