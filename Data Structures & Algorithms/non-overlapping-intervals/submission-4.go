
func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i,j int) bool {

		return intervals[i][1] <= intervals[j][1]
	})
	fmt.Println(intervals)

	count := 0
	last := 0


	for i:=1; i<len(intervals); i++{
		if intervals[i][0] < intervals[last][1]{
			count++
			continue
		} 
		last = i
	}
	return count
}
