func orangesRotting(grid [][]int) int {

	rows, cols := len(grid), len(grid[0])

	queue:=[][2]int{}
	freshCount := 0
	for i := 0; i<rows; i++{
		for j:=0; j<cols; j++{
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i,j})
			}
			if grid[i][j] == 1{
				freshCount++
			}
		}
	}
	
	// if len(queue) == 0 {
	// 	return 0
	// }


	directions := [4][2]int{{0,1},{0,-1},{1,0},{-1,0}}
	count := 0
for len(queue) > 0 {
    ringSize := len(queue)
    for i := 0; i < ringSize; i++ {
        node := queue[0]
        queue = queue[1:]
        row, col := node[0], node[1]

        for _, dir := range directions {
            r, c := row+dir[0], col+dir[1]
            if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] != 1 {
                continue
            }
            freshCount--
            queue = append(queue, [2]int{r, c})
            grid[r][c] = 2
        }
    }
    // finished one ring; if anything's left, it's next minute's ring
    if len(queue) > 0 {
        count++
    }
}

if freshCount != 0 {
    return -1
}
return count
}