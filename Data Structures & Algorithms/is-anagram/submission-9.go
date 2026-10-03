func isAnagram(s string, t string) bool {
    if len(s) != len (t) {
        return false;
    }
    count_s := make(map[byte]int)
    count_t := make(map[byte]int)
    for i := 0; i< len(s); i++ {
        count_s[s[i]]++
        count_t[t[i]]++
    }
    if len(count_s) != len(count_t){
        return false
    }
    for k,v := range count_s {
        if count_t[k] != v {
            return false
        }
    }
    return true

}
