import "slices"
func eraseOverlapIntervals(intervals [][]int) int {
    slices.SortFunc(intervals, func(a, b []int) int {
        if a[1]>=b[1] {
            return 1
        } else {
            return -1
        }
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
