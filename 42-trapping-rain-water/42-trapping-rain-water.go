func getMax( a,b int)int{if a > b {return a}else{return b}}
func trap(height []int) int {
    res,l,r := 0,0,len(height)-1
    lm,rm := height[l],height[r]
    for l < r {                
        if height[l] < height[r]{            
            lm = getMax(lm,height[l])
            res += lm-height[l]
            l++
        }else{            
            rm = getMax(rm,height[r])
            res += rm- height[r]            
            r--
        }
    }
    
    
    return res
}

// there must be 1 of distance between l and r
// l and r must be >= than the "middle" elevations
// if l >= r or r >= l iterate from l to r and add to the result
// update l and r after that count