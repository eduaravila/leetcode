func minSwaps(s string) int {
    op,cl := []int{},[]int{}
    var i,res int
    
    r := len(s)-1
    for i < len(s) {
        
        if len(cl) > len(op) {
            for s[r] == '[' {
                r--
            }            
            res++
            cl = cl[:len(cl)-1]
            op = append(op,i)
            r--
        }
        if s[i] == ']'{
            cl = append(cl,i)
        }
        if s[i] == '['{
            op = append(op,i)
        }
        i++
    }
    return res
}