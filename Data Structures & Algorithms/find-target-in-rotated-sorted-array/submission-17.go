func search(nums []int, target int) int {
	right := len(nums)
	left := 0

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else{
			return -1
		}
	}

	for left < right {
		mid := left + (right - left ) / 2

		fmt.Println(mid)
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > nums[left] {
			if target <= nums[mid] && target >= nums[left]{
				right = mid
			} else {
				left = mid +1
			}
		} else{
			if target >= nums[mid] && target <= nums[right - 1] {
				left = mid + 1
			} else {
				right = mid 
			}
		}
	}
	return -1

}