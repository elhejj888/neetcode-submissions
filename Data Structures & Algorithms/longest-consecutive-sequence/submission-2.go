import (
	"slices"
)
func longestConsecutive(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	slices.Sort(nums)
	res := 1
	maxN := 1
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i+1] - nums[i] == 1{
			res++
		} else if nums[i+1] - nums[i] != 1 && nums[i+1] != nums[i] && res != 1  {
			if res > maxN {
			maxN = res
			} 
			res = 1
		}
	} 
	return max(res, maxN)
	

}
