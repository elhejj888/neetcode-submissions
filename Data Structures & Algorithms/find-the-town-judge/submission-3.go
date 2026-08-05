func findJudge(n int, trust [][]int) int {
	trusted := make(map[int] bool)
	counts := make(map[int]int)

	for i:=0; i<len(trust);i++{
		trusted[trust[i][0]] = true
		counts[trust[i][1]]++
		
	}

	for i:=1;  i <= n; i++{
		if !trusted[i] && counts[i] == n-1{
			return i
		}
	}
	return -1
}
