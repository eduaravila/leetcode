func subarraySum(nums []int, k int) int {
    var res, sum int
    
    
    for i := range nums{
        sum = 0
        for y := i ; y < len(nums); y++{
            sum += nums[y]
            if sum == k{
                res++
                
            }
            
        }
    }
    return res
}