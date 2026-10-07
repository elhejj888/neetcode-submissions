import "slices"
func threeSum(nums []int) [][]int {
	res := make(map[int]int)
	ver := make(map[string]bool)
	for i := range nums {
		res[nums[i]] = i
	} 
	l := [][]int{}
	for i := 0; i < len(nums); i++{

		for j:=i+1; j<len(nums); j++{
			tmp := nums[i] + nums[j]

			tmp = tmp * -1

			if v, ok := res[tmp]; ok && v > j && v > i {
				arr := []int{nums[i],nums[j],tmp}
				slices.Sort(arr)
		if !ver[strconv.Itoa(arr[0])+","+strconv.Itoa(arr[1])+","+strconv.Itoa(arr[2])]{
							ver[strconv.Itoa(arr[0])+","+strconv.Itoa(arr[1])+","+strconv.Itoa(arr[2])] = true
				l = append(l, arr)
		}

			}
		}

	}
	return l
}
