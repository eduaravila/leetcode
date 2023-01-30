func getMin(a,b int)int{
    if a < b{
        return a        
    }
    return b
}

func minCostClimbingStairs(cost []int) int {   
    memo := make(map[int]int)
    solution(cost, len(cost)-1, memo)
    
    return getMin(memo[len(cost)-1],memo[len(cost)-2])
}

func solution(cost []int, i int, memo map[int]int) int{
    if val,e := memo[i];e{
        return val
    }
    if i < 0 {        
        return 0
    }
    
    memo[i] = cost[i] + getMin( solution(cost,i-1,memo) , solution(cost,i-2, memo))
       
    return memo[i]
}
