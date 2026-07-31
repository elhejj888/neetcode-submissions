type TimeMap struct {
	fullMap map[string]map[int]string
}

func Constructor() TimeMap {
	return TimeMap {
		fullMap : make(map[string]map[int]string),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.fullMap[key] == nil {
    	this.fullMap[key] = make(map[int]string) 
	}
	this.fullMap[key][timestamp] = value
}

func (this *TimeMap) Get(key string, timestamp int) string {
	fmt.Println(this.fullMap[key])
	res := this.fullMap[key][timestamp]
	if res != "" {
		return res 
	} else {
		index := -1
		var keys []int
		for k := range this.fullMap[key]{
			keys = append(keys,k)
		}
		for _, v := range keys{
			fmt.Println("value: ", v)
			if v < timestamp && v > index {
				index = v
			}
		} 
		fmt.Println(index)
		return this.fullMap[key][index]
		
	}

}
