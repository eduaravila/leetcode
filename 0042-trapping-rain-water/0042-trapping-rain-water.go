func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func trap(height []int) int {
    l,r := 0,len(height)-1
    lm,rm := height[l], height[r]
    var res int
 
    for l < r {
        if height[l] < height[r]{
            lm = getMax(height[l],lm)
            res += lm - height[l]            
            l++    
        }else{            
            rm = getMax(height[r],rm)
            res += rm - height[r]
            r--
        }                
    }
    
    return res
}