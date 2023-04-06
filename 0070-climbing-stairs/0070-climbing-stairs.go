func climbStairs(n int) int {
    dp := make([]int,n)
    if n < 3{
        return n
    }
    
    dp[0] = 1
    dp[1] = 2
    
    for i := 2 ; i < n ; i++{
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n-1]
}

/*
base case: 


[1,1,1,1]

[1,2,3,5,8]

n = 1 -> 1 | 1,

n = 2 -> 2 | 1, 2


n 3 -> 1,1,1 | 1,2 | 2,1

invalid: 1,1,2 | 2, 2


n = 4 -> 1,1,1,1 | 2,2 | 2,1,1 | 1,1,2 | 1,2,1 |


*/