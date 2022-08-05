func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2){
        return false
    }
    count1 := make(map[rune]int)
    count2 := make(map[rune]int)
    
    for i,val := range s1{
        count1[val]++
        count2[rune(s2[i])]++
    }
    
    matches := 0 
    for k, val1 := range count1{
        if val2,e := count2[k];e && val1 == val2 {
            matches+= count2[k]
        }
    }   
    l := 0
    
    for i:= len(s1) ; i < len(s2) ; i++{
        if matches == len(s1){
            return true
        }
        k := rune(s2[i])
        count2[k]++
        
        if val,e := count1[k]; e && val == count2[k]{ 
            matches+=val
        }else if val,e := count1[k]; e && val+1 == count2[k]{
            matches-=val
        }
        
        k = rune(s2[l])
        count2[k]--
        if val,e := count1[k]; e && val == count2[k]{ 
            matches+=val
        }else if val,e := count1[k]; e && val-1 == count2[k]{
            matches-=val
        }
        
        l++
    }
    
    return matches == len(s1) 
}