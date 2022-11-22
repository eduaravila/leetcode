func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func isSubset(s string,p string, removed map[int]bool)bool{
    l,r := 0,0    
    for l < len(p) {
        if r > len(s)-1{
            return false
        }
        if _,e := removed[r] ;e{
            r++
            continue                
        }
        if s[r] == p[l]{
            l++
        }     
        r++
    }
    return true
}

func maximumRemovals(s string, p string, removable []int) int {
    var max int
    
    l,r := 0,len(removable)-1
    for l <=r {
        m := (l+r)>>1
        set := make(map[int]bool)
        for _,r := range removable[:m+1]{
            set[r] =true
        }
        if isSubset(s,p,set){
            max = getMax(max,m+1)
            l=m+1
        }else{
            r= m-1
        }
    }
    return max
}