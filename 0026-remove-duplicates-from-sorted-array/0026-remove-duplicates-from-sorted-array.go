func removeDuplicates(nums []int) int {
    var current,k int
    
    for current < len(nums){
        
        nums[k] = nums[current]
        for current < len(nums) && nums[current] == nums[k]{
            current++
        }
        k++        
    }
    return k
}