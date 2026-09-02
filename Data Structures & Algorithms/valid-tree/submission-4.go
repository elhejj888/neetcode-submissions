func validTree(n int, edges [][]int) bool {

	if len(edges) != n - 1 {
		return false
	}
	seen := make(map[int]bool)
	adj := make(map[int][]int)
	for _,v := range edges {
		adj[v[0]] = append(adj[v[0]],v[1])
		adj[v[1]] = append(adj[v[1]],v[0])
	}
	var dfs func(int)
	dfs = func(node int) {
		if seen[node]{
			return
		}
		seen[node] = true
		for _,v := range adj[node]{

			dfs(v)
		}
	}
	dfs(0)
	return len(seen) == n
}
