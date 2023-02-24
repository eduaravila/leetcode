func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func rob(nums []int) int {
    if len(nums) < 2{
        return nums[0]
    }
    return getMax(robert(nums[1:]),robert(nums[:len(nums)-1]))
}

func robert(nums []int) int{
    n := len(nums)
    dp := make([]int,n)   
    var max int
    for i:=0 ; i<len(nums); i++{
        if i < 2{
            dp[i] = nums[i]
        }else{
            dp[i] = getMax(nums[i] + dp[i-2], dp[i-1] - nums[i-1] + nums[i])
        }
        
        max = getMax(max,dp[i])
    }
    
    return max
}