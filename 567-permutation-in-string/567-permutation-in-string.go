func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2){ 
        return false
    }
    s1C := make([]int,26)
    s2C := make([]int,26)
    for i,val := range s1{
        s1C[val - 'a']++
        s2C[rune(s2[i]) - 'a']++
    }
    matches := 0
    
    for i := 0; i<26 ;i++{
        if s1C[i] == s2C[i] {
            matches++
        }
    }
    l := 0
    // start at window end
    for r:= len(s1) ; r < len(s2) ;r++{
        if matches == 26{
            return true
        }
        
        index := rune(s2[r]) - 'a'
        
        s2C[index]++
        if s1C[index] == s2C[index]{
            matches++
        }else if s1C[index] +1 == s2C[index]{ // exceed char count 
            matches--
        }
        index = rune(s2[l]) - 'a'
        s2C[index]--
        if s1C[index] == s2C[index]{
            matches++
        }else if s1C[index] - 1 == s2C[index]{ // removed required values
            matches--
        }
        l++
    }
    return matches == 26
}