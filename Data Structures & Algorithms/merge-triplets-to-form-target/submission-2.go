func mergeTriplets(triplets [][]int, target []int) bool {
	
	a,b,c := false, false,false

	for i := 0; i<len(triplets); i++ {
		if triplets[i][0] > target[0] || triplets[i][1] > target[1] || triplets[i][2] > target[2] {
            continue
        }
		if triplets[i][0] == target[0] {
			a = true
		}
		if triplets[i][1] == target[1]{
			b = true
		}
		if triplets[i][2] == target[2]{
			c = true
		}
	}

	return a && b && c
}
