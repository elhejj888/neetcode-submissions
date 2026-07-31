func isValid(s string) bool {
	if (len(s) % 2 != 0){
		return false
	}

    var parentheses []string
	for i:=0; i<len(s); i++ {
		if strings.Contains("([{",string(s[i])){
			parentheses = append(parentheses, string(s[i]))
		}
		switch {
			case string(s[i]) == ")" :
				if len(parentheses) > 0 && parentheses[len(parentheses) - 1] == "("{
					parentheses = parentheses[:len(parentheses) - 1]
				} else {
					return false
				}
			case string(s[i]) == "]" :
				if len(parentheses) > 0 && parentheses[len(parentheses) - 1] == "["{
					parentheses = parentheses[:len(parentheses) - 1]
				} else {
					return false
				}
			case string(s[i]) == "}" :
				if len(parentheses) > 0 && parentheses[len(parentheses) - 1] == "{"{
					parentheses = parentheses[:len(parentheses) - 1]
				} else {
					return false
				}
		}
	}
	
	if len(parentheses) == 0 {
		return true
	} else {
		return false
	}
}
