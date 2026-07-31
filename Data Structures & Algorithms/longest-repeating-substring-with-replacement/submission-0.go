func characterReplacement(s string, k int) int {
	count := make(map[string]int)
	left := 0
	best := 0
	maxfreq := 0
	for i := 0; i<len(s); i++ {
		count[string(s[i])]++

		if count[string(s[i])] > maxfreq {
			maxfreq = count[string(s[i])]
		}
		for (i - left+1) - maxfreq > k {
			count[string(s[left])]--
			left++
		}
		best = max(best, i-left+1)
	}
	return best
}




