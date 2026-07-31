func partitionLabels(s string) []int {
    frq := make(map[byte]int)
	for i := 0; i< len(s);i++{
		frq[s[i]]++
	}

	set := make(map[byte]struct{})
	res := []int{}
	start := 0
	for i:=0; i<len(s); i++ {
		frq[s[i]]--
		set[s[i]]=struct{}{}
		if frq[s[i]] == 0{
			delete(set,s[i])
		}
		if len(set) == 0 {
			res = append(res, i - start + 1)
			start = i+1
		}
	
	}
	return res
}
