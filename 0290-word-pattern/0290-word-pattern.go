func wordPattern(pattern string, s string) bool {
    patterCounter := make(map[rune]string)
    patterCounter2 := make(map[string]string)
    sarr := strings.Split(s," ")
    if len(pattern) != len(sarr){
        return false
    }
    for i, c := range pattern{
        v1,e1 := patterCounter[c]
        v2,e2 := patterCounter2[sarr[i]]
        if (e1 && v1!= sarr[i]) || (e2 && v2 != string(c)) {
            return false
        }else{
            patterCounter[c] = sarr[i]
            patterCounter2[sarr[i]] = string(c)
        }
    }
    
    return true
}