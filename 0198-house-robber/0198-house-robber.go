func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func rob(nums []int) int {
    dp := make([]int,len(nums))    
    var max int
    for i := 0 ; i < len(nums) ; i++{                
        if i < 2{
            dp[i] = nums[i]
        }else{
            dp[i] = getMax(nums[i] + dp[i-2], dp[i-1] - nums[i-1] + nums[i])        
        }
        max = getMax(max,dp[i])
    }
    
    fmt.Println(dp)
    
    return max
}


/*
[2,1,1,2]

[2, 1,3]

*/