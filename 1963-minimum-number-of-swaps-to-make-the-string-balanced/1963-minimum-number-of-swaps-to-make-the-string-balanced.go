func minSwaps(s string) int {
    
    var res,op,cl int
    
    r := len(s)-1
    for i := range s {
        
        if cl > op {
            for s[r] == '[' {
                r--
            }            
            res++
            cl--
            op++
            r--
        }
        if s[i] == ']'{
            cl++
        }
        if s[i] == '['{
            op++
        }
        i++
    }
    return res
}