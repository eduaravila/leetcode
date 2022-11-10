
func rotate(nums []int, k int)  {
    if k < 1{
        return 
    }
    n:= len(nums)
    l,r := 0,n-1
    for l < r {
        nums[l],nums[r] = nums[r],nums[l]
        l++
        r--
    }
    
    limit := k%n-1
    l,r =0, limit
    
    for l<r{
        nums[l],nums[r] = nums[r],nums[l]
        l++
        r--
    }
    
    l = limit+1
    r = n-1
    for l<r{
        nums[l],nums[r] = nums[r],nums[l]
        l++
        r--
    }
}