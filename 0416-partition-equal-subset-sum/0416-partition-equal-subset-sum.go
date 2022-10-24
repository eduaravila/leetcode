func canPartition(nums []int) bool {
    var sum int
    
    for _, num := range nums{
        sum+= num
    }
    if sum % 2 != 0{
        return false
    }
    
    return solution(nums,0,sum/2,make(map[string]bool))
}

func solution(nums []int, current,sum int, memo map[string]bool)bool{
    key := fmt.Sprintf("%d-%d", sum,current)
    if val,e := memo[key]; e {
        return val
    }
    if sum < 0 || (sum > 0 && current > len(nums)-1){
        return false
    }
    
    if sum == 0{
        return true
    }
    
    memo[key] = solution(nums,current+1,sum-nums[current],memo) || solution(nums,current+1, sum,memo)
    return memo[key]
}