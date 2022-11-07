func removeDuplicates(nums []int) int {
    var current,k int
    
    for current < len(nums){
        num := nums[current]
        nums[k] = num
        for current < len(nums) && nums[current] == num{
            current++
        }
        k++
        
    }
    return k
}