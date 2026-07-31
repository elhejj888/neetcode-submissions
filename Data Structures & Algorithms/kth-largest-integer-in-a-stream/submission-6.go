type Heap []int
type KthLargest struct {
    K int
	Items Heap
}
func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int){
	h[i],h[j] = h[j], h[i]
}

func (h *Heap) Pop() any {
	
	old := *h
	tmp := old[len(old) - 1]
	*h = old[:len(old) - 1]
	return tmp
}

func (h *Heap) Push(item any){
	*h = append(*h, item.(int))
}

func Constructor(k int, nums []int) KthLargest {
 h := Heap{}
 heap.Init(&h)

 obj := KthLargest{
	K :k,
	Items: h,
 }

 for _, n := range nums{
	obj.Add(n)
 }

return obj
}

func (this *KthLargest) Add(val int) int {

	if this.Items.Len() < this.K {
		heap.Push(&this.Items, val)
	} else if val > this.Items[0] {
		heap.Pop(&this.Items)
		heap.Push(&this.Items, val)
	}

	return this.Items[0]

}
