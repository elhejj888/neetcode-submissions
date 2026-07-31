import ("slices")
func canCompleteCircuit(gas []int, cost []int) int {
	index := -1

	visited := [] int {}

	for i := 0; i< len(gas); i++{
		if gas[i] - cost[i] < 0 {
			continue
		}
		currentGas := 0
		for j := i ; j <= len(gas) ; j++ {
			if j >= len(gas) {
				j = 0
			}
			if slices.Contains(visited, j){
				break
			}
			
			currentGas += gas[j]

			if currentGas - cost[j] < 0 {
				break
			} else {
				visited = append(visited , j)
				currentGas -= cost[j]
			}

			
		} 
		if len(visited) == len(gas) {
			index = i 
			break
		} else {
			visited = []int{}
		}

	}

	return index
}
