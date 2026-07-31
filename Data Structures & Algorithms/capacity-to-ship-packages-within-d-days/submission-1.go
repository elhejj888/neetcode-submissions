import ("slices")
func shipWithinDays(weights []int, days int) int {
	if days == len(weights) {
		return slices.Max(weights)
	}
	
	
	min := slices.Max(weights)
	max := 0
	for _, val := range weights {
		max+= val
	}
	

	for min < max {
	mid := min + (max - min)/2
	isValid := divideOnDays(weights,days, mid)
	if isValid <= days  {
		//max is too big, i need a smaller value
		max = mid
		
		
	} else if isValid > days {
		//max is too small, i need a bigger max
		min = mid + 1
		
	}
	}
	return min
}

func divideOnDays(weights []int, days int, maxWeight int) int {
	index := 0
	res := make([]int, 1)
	for i :=0 ; i < len(weights); i++{

		res[index] += weights[i]

		if res[index] > maxWeight {
			
			res = append(res,res[index] - weights[i])
			index++
			res[index] = weights[i]

		}

	}
	return len(res)
}