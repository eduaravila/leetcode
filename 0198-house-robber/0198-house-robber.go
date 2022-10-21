func getMax(a,b int)int{
    if a > b{
        return a
    }
    return b
}

func rob(nums []int) int {
    var max int
    sums:= make([]int,len(nums))
    
    for i := range nums{
        if i > 1{
            sums[i] = getMax((nums[i]+(sums[i-1] - nums[i-1])), nums[i]+sums[i-2])
        }else{
            sums[i] = nums[i]
        }
        max = getMax(max,sums[i])
    }    
    return max
}

