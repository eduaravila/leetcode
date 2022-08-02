func checkInclusion(s1 string, s2 string) bool {
    s1C := make(map[rune]int)
    
    for _,val := range s1{
        s1C[val]++
    }
    
    l,r := 0,0
    wC := make(map[rune]int)
    for r < len(s2){
        wC[rune(s2[r])]++        
        if r - l+1 == len(s1){
            valid := true            
            for k,val := range s1C {                
                if wC[k] != val {
                    valid = false
                    break
                }
            }
            if valid {
                return valid
            }
            wC[rune(s2[l])]--
            l++
        }
        r++
    }
    return false
}