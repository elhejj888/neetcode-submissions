func solve(board [][]byte) {
    rows, cols := len(board), len(board[0])
	seen:=make(map[string]bool)
	queue := [][2]int{}
	seenO:=make(map[string]bool)


	for i:=0; i<cols;i++{
		if board[0][i] == 'O'{
			queue = append(queue, [2]int{0,i})
			seenO[string(0)+","+string(i)] = true
		}
		if board[rows-1][i] == 'O'{
			queue = append(queue, [2]int{rows-1,i})
			seenO[string(rows-1)+","+string(i)] = true

		}
		seen[string(0)+","+string(i)] = true
		seen[string(rows-1)+","+string(i)] = true
	}

	for i := 0; i<rows; i++{
		if board[i][0] == 'O'{
			queue = append(queue, [2]int{i, 0})
			seenO[string(i)+","+string(0)] = true

		}
		if board[i][cols-1]=='O'{
			queue = append(queue, [2]int{i, cols-1})
			seenO[string(i)+","+string(cols-1)] = true

		}
		seen[string(i)+","+string(0)] = true
		seen[string(i)+","+string(cols-1)] = true
	}

	directions := [4][2]int{{0,1},{0,-1},{1,0},{-1,0}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		row, col := node[0], node[1]
		for _,dir := range directions{
			r, c:= row+dir[0], col+dir[1]
			if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] == 'X' || seen[string(r)+","+string(c)]{
				continue
			} 
		queue = append(queue, [2]int{r,c})
		seen[string(r)+","+string(c)] = true
		seenO[string(r)+","+string(c)] = true

		}
	}
	for i:=0; i<rows; i++{
		for j:=0; j<cols;j++{
			if seen[string(i)+","+string(j)] {
				continue
			}
			board[i][j]='X'
		}
	}
}
