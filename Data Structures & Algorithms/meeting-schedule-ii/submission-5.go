/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

type Heap []int

func (h Heap) Len()int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h Heap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Pop() any {
	
	old := *h
	tmp := old[len(old) - 1]
	*h = old[:len(old)-1]
	return tmp
} 

func (h *Heap) Push(item any) {
	*h = append(*h, item.(int))
}

func (h *Heap) Peek() int {
	return (*h)[0]
}
func minMeetingRooms(intervals []Interval) int {

	sort.Slice(intervals, func(i,j int) bool {
		return intervals[i].start < intervals[j].start
	})

	queue := Heap{}
	heap.Init(&queue)

	for _, interval := range intervals {
		if queue.Len() > 0 {
			if top := queue.Peek(); top <= interval.start {
				heap.Pop(&queue)
			}
		}
		heap.Push(&queue, interval.end)
	}
	return queue.Len()
}
