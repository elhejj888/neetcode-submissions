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

func leastInterval(tasks []byte, n int) int {
    count := make(map[byte]int)
    for _, task := range tasks {
        count[task]++
    }

    maxHeap := Heap {}
	heap.Init(&maxHeap)
    for _, cnt := range count {
       heap.Push(&maxHeap, cnt)
    }

    time := 0
    q := make([][2]int, 0)

    for maxHeap.Len() > 0 || len(q) > 0 {
        time++

        if maxHeap.Len() == 0 {
            time = q[0][1]
        } else {
            cnt := heap.Pop(&maxHeap)
            cnt = cnt.(int) - 1
            if cnt.(int) > 0 {
                q = append(q, [2]int{cnt.(int), time + n})
            }
        }

        if len(q) > 0 && q[0][1] == time {
           heap.Push(&maxHeap,q[0][0])
            q = q[1:]
        }
    }

    return time
}