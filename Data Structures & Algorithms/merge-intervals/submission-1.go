func merge(intervals [][]int) [][]int {

    res := [][]int{}
	

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	interval := intervals[0]
	fmt.Println(intervals)

	for i:= 1; i<len(intervals) ; i++ {
		if interval[1] < intervals[i][0]{
			res = append(res, interval)
			interval = intervals[i]
		} else if interval[1] >= intervals[i][0]{
			interval[0] = min(interval[0],intervals[i][0])
			interval[1] = max(interval[1], intervals[i][1])
		}
	}
	
	res = append(res, interval)
	
	return res
}
