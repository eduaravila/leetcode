func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func trap(height []int) int {
    
    l,r := 0,len(height)-1
    lm,rm := height[l],height[r]
    var res int
    for l < r {
        lc,rc := height[l],height[r]
        if lc < rc {
            lm = getMax(lc,lm)
            res += lm - lc
            l++
        }else{
            rm = getMax(rc,rm)
            res += rm - rc
            r--
        }        
    }
    return res
}