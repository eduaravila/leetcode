func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func partitionLabels(s string) []int {
    indexes := make([]int,('z'+1) - 'a')
    
    for i,c := range s{
        indexes[c-'a']= i
    }
    
    res := []int{}
    var l,c,r int
    for  c < len(s){
        r = getMax(r, indexes[s[c]-'a'])
        
        if c == r {
            res = append(res,r-l +1)
            l = c +1
            
        }
        c++
    }
    return res    
}