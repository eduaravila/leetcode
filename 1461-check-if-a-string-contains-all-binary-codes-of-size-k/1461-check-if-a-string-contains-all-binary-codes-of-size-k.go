func hasAllCodes(s string, k int) bool {       
    seen  := make(map[string]bool)
    l,r := 0,k
    for r <= len(s){
        
        seen[s[l:r]] = true
        l++
        r++        
    }    

    return len(seen) >= (1 << k)
}

// generate binary codes 
// save them in a hash table
// k = 2
// 00, 01, 10, 11

// 
