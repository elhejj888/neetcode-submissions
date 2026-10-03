import (
	"slices"
	"cmp"
	)
func longestCommonPrefix(strs []string) string {
    var pref strings.Builder
	slices.SortFunc(strs, func(a, b string)int{
		return cmp.Compare(len(a), len(b))
	})
	i := 0
	for i < len(strs[0]) {
		k := 1
		for k < len(strs) && strs[0][i] == strs[k][i]{
			k++
		}
		if k == len(strs) {
			fmt.Fprintf(&pref, "%c", strs[0][i])
		} else {
			return pref.String()
		}

	i++
	}
	return pref.String()
}

