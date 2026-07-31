func twoSum(nums []int, target int) []int {
    sumList := make(map[int]int)
	for indice, value := range(nums) {
		result := target - value
		if _, found := sumList[result]; found {
			return [] int {sumList[result], indice }
		}
		sumList[value] = indice
	}
	fmt.Println(sumList)
	return [] int{}
}
