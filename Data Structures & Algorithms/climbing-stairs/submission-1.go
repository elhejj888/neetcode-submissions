func climbStairs(n int) int {
    str := make(map[int]int)

	var count func(int)int

	count = func(n int) int { 

	if _,ok := str[n] ; ok {
		return str[n]
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	str[n] = count(n-2) + count(n-1)
	return str[n]
	}
	
	return count(n)
}
