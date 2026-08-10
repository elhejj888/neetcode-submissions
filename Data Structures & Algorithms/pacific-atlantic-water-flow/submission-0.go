func pacificAtlantic(heights [][]int) [][]int {
    rows, cols := len(heights), len(heights[0])
    
    seenPacific := make(map[string]bool)
    seenAtlantic := make(map[string]bool)
    

    pacific := [][2]int{}
    atlantic := [][2]int{}
    
    for i:=0; i < cols; i++{
        pacific = append(pacific, [2]int{0, i})
        atlantic = append(atlantic, [2]int{rows-1,i})
        seenPacific[string(0) + string(i)] = true
        seenAtlantic[string(rows-1) + string(i)] = true
    } 
    for i:=0; i < rows; i++{
        pacific = append(pacific, [2]int{i,0})
        atlantic = append(atlantic, [2]int{i, cols-1})
        seenPacific[string(i) + string(0)] = true
        seenAtlantic[string(i) + string(cols-1)] = true
    }
    directions:=[4][2]int{{0,1},{0,-1},{1,0},{-1,0}}

    pacificRes := append([][2]int(nil), pacific...)
    atlanticRes := append([][2]int(nil), atlantic...)   

    for len(pacific) > 0{
        node := pacific[0]
        pacific = pacific[1:]
        row, col := node[0],node[1]
        for _,dir := range directions{
            r , c := row+dir[0], col+dir[1]
            if r < 0 || c < 0 || r >= rows || c >= cols || heights[r][c] < heights[row][col] || seenPacific[string(r) + string(c)]{
                continue
            }
            seenPacific[string(r) + string(c)] = true
            pacificRes = append(pacificRes, [2]int{r,c})
            pacific = append(pacific, [2]int{r,c})
        }

    }
        for len(atlantic) > 0{
        node := atlantic[0]
        atlantic = atlantic[1:]
        row, col := node[0],node[1]
        for _,dir := range directions{
            r , c := row+dir[0], col+dir[1]
            if r < 0 || c < 0 || r >= rows || c >= cols || heights[r][c] < heights[row][col] || seenAtlantic[string(r) + string(c)]{
                continue
            }
            seenAtlantic[string(r) + string(c)] = true
            atlanticRes = append(atlanticRes, [2]int{r,c})
            atlantic = append(atlantic, [2]int{r,c})
        }

    }
    res := [][]int{}
    seen := make(map[string]bool)
    for i:=0;i<len(atlanticRes);i++{
        for j:=0; j<len(pacificRes);j++{
            if atlanticRes[i][0]==pacificRes[j][0] && atlanticRes[i][1]==pacificRes[j][1] && !seen[string(atlanticRes[i][0])+string(atlanticRes[i][1])]{
                seen[string(atlanticRes[i][0])+string(atlanticRes[i][1])] = true
                res = append(res, []int{atlanticRes[i][0],atlanticRes[i][1]})
            }
        }
    }

    return res

}
