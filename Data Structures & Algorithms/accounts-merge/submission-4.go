import ("slices")
func accountsMerge(accounts [][]string) [][]string {
	parent := make(map[string]string)
	names := make(map[string]string)

	for i := 0; i < len(accounts); i++ {
		for j := 1; j < len(accounts[i]); j++ {
			email := accounts[i][j]
			parent[email] = email  // every email starts as its own parent
			names[email] = accounts[i][0]
		}
	}

	var find func(string)string
	find = func(acc string) string {
		if acc != parent[acc]{	
			parent[acc] = find(parent[acc])
		}
		return parent[acc]
	}

	for _,v := range accounts{
		for j:=1; j<len(v); j++ {
			a, b := find(v[j]), find(v[1])
			if a == b{
				continue
			}
			parent[a]=b

		}
	}

	
	groups := make(map[string][]string)
	for email := range parent {
		root := find(email)
		groups[root] = append(groups[root], email)
	}


	l:=[][]string{}
	for k,v := range groups {

		slices.Sort(v)
		row := append([]string{names[k]}, v...)
		l = append(l, row)
	}
	return l

}
