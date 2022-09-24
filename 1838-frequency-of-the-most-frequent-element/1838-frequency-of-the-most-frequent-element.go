func getMax(a,b int )int{
    if a > b{
        return a
    }
    return b
}

func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)    
    var max int        
    r := len(nums)-1
    l := r
    for l >= 0 {
        k-= nums[r] - nums[l]
        if k < 0 {
            k += (nums[r] - nums[r-1]) * (r - l)
            r--
        }
        max = getMax(max, (r-l) + 1 )

        l--
    }
    
    return max
}