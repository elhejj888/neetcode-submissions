type Heap []int
type KthLargest struct {
    k     int
    items Heap
}
func (h Heap) Len() int { return len(h)}
func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *Heap) Pop() any {
	old := *h 
	n := len(old)
	tmp := old[n-1]
	*h = old[:n-1]
	return tmp
}
func (h *Heap) Push(item any) {
	*h = append(*h, item.(int))
}
func Constructor(k int, nums []int) KthLargest {
    h := Heap{}
    heap.Init(&h)

    obj := KthLargest{
        k:     k,
        items: h,
    }

    for _, n := range nums {
        obj.Add(n)
    }

    return obj
}

func (this *KthLargest) Add(val int) int {

    if this.items.Len() < this.k {
        heap.Push(&this.items, val)
    } else if val > this.items[0] {
        heap.Pop(&this.items)
        heap.Push(&this.items, val)
    }

    return this.items[0]
}