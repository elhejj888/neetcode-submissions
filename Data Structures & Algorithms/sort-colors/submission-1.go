func sortColors(nums []int) {
	left, right := 0, len(nums) - 1
	index := 0
	for index <= right {
		if nums[index] == 0 && nums[left] == 0{
			index++
			left++
		} else if nums[index] == 0 {
			nums[index], nums[left] = nums[left], nums[index]
			left++
		} else if nums[index] == 2 && nums[right] == 2 {
			right--
		} else if nums[index] == 2 {
			nums[index], nums[right] = nums[right], nums[index]
			right--
		} else {
			index++
		}
	} 
}
