import (
	"slices"
	"cmp"
)

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	res := []int{}
	for i:=0; i<len(nums); i++{
			count[nums[i]]++

	}
	values := []int{}
	for _,v := range count {
		values = append(values, v)
	}

	slices.SortFunc(values, func(a, b int) int {
        return cmp.Compare(b, a)
    })
	for i:=0; i<k; i++ {
		for j:=0; j<len(nums); j++ {
			if count[nums[j]] == values[i]{
				res = append(res, nums[j])
				count[nums[j]] = -1
			}
		}
	}

	return res

}
