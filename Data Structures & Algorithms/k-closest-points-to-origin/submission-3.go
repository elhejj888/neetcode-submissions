type Heap [][]int

func (h Heap) Len()int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	r1 := math.Sqrt(math.Pow( float64(h[i][0]), 2) + math.Pow(float64(h[i][1]) , 2))
	r2 := math.Sqrt(math.Pow(float64(h[j][0]), 2) + math.Pow(float64(h[j][1]) , 2))
	return  r1 < r2
}
func (h Heap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Pop() any {
	
	old := *h
	tmp := old[len(old) - 1]
	*h = old[:len(old)- 1]
	return tmp
} 

func (h *Heap) Push(item any) {
	*h = append(*h, item.([]int))
}

func (h *Heap) Peek() []int {
	return (*h)[0]
}



func kClosest(points [][]int, k int) [][]int {
	res := Heap{}
	heap.Init(&res)
	res2 := [][]int{}
	for _,v := range points {
		heap.Push(&res, v)
	}
	fmt.Println(res)
	i := 0
	for i<k{
		
		res2 = append(res2,heap.Pop(&res).([]int))
		i++
	}
	return res2
}
