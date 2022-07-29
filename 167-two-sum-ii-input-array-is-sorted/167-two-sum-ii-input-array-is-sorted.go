func twoSum(nums []int, target int) []int {        
    l,r := 0,len(nums)-1
    for l < r{
        currentSum  := nums[l] + nums[r]
        if currentSum > target {
            r--
        }else if currentSum < target {
            l++
        }else{
            return []int{l+1,r+1}
        }        
    }
    
    return []int{}
}