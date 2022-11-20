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

// 1 min
// get the largest value in the array

// from 1 to max
// divide arr[current] / bananas, this is the amount of time it will take to each this stack of bananas if you can only eat k bananas
// remember to ceil the result, even is takes you 10 min extra it rounds to the complete hour
// do this with all the elements compare the total with h
// if the total < h, search in the left side to get a smaller amount 
/// else try to pick a bigger value

// observation
// if h == len(arr), pick the biggest value as a result


