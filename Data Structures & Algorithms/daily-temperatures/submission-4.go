type Stack struct {
		items [] int
}
func (s* Stack) Push(data int){
	s.items = append(s.items, data)
}

func (s* Stack) Pop(){
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s* Stack) Top() (int, error){
	if s.IsEmpty() {
		return 0, fmt.Errorf("Stack is Empty")
	}
	return s.items[len(s.items) - 1], nil
}

func (s* Stack) IsEmpty() bool{
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s* Stack) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()

}


func dailyTemperatures(temperatures []int) []int {
	warmDates := make([]int, len(temperatures))
	stackOfDays := Stack{}
	for index:= 0; index < len(temperatures); index++ {
		if stackOfDays.IsEmpty() {
			stackOfDays.Push(index)
			continue
		} else if item, err := stackOfDays.Top(); err == nil && temperatures[index] <= temperatures[item] {
			stackOfDays.Push(index)
			continue
		}




	for !stackOfDays.IsEmpty() {
		item, _ := stackOfDays.Top()

		if temperatures[index] <= temperatures[item] {
			break
		}

		warmDates[item] = index - item
		stackOfDays.Pop()
	}

	stackOfDays.Push(index)
	}
	fmt.Println(warmDates)
	return warmDates
}
