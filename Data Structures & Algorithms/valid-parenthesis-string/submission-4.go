func checkValidString(s string) bool {

	prtKeys := []int{}
	strKeys := []int{}
	for i:= 0; i < len(s); i++{
		if string(s[i]) == "(" {
			prtKeys = append(prtKeys, i)
		} else if string(s[i]) == "*"{
			strKeys = append(strKeys, i)
		} else if string(s[i]) == ")" {
			if len(prtKeys) > 0 {
				prtKeys = prtKeys[:len(prtKeys)-1]
			} else if len(strKeys) > 0{
				strKeys = strKeys[:len(strKeys)-1]
			}else {
				return false
			}
		}
		
	}

	fmt.Println(prtKeys)
	fmt.Println(strKeys)

	// if len(prtKeys) > len(strKeys) {
	// 	return false
	// }
	j := 0
	for len(prtKeys) > 0 && j < len(strKeys) {
		if prtKeys[0] > strKeys[j] {
			j++
		} else {
		
		prtKeys = prtKeys[1:]
		strKeys = strKeys[1:]

		}
	}
	fmt.Println(prtKeys)
	fmt.Println(strKeys)

	return len(prtKeys) == 0
	

    
}
