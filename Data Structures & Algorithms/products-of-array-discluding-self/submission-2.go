func productExceptSelf(nums []int) []int {
	var results []int
	zerosCount := 0
	res := 1
	for i:=0; i<len(nums); i++{
		if nums[i] == 0{
			zerosCount ++
		} else {
			res = res * nums[i]
		}
	}
	if zerosCount > 1 {
		return make([]int, len(nums))
	}  
	for i:= 0; i<len(nums);i++ {
		if nums[i] == 0 {
			results = append(results, res)
		} else if zerosCount == 0 {
		res_i := res / nums[i]
		results = append(results, res_i)
		} else if zerosCount == 1 && nums[i] != 0 {
			results = append(results, 0)
		}
	}
	return results
}
