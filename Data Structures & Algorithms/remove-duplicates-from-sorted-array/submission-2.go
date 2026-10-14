import "slices"
func removeDuplicates(nums []int) int {
	dup := make(map[int]bool)
	for i:=0; i<len(nums); i++{
		dup[nums[i]] = true 
	}
	i := 0
	for k,_ := range dup {
		nums[i] = k
		i++
	}
	nums = nums[:len(dup)]
	slices.Sort(nums)
	return len(dup)

}
