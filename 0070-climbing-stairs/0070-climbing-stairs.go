func climbStairs(n int) int {
    return solution(n, make(map[int]int))
}

func solution(n int, memo map[int]int)int{
    if val,e := memo[n];e {
        return val
    }
    
    if n < 0{
        return 0
    }
    if n == 0{
        return 1
    }
    
    memo[n-1] = solution(n-1, memo)
    memo[n-2] = solution(n-2, memo)
    return memo[n-1] + memo[n-2]
}