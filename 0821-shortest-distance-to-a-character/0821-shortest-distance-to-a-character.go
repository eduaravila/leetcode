func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}
func shortestToChar(s string, c byte) []int {    
    res := make([]int,len(s))
    inf := int(^uint(0)>>1)
    for i := range s{
        if s[i] == c{            
            continue
        }
        res[i]=inf
    }
    
    for i := range s{        
        if s[i] != c{
            continue
        }
        l,r,sum := i-1,i+1,0           
        for l >= 0 && s[l]!= c{
            sum++            
            res[l]=getMin(res[l],sum)
            l--
        }
        sum = 0
        for r < len(s) && s[r]!=c{
            sum++            
            res[r]=getMin(res[r],sum)
            r++
        }
    }
    return res
}