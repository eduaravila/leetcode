func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2){
        return false
    }
    count1,count2 := make(map[rune]int),make(map[rune]int)
    
    for i,k := range s1{
        count1[k]++
        count2[rune(s2[i])]++
    }
    
    matches :=0
    
    for k,v := range count1{
        if v == count2[k]{
            matches+=v
        }
    }
    var l int
    for i := len(s1) ; i < len(s2) ; i++{
        if matches == len(s1){
            return true
        }
        key := rune(s2[i])
        
        count2[key]++
        if count2[key] == count1[key]{
            matches+=count1[key]
        }else if count1[key] +1 == count2[key]{
            matches -= count1[key]
        }
        
        key = rune(s2[l])
        
        count2[key]--
        if count2[key] == count1[key]{
            matches+=count1[key]
        }else if count1[key] - 1 == count2[key]{
            matches -= count1[key]
        }
        l++
    }
    return matches == len(s1)
}

// we need to find matches
// count the letter of s1 and s2
// total matches == len(s1)

// increase matches if counter1[key] == counter2[key]
// decrease otherwise
