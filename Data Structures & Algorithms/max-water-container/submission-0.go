func maxArea(heights []int) int {
	left, right := 0, len(heights) - 1
	best := 0

	for left < right {

		area := (right - left) * min(heights[left], heights[right])
		best = max (best, area)

		if heights[left] < heights[right]{
			left += 1
		} else {
			right -= 1
			}

	}
	return best

}
