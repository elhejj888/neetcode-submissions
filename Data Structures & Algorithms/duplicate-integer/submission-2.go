func hasDuplicate(nums []int) bool {
    exist := make(map[int]int)
	for i := range nums {
		if _,ok := exist[nums[i]]; ok {
			return true
		}
		exist[nums[i]] = nums[i]
	}
	return false
}
