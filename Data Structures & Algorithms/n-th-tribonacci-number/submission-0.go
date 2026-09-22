func tribonacci(n int) int {
	res := make(map[int]int)
	res[0] = 0
	res[1] = 1
	res[2] = 1

	var fn func(int)int
	fn = func(n int) int {
		if _,ok := res[n] ; ok {
			return res[n]
		}
		res[n] = fn(n-1) + fn(n-2) + fn(n-3)
		return res[n]
	}
	fmt.Println(res)
	return fn(n)
}
