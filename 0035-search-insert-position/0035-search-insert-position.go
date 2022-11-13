func searchInsert(nums []int, target int) int {
    l,r := 0,len(nums)-1
    
    for l <= r{
        m := int((l+r) >> 1)
        
        if m > len(nums)-1 || nums[m] == target {
            return m
        }
        
        if nums[m] > target{
            r= m-1
        }else{
            l=m+1
        }
    }
    
    return l
}