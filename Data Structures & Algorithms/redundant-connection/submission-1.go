func findRedundantConnection(edges [][]int) []int {
	parent := make(map[int]int)
	for i := 1; i<=len(edges); i++ {
		parent[i] = i
	}
	var find func(int)int
	find = func(x int)int{

		for x != parent[x]{
			x = parent[x]

		}
		return x
	}
	res := []int{}
	for _,v := range edges {
		if find(v[0]) == find(v[1]) {
			res = v
			break
		}
		parent[find(v[0])] = find(v[1])
	}
	return res
}
