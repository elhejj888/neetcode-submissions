type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	code := []byte{}
	for i:=0; i<len(strs); i++{
		for j:=0; j<len(strs[i]); j++{
		code = append(code, byte(strs[i][j]))
		}
		code = append(code, '\t')
	}
	var res strings.Builder
	for i := range code {
		res.WriteString(string(code[i]))
	} 
	return res.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	tmp := ""
	for i:=0; i<len(encoded); i++{
		if encoded[i] == '\t'{
			res = append(res, tmp)
			tmp = ""
			continue
		}
		tmp += string(encoded[i])
	}
	return res
}
