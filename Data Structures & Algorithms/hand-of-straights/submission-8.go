import("slices")

func isNStraightHand(hand []int, groupSize int) bool {

    if len(hand) % groupSize != 0 {
		return false
	}

	frq := make(map[int]int)

	for i := 0; i<len(hand); i++{
		frq[hand[i]]++
	}

	slices.Sort(hand)

	for i:=0; i<len(hand); i++ {
		if frq[hand[i]] == 0{
			continue
		}
		for j := hand[i]; j<hand[i] + groupSize; j++{
			if frq[j] == 0 {
				return false
			}
			frq[j]--
		}
	}
	return true
}
