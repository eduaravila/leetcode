func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func candy(ratings []int) int {
    candies := make([]int,len(ratings))        
    var i int
    for i < len(ratings) {                
        if i > 0 && ratings[i] > ratings[i-1]{
            candies[i] = candies[i-1]+1
        }else{
            candies[i] = 1
        }
        i++
    }
    
    i = len(ratings)-1
    
    for i >= 0 {        
        if i > 0 && ratings[i-1] > ratings[i]{
            candies[i-1] = getMax(candies[i]+1, candies[i-1])
        }  
        i--
    }
    var res int
    
    for _,candie := range candies{
        res+=candie
    }
    return res
}