func solve(board [][]byte) {
    rows, cols := len(board), len(board[0])
	queue := [][2]int{}
	seen:=[][]bool{}
	for i:=0; i<rows; i++{
		seen = append(seen, []bool{})
		for j:=0; j< cols; j++{
			seen[i] = append(seen[i],false)
		}
	}


	for i:=0; i<cols;i++{
		if board[0][i] == 'O'{
			seen[0][i]=true
			queue = append(queue, [2]int{0,i})
			
		}
		if board[rows-1][i] == 'O'{
			seen[rows-1][i]=true
			queue = append(queue, [2]int{rows-1,i})
			

		}

	}

	for i := 0; i<rows; i++{
		if board[i][0] == 'O'{
			seen[i][0]=true
			queue = append(queue, [2]int{i, 0})

		}
		if board[i][cols-1]=='O'{
			seen[i][cols-1]=true
			queue = append(queue, [2]int{i, cols-1})

		}
	}

	directions := [4][2]int{{0,1},{0,-1},{1,0},{-1,0}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		row, col := node[0], node[1]
		for _,dir := range directions{
			r, c:= row+dir[0], col+dir[1]
			if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] == 'X' || seen[r][c]{
				continue
			} 
		queue = append(queue, [2]int{r,c})
		seen[r][c]=true
		}
	}
	for i:=0; i<rows; i++{
		for j:=0; j<cols;j++{
			if seen[i][j] {
				continue
			}
			board[i][j]='X'
		}
	}
}
