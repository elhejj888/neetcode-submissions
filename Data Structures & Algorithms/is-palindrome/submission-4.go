func isPalindrome(s string) bool {
	stripedS := ""
	s = strings.ToLower(s)
	for i := 0; i<len(s); i++ {
		if int(s[i]) > 122 || int(s[i]) < 65{
			if int(s[i]) >= 48 && int(s[i]) <= 57 {
				stripedS+=string(s[i])
			} else {
				continue
			}
		}else {
			stripedS+=string(s[i])
		}
	}
	fmt.Println(stripedS)
	for i , j:= 0, len(stripedS)-1; i < j ; i ,j= i+1, j-1 {
		if stripedS[i] != stripedS[j]{
			return false
		}

	}
	return true
}
