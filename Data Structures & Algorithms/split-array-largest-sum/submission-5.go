import ("slices")
func splitArray(nums []int, k int) int {
	if k == len(nums) {
		return slices.Max(nums)
	}
	left := nums[0]
	right:= 0
	for _, value := range nums {
		right += value
		left = max(left, value)
	}
	for left < right {
		mid := left + (right - left) / 2

		isTrue := isMaxSplit(nums, k, mid)

		if isTrue {
			right = mid
		} else {
			left = mid+1
		}
	 
	}
	return left
}

func isMaxSplit(nums []int, k int, limit int) bool {
    count := 1
    currentSum := 0

    for _, num := range nums {
        if currentSum+num > limit {
            count++

            if count > k {
                return false
            }

            currentSum = num
        } else {
            currentSum += num
        }
    }

    return true
}
