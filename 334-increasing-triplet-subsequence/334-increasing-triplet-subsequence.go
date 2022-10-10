func increasingTriplet(nums []int) bool {
    
    m := int(^uint(0)>>1)
    l :=m
    for i := range nums {
        if nums[i] <= l{
            l = nums[i]
        }else if nums[i] <=m{
            m = nums[i]
        }else {
            return true
        }
    }
    return false
}
