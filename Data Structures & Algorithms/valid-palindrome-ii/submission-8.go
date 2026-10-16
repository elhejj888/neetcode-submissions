import "slices"
func validPalindrome(s string) bool {
	if len(s) < 3 {
		return true
	}

	var checkPalindrome func(string) bool
	checkPalindrome = func(s string) bool {
		var k []rune
		if len(s) % 2 == 0 {
		k = []rune (s[len(s)/2:])
		} else {
			k = []rune (s[(len(s)/2)+1:])
		}
		slices.Reverse(k)

		if string(s[:(len(s)/2)]) == string(k) {
			return true
		}
		return false
	}

	first, last := 0, len(s)-1
	for first < last {
		if s[first] != s[last] {
			var sb strings.Builder
			sb.WriteString(s[:first])
			sb.WriteString(s[first+1:])
			if checkPalindrome(sb.String()) {
				return true
			}
			sb.Reset()
			sb.WriteString(s[:last])
			sb.WriteString(s[last+1:])
			return checkPalindrome(sb.String())
		}
		first++
		last--
	}
	first, last = 0, len(s)-1

	


	return true

}
