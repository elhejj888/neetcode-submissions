func dailyTemperatures(temperatures []int) []int {
	index := 0
	nextIndex := 1
	warmerDates := []int{}
	n := len(temperatures)
	for index < n {
		if nextIndex == n {
			warmerDates = append(warmerDates,0)
			
			index++
			nextIndex = index+1
			
		} else if temperatures[nextIndex] > temperatures[index] && nextIndex < n{
			
			warmerDates = append(warmerDates,nextIndex - index)
			index++
			nextIndex = index + 1
			
		} else if temperatures[nextIndex] <= temperatures[index]  && nextIndex < n{
			nextIndex++
			
		} 

	}
	return warmerDates
}
