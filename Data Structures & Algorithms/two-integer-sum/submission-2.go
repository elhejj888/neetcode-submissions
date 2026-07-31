func twoSum(nums []int, target int) []int {
    sumList := make(map[int]int)
	for i := 0; i< len(nums); i++ {
		result := target - nums[i]
		if _, found := sumList[result]; found {
			return [] int {sumList[result], i}
		}
		sumList[nums[i]] = i
	}
	fmt.Println(sumList)
	return [] int{}
}
