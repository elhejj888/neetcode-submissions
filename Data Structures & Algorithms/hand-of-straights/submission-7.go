func isNStraightHand(hand []int, groupSize int) bool {

	if len(hand) % groupSize != 0 {
		return false
	}

	frequencies := make(map[int]int)

	for i := 0; i< len(hand); i++ {
		frequencies[hand[i]]++
	}

	sort.Ints(hand)

	for i := 0; i < len(hand); i++ {
		if frequencies[hand[i]] == 0 {
			continue
		}
	
		for j := hand[i]; j < hand[i] + groupSize; j++{
			if frequencies[j] == 0 {
				return false
			}

			frequencies[j]--
		}
	}

	return true
    
}
