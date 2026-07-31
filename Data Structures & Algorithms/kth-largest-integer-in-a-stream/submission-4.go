import ("slices")
type KthLargest struct {
   K int
   Stream []int
}


func Constructor(k int, nums []int) KthLargest {
    return KthLargest{
		K : k,
		Stream : nums,
	}
}


func (this *KthLargest) Add(val int) int {
    this.Stream = append(this.Stream, val)
	slices.Sort(this.Stream)
	return this.Stream[len(this.Stream) - this.K]

}
