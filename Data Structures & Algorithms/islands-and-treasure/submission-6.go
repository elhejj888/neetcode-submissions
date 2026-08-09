func islandsAndTreasure(grid [][]int) {
    rows, cols := len(grid), len(grid[0])

	INF := 2147483647
	queue := [][2]int{}
	for i := 0; i<rows; i++ {
		for j:=0; j<cols; j++{
			if grid[i][j] == 0 {
				queue = append(queue, [2]int{i,j})
			}
		}
	}
	if len(queue) == 0 {
		return
	}
	directions := [4][2]int{{0,1},{0,-1},{1,0},{-1.0}}

	for len(queue) > 0{
		node := queue[0]
		queue = queue[1:]

		row, col := node[0], node[1]
		for _,dir := range directions{
			r,c := row+dir[0], col+dir[1]
			if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] != INF {
				continue
			} 
			queue = append(queue, [2]int{r,c})
			grid[r][c] = grid[row][col] + 1
		}
	}
	
}
