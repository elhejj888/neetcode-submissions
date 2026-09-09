import ("slices")
func accountsMerge(accounts [][]string) [][]string {
	parent := make(map[string]string)
	names := make(map[string]string)
	seen := make(map[string]bool)

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
			acc = find(parent[acc])
		}
		return acc
	}

	res:=make(map[string][]string)

	for _,v := range accounts{
		for j:=1; j<len(v); j++ {
			a, b := find(v[j]), find(v[1])
			if a == b{
				acc := v[j]
					if !seen[acc] {
					res[parent[acc]] = append(res[parent[acc]] , acc)
					seen[acc] = true
					}
				continue
			}
			parent[a]=b

		}
	}

	
	for k,_ := range parent {

			parent[k] = find(k)
			
	}
	parent2 := make(map[string][]string)
	for k,v := range parent {
		parent2[v] = append(parent2[v],k)
	}


	l:=[][]string{}
	j:=0
	for k,v := range parent2 {
		l = append(l, []string{})
		l[j] = append(l[j], names[k])
		slices.Sort(v)
		l[j] = slices.Concat(l[j],v)
		j++
	}
	return l

}
