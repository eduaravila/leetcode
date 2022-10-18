func canPartition(nums []int) bool {
    var sum int 
    
    for _, num := range nums{
        sum+= num
    }
    if sum % 2 != 0{
        return false
    }    
    memo := make(map[int]map[int]bool)
    

    return getSum(nums, sum/2, 0, memo)
}
func getSum(nums []int,sum, current int, memo map[int]map[int]bool )bool{
    
    if val, e:= memo[current][sum]; e {        
        return val
    }
    
    if sum == 0 {
        return true
    }
    
    if current > len(nums)-1 || sum < 0  ||( current > len(nums)-1 && sum > 0){
        return false
    }
    
    
    
    if getSum(nums, sum,current + 1,memo) || 
    getSum(nums,sum - nums[current],current + 1,memo){
        if _,e :=memo[current];!e{
            memo[current] = make(map[int]bool)
        }
        memo[current][sum]= true
        return true
    }
    
    if _,e :=memo[current];!e{
        memo[current] = make(map[int]bool)
    }
    memo[current][sum] = false
    return false
}