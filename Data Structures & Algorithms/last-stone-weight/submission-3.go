type Heap []int

func (h Heap) Len()int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
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

func lastStoneWeight(stones []int) int {
	items := Heap{}
	heap.Init(&items)
	for _,n := range stones {
		heap.Push(&items, n)
	}

	for items.Len() > 0 {
		if items.Len() == 1 {
			return items[0]
		}

		stone1 := heap.Pop(&items).(int)
		stone2 := heap.Pop(&items).(int)
		if stone1 - stone2 > 0 {
			heap.Push(&items, stone1 - stone2)
		}
	}

	return 0

}
