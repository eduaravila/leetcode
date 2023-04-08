/*

we need the min of sum of perfect square numbers -> int
input -> int -> target 

bottom up
1+1+1+1+1+1.... -> 14 


var currentSum, base, target, min int 

base = base * base 

currentSum + 1 -> how many numbers have added in the current attempt 


* base case if base == 0 
* base case if target == 0 
    update the min value with the currentSum 


solution(0,target, target, inf)


for base > 0 {
    total := base * base 

    if total > target{
        base--
        continue
    }
    solution(currentSum+1, base, target - total, min)
    base--
}
*/

func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}


func numSquares(n int) int {
    return solution(n)
}
/*

type key struct{
    a,b int
}

func solution(currentSum, base, target int, memo map[key]int) int{ 
    k := key{base,target}
    if val, e := memo[k];e{
        return val
    }
    if target < 0{    
        return int(^uint(0)>>1)
    }
    if target == 0{        
        return currentSum
    }    
    min := int(^uint(0)>>1)       
    for base <= target  {
        total := base * base
        min = getMin(min, solution(currentSum+1, 1, target - total,memo))
        memo[key{base,target-total}] = min       
        base++
    }
    memo[k] = min
    return min 
}
.*/

func solution(target int) int {
    dp := make([]int,target+1)
    inf := int(^uint(0)>>1)
    
    for i := range dp {
        dp[i]= inf
    }
    dp[0] = 0
    
    for i:= 1; i<=target;i++{
        j :=1 
        for i-j*j >= 0{
            dp[i] = getMin(dp[i], dp[i-j*j] +1)
            j++
        }
    }
    
    return dp[target]
}