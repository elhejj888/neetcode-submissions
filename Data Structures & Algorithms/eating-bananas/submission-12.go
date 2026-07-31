import ("slices")
func minEatingSpeed(piles []int, h int) int {
	if h==len(piles){
		return slices.Max(piles)
	}
	
	left := 1
	right:= slices.Max(piles)
	var mid int

	for left < right{
		mid = (left+right) / 2
		if isWorking(piles, mid, h) {
			right = mid
		} else {
			left = mid+1
		}
	
	}
	return left
}

func isWorking(piles []int, nBananas int, h int) bool {
	count := 0
	for i := 0 ; i < len(piles); i++ {
		count += int(math.Ceil(float64(piles[i])/float64(nBananas)))
	}
	if count > h {
		return false
	}
	return true
}
