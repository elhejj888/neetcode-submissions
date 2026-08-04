func islandPerimeter(grid [][]int) int {
count := 0
for i:=0; i < len(grid); i++{
	for j:=0; j<len(grid[i]); j++{
		add := 4
		if grid[i][j] == 1 {

			if i+1 < len(grid) && grid[i+1][j] == 1 {
				add--
			} 

			if i-1 >= 0 && grid[i-1][j] == 1 {
				add--
			}

			if j+1 < len(grid[i]) && grid[i][j+1] == 1{
				add--
			}

			if j-1 >= 0 && grid[i][j-1] == 1 {
				add--
			}

			fmt.Println("add: ",add," i: ",i," j: ",j )
			count+=add
		}
	}
}
return count
}
