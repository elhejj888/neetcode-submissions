func twoSum(nums []int, target int) []int {
    history := make(map[int]int)
	for i:=0; i<len(nums); i++{
		res := target - nums[i]
		if v, ok := history[res]; ok {
			return []int{v, i}
		}
		history[nums[i]] = i
		
	}
	return []int{}
}
