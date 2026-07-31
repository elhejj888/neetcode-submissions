func dailyTemperatures(temperatures []int) []int {
	daysCount := 0
	index := 0
	nextIndex := 1
	warmerDates := []int{}
	n := len(temperatures)
	for index < n {
		if nextIndex == n {
			warmerDates = append(warmerDates,0)
			
			index++
			nextIndex = index+1
			daysCount = 0
		} else if temperatures[nextIndex] > temperatures[index] && nextIndex < n{
			daysCount++
			warmerDates = append(warmerDates,daysCount)
			index++
			nextIndex = index + 1
			daysCount = 0
		} else if temperatures[nextIndex] <= temperatures[index]  && nextIndex < n{
			nextIndex++
			daysCount++
		} 

	}
	return warmerDates
}
