func minInterval(intervals [][]int, queries []int) []int {
    
	sort.Slice(intervals, func(i,j int)bool {
		return intervals[i][0] <= intervals[j][0]
	})

	fmt.Println(intervals)
	minInterval := -1
	
	res := make([]int,len(queries))
	for i:=0; i<len(queries); i++{

		for j:= 0; j<len(intervals); j++{
			if queries[i] <= intervals[j][1] && queries[i] >= intervals[j][0]{
				interval := intervals[j][1] - intervals[j][0] + 1
				if minInterval != -1 {
					if interval < minInterval {
						minInterval = interval
					}
				} else {
					minInterval = interval
				}
				fmt.Println(interval)
			} else if queries[i] < intervals[j][0]{
				break
			}
		}

		res[i] = minInterval
		minInterval = -1

	}
	return res
}
