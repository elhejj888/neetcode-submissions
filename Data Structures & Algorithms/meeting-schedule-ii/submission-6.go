/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	n := len(intervals)
    if n == 0 {
        return 0
    }
	// we separate the intervals on two tables, the 1st for the starts and the 2nd for the ends
    starts := make([]int, n)
    ends := make([]int, n)

	//a loop that collects all elements for each interval and assign each value to the right table
    for i, interval := range intervals {
        starts[i] = interval.start
        ends[i] = interval.end
    }

	//sorting both tables to get the first starting meeting fitting with the first ending time
    sort.Ints(starts)
    sort.Ints(ends)

	//two pointers to iterate each table on a different speed s: start, e:end
	sPtr,ePtr := 0,0

	//two variables to track both the number of rooms used after each iteration and the maximum number of rooms in total
	usedRooms, maxRooms := 0,0

	// a loop that iterates both tables, so whenever we have a starting meeting on a time smaller than the ending time of the current meeting, we will need to use another room for that meeting, else if the start time is bigger then a meeting has ended and a room is no longer used, so we decrement the number of rooms used, and we assign that number to the maxRooms  
	for sPtr< n{
		if starts[sPtr] < ends[ePtr] {
			usedRooms++
			sPtr++
		} else {
			usedRooms--
			ePtr++
		}

		if usedRooms > maxRooms {
			maxRooms = usedRooms
		}
	}

	return maxRooms
}
