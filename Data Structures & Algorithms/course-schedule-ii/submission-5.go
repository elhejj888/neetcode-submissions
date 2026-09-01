func findOrder(numCourses int, grid [][]int) []int {
	res := []int{}
    rows := len(grid)
    seen := make(map[int]int)
    adj := make(map[int][]int)
    for i:=0; i<rows; i++ {

        adj[grid[i][0]] = append(adj[grid[i][0]], grid[i][1])
    }
	
    var dfs func(course int) bool

    dfs = func(course int) bool {
        if seen[course] == 2 {
            return true
        } 
        if seen[course] == 1 {
            return false 
        }
        
        seen[course] = 1
        for _, v := range adj[course]{

            if !dfs(v) {
                return false
            }
        }
        seen[course] = 2
		res = append(res,course)
        return true
    }
    for i:=0; i<numCourses; i++ {
        if !dfs(i) {
            return []int{}
        }
    }

	return res

}

