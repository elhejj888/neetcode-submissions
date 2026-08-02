
func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i,j int) bool {

		return intervals[i][1] <= intervals[j][1]
	})
	fmt.Println(intervals)

	count := 0
	lastIndex := 0

	for i:=1; i<len(intervals); i++{
		if intervals[lastIndex][1] <= intervals[i][0]{
			lastIndex = i
			
		} else {
		count++	
		}	
		
		
	}
	return count
}
