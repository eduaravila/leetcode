func climbStairs(n int) int {
    return solution(n, make([]int,46))
}


func solution(n int, memo []int)int{
    if memo[n] > 0 {
        return memo[n]
    }
    
    if n == 0 {
        return 1        
    }
    res := solution(n-1,memo)
    
    if n >=2 {
        res += solution(n-2,memo)        
        
    }
    memo[n] = res
    return res
}