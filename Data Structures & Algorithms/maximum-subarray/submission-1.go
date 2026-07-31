import ("slices")
func maxSubArray(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	max := 0 
	table := [] int {}

	for i := 0 ; i < len(nums); i++ {

		max += nums[i]
		table = append(table, max)
		if max < 0{
		
		max = 0
		}

		
	}

	fmt.Println(table)
	return slices.Max(table)
}
