func partitionLabels(s string) []int {
    lastIndex := make(map[byte]int)

	for i:= 0; i<len(s); i++{
		lastIndex[s[i]] = i
	}

	startIndex := 0
	endIndex := 0

	currentIndex := 0

	res := []int{}

	for currentIndex < len(s) {
		if endIndex < lastIndex[s[currentIndex]]{
			endIndex = lastIndex[s[currentIndex]]
		}
		currentIndex++

		if currentIndex > endIndex{
			res = append(res, endIndex - startIndex + 1)
			startIndex = currentIndex
			endIndex = currentIndex
		}
	}
	return res
}
