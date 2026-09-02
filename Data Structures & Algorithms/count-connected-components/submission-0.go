func countComponents(n int, edges [][]int) int {
    seen := make(map[int]bool)
	adj := make(map[int][]int)
	count := 0

	for _,v := range edges {
		adj[v[0]] = append(adj[v[0]],v[1])
		adj[v[1]] = append(adj[v[1]],v[0])
	}

	var dfs func(int)
	dfs = func(node int){
		if seen[node]{
			return
		}
		seen[node]=true
		for _,v := range adj[node]{
			dfs(v)
		}
	}
	for i:=0 ; i < n ; i++ {
		if seen[i] {
			continue
		}
		dfs(i)
		count++
	}
	return count
}
