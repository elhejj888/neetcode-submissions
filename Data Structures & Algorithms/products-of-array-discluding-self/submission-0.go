func productExceptSelf(nums []int) []int {
	var results []int
	for i:=0; i<len(nums); i++{
		res := 1;
		for j := 0; j<len(nums);j++ {
			if i != j {
				res = res*nums[j]
			}
		}
		results = append(results, res)
	}
	return results
}
