func canPartition(nums []int) bool {
    var sum int
    
    for _, num := range nums{
        sum+= num
    }
    if sum % 2 != 0{
        return false
    }
    
    return solution(nums,0,sum/2,make(map[int]map[int]bool))
}

func solution(nums []int, current,sum int, memo map[int]map[int]bool)bool{
   
    if val,e := memo[current][sum]; e {
        return val
    }
    if sum < 0 || (sum > 0 && current > len(nums)-1){
        return false
    }
    
    if sum == 0{
        return true
    }
    
    if solution(nums,current+1,sum-nums[current],memo) || solution(nums,current+1, sum,memo){
        if _, e := memo[current];!e{
            memo[current] = make(map[int]bool)
        }
        memo[current][sum] = true
        return true
    }
    
    if _, e := memo[current];!e{
        memo[current] = make(map[int]bool)
    }
    memo[current][sum] = false
    return false
}