func mergeAlternately(word1 string, word2 string) string {
	k := max(len(word1), len(word2))
	i:=0
	var res strings.Builder
	for i < k {
		if i < len(word1) && i < len(word2){
			res.WriteString(string(word1[i]))
			res.WriteString(string(word2[i]))
			i++
		} else {
			if i >= len(word1) {
				res.WriteString(string(word2[i]))
				i++
			} else if i >= len(word2) {
				res.WriteString(string(word1[i]))
				i++
			}
		}

	}
	return res.String()
}
