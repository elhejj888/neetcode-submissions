func validTree(n int, edges [][]int) bool {
    //rows := len(edges)
	seen := make(map[int]bool)
	adj := make(map[int][]int)
	for _,v := range edges {
		adj[v[0]] = append(adj[v[0]],v[1])
		adj[v[1]] = append(adj[v[1]],v[0])
	}
	var dfs func(int, int)bool
	dfs = func(node, parent int) bool {
		if seen[node]{
			return false
		}
		seen[node] = true
		for _,v := range adj[node]{
			if v == parent {
				continue
			}
			if !dfs(v,node) {
				return false
			}
		}
		return true
	}
	if !dfs(0,-1){
		return false
	}
	return len(seen) == n
}
