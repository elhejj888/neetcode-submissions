import "slices"
func isAnagram(s string, t string) bool {
    if len(s) != len (t) {
        return false;
    }
    s_slice := asciiArray(s)
    t_slice := asciiArray(t)
    for i := 0; i< len(s_slice); i++ {
        if s_slice[i] != t_slice[i]{
            return false
        }
    }
    return true

}

func asciiArray(s string) []int{
    var sum []int

        for i := 0; i<len(s); i++{
            sum = append(sum, int(s[i]))
        }
    slices.Sort(sum)
    return sum 
}
