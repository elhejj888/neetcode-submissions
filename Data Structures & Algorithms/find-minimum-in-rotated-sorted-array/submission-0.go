func findMin(nums []int) int {
	leftPointer := 0
	rightPointer := len(nums) - 1
	res := nums[leftPointer]
	for leftPointer < rightPointer {
		if nums[leftPointer] < nums[rightPointer] {
			res = nums[leftPointer]
			rightPointer--
		}else if nums[leftPointer] > nums[rightPointer] {
			res = nums[rightPointer]
			leftPointer++
		}else if nums[leftPointer] == nums[rightPointer]{
			leftPointer++
			rightPointer--
		}
	}
	return res
}
