func canFinish(numCourses int, grid [][]int) bool {
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
        return true
    }
    for _,v := range grid {
        if !dfs(v[0]) {
            return false
        }
    }
    return true

    
}
