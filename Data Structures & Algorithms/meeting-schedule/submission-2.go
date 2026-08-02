/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	sort.Slice(intervals, func(i,j int) bool{
		return intervals[i].start < intervals[j].end
	})
	lastIndex := 0
	fmt.Println(intervals)
	for i:=1; i<len(intervals); i++{
		if intervals[lastIndex].end > intervals[i].start {
			return false
		} else {
			lastIndex = i
		}
	}
	return true
}
