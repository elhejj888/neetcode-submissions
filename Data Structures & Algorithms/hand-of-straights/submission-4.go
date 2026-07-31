import ("slices")
func isNStraightHand(hand []int, groupSize int) bool {
    
	if len(hand)/groupSize == 0 || len(hand) % groupSize != 0 {
		return false
	}

	slices.Sort(hand)

	frequencies := make(map[int]int)
	
	for i:=0; i < len(hand); i++ {
		frequencies[hand[i]]++
		if frequencies[hand[i]] > len(hand)/groupSize {
			return false
		}
	}

	count := 0
	i := 0
	n:=0
	tables := make(map[int][]int)
	
	for n < len(hand) / groupSize { //2
		for count < groupSize && i < len(hand) {  //2

			if frequencies[hand[i]] > 0 && !slices.Contains(tables[n],hand[i]){

			frequencies[hand[i]]--
			count++
			tables[n]=append(tables[n], hand[i])
		}
		
		i++ //2
		fmt.Println(i)
		}
		fmt.Println(tables)
		fmt.Println(frequencies)
		count = 0
		n++
		i = 0
	}

	if len(tables) != len(hand)/groupSize{
		return false
	}

	for i:=0 ; i< len(tables); i++ {
		if len(tables[i]) < groupSize {
			return false
		}
		for j:=0; j<len(tables[i]) - 1; j++{
			if tables[i][j+1] - tables[i][j] != 1 {
				return false
			}
		}
	}

	return true


}
