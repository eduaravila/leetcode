func minSwaps(s string) int {
    cl := []int{}
    var i,res,op int
    
    r := len(s)-1
    for i < len(s) {
        
        if len(cl) > op {
            for s[r] == '[' {
                r--
            }            
            res++
            cl = cl[:len(cl)-1]
            op++
            r--
        }
        if s[i] == ']'{
            cl = append(cl,i)
        }
        if s[i] == '['{
            op++
        }
        i++
    }
    return res
}