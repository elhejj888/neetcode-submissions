import ("slices")
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
	res := this.fullMap[key][timestamp]
	if res != "" {
		return res 
	} else {
		var keys []int
		for k := range this.fullMap[key] {
			keys = append(keys, k)
		}
		slices.Sort(keys)
		if len(keys) == 0 || keys[0] > timestamp {
			return ""
		}

		right := len(keys) - 1
		left := 0
		
		for left < right {
			mid := left + (right - left) + 1 / 2
			if timestamp <= keys[mid] {
				right = mid - 1
			} else if timestamp >= keys[mid] {
				left = mid 
			} 
		}
		
		return this.fullMap[key][keys[left]] 
	}

}
