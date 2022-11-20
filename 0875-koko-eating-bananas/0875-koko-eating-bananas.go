func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}


func minEatingSpeed(piles []int, h int) int {
    var max int
        
    for _, pile := range piles{
        max = getMax(pile,max)
    }
    
    l,r,res := 1, max,max
    
    for l <= r {
        m := (l+r)>>1
        var totalHours int
        for _,pile := range piles{            
            totalHours += int(math.Ceil(float64(pile)/float64(m)))
        }
        
        if totalHours <= h{
            res = getMin(res,m)
            r = m-1
        }else{
            l = m+1
        }        
    }
    
    return res
}


//30,11,23,4,20

// 4,11,20,23,30


// Input: piles = [3,6,7,11], h = 8

// 27 
