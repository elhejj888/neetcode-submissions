func leastInterval(tasks []byte, n int) int {
	if tasks == nil {
		return 0
	}

	mapCounts := make(map[byte] int)

	for _, val := range tasks {
		mapCounts[val]++
	}

	maxVal := 0

	for _, val := range mapCounts{
		if val > maxVal{
			maxVal = val
		}
	}
	countMaxVal := 0
	for _, val := range mapCounts{
		if maxVal == val {
			countMaxVal++
		}
	}
	res := (maxVal - 1) * (n + 1) + countMaxVal

	return max(res, len(tasks))  

}
