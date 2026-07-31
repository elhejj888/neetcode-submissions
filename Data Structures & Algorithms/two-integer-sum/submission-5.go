func twoSum(nums []int, target int) []int {
    sumList := make(map[int]int)
	for indice, value := range(nums) {
		result := target - value
		if output, found := sumList[result]; found {
			return [] int {output, indice }
		}
		sumList[value] = indice
	}
	return [] int{}
}
