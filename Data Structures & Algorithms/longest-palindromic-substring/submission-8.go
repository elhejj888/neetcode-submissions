import (
	"slices"
	"cmp"
)
func longestPalindrome(s string) string {
	if len(s) <= 2{
		if len(s) == 2 && s[0] != s[1] {
			return string(s[0])
		} else if len(s) < 2 {
			return string(s[0])
		}
	}
	res := []string{}
for i:=0; i<len(s); i++{
	k, j, d := i, len(s)-1,0
	start,end := "",""
	for k<j {
		if s[i] == s[j] && start != ""{
			d = j
		}
		if s[k] == s[j]{
			start += string(s[k])
			end = string(s[j]) + end
			k++
			j--
		}else {
			j--
		}
		if s[k] != s[j] && start != ""{
			start = ""
			end =""
			k=i
			if d != 0 {
				j = d
				d = 0
			}
		}
		if k == j && start != "" {
			start += string(s[k])
		}
		
	}

	if start != "" {
		res = append(res, start+end)
	}
}

return slices.MaxFunc(res, func(a, b string)int{
	return cmp.Compare(len(a), len(b))
})
}
