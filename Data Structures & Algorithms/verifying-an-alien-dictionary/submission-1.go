func isAlienSorted(words []string, order string) bool {
	mapping := make(map[byte]int)
	for i:=0; i<len(order); i++{
		mapping[order[i]] = i
	}

	for i := 0; i<len(words) - 1; i++{

		word1 := len(words[i])
		word2 := len(words[i+1])

		isBreak := false
		
		length := min(word1, word2)

		for j:=0; j<length; j++{
			if mapping[words[i][j]] == mapping[words[i+1][j]]{
				continue
			} else if mapping[words[i][j]] < mapping[words[i+1][j]]{
				isBreak = true
				break
			} else if mapping[words[i][j]] > mapping[words[i+1][j]]{
				return false
			}  

		}
		if !isBreak && word1 > word2{
			return false
		}

		isBreak = false
	}
	return true
}
