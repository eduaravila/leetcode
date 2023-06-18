func getMin(a,b int)int{
    if a < b {
        return a
    }
    return b
}

func coinChange(coins []int, amount int) int {
    dp := []int{}
    
    inf := int(^uint(0)>>1)
    
    for i := 0 ; i <= amount ; i++{
        dp = append(dp, inf)
    }
    
    dp[0] = 0
    
    for _, coin := range coins{
        var i int
        next := i + coin
        for next < amount+1 {
            if dp[i] == inf{
                i++
                next = i + coin
                continue
            }
            dp[next] = getMin(dp[next], dp[i]+1) 
            i++
            next = i + coin
        }
    }
    
 
    if dp[amount] == inf{
        return -1
    }

    return dp[amount]       
}


/*

Input: coins = [1,2,5], amount = 11, answ = 3


[0,1,1,2,2,3,3,4,4,5,5,6] 
       
       
       
dp := []int, len(coints) +1

next = current + coin

if next < len(coins){
    next = min(dp[next], current + 1)
}

if min == inf{
    return -1
}

return dp[n]
*/