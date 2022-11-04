func hasAllCodes(s string, k int) bool {   
    keys := make(map[int]bool)
    for i := 0 ; i < (1 << k) ; i++{
        keys[i] = true
    }
    l,r := 0,k
    for r <= len(s){
        key,_ := strconv.ParseInt(s[l:r],2,64)
        if _,e := keys[int(key)];e{
            delete(keys,int(key))
        }
        l++
        r++        
    }    

    return len(keys) < 1
}

// generate binary codes 
// save them in a hash table
// k = 2
// 00, 01, 10, 11

// 
