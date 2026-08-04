func findJudge(n int, trust [][]int) int {
	trusted := make(map[int][]int)
	counts := make(map[int]int)

	for i:=0; i<len(trust);i++{
		trusted[trust[i][0]] =append(trusted[trust[i][0]], trust[i][1])
		counts[trust[i][1]]++
	}

	for i:=1;  i <= n; i++{
		if _,found := trusted[i]; !found && counts[i] == n-1{
			return i
		}
	}
	return -1
}
