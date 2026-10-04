func majorityElement(nums []int) int {
max := 0
val := len(nums)/2
counter := make(map[int]int)
for i:=0; i<len(nums); i++ {
	counter[nums[i]]++
	if v,_ := counter[nums[i]]; v > val {
		if v > max {
			max = nums[i]
		}
	}
}   
return max
}
